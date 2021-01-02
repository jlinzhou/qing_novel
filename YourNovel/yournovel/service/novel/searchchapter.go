package novel

import (
	"encoding/json"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"html/template"
	"log"
	"net/url"
	"strings"
	"yournovel/yournovel/conf"
	"yournovel/yournovel/db/redis"
	"yournovel/yournovel/fetcher"
	"yournovel/yournovel/model"
)

func SearchChapterOfNovelApi(novelUrl string, novelName string) (*model.NovelChapterApi, error) {
	c := fetcher.NewNovelConnector()

	novelChapterApi := model.NovelChapterApi{
		ChapterInfo: make([]*model.ChapterInfo, 0),
	}
	requestURI, err := url.ParseRequestURI(novelUrl)
	if err != nil {
		return &novelChapterApi, err
	}
	//chapterInfoList := make([]*model.ChapterInfo, 0)
	host := requestURI.Host

	//if
	chapterInfoList, err1 := redis.RedisConnector.GetBook(novelUrl)
	if err1 != nil {

		chapterSelector, ok := conf.RuleConfig.Rule[host]["chapter_selector"].(string)
		if !ok {
			fmt.Println("ok")
			return &novelChapterApi, err
		}

		c.OnHTML(chapterSelector, func(element *colly.HTMLElement) {

			//.Parent()
			html, _ := element.DOM.Html()
			if err != nil {
				log.Println(err)
				return
			}
			doc, err := htmlquery.Parse(strings.NewReader(html))
			if err != nil {
				log.Println(err)
				return
			}
			nodes := htmlquery.Find(doc, `//a`)
			for _, node := range nodes {
				//fmt.Println(node)
				a := htmlquery.FindOne(node, "/")
				novelChapterApi.ChapterInfo = append(novelChapterApi.ChapterInfo, &model.ChapterInfo{
					ChapterName: htmlquery.InnerText(a),
					ChapterUrl:  htmlquery.SelectAttr(a, "href"),
				})
			}
			data, err := json.Marshal(novelChapterApi.ChapterInfo)
			if err != nil {
				log.Println(err)
				return
			}
			err = redis.RedisConnector.SetBook(novelUrl, data)
			if err != nil {
				log.Println(err)
				return
			}
		})
		novelChapterApi.Name = novelName
		novelChapterApi.OriginUrl = novelUrl
		novelChapterApi.LinkPrefix = conf.RuleConfig.Rule[host]["link_prefix"].(string)
		novelChapterApi.Domain = fmt.Sprintf("%s://%s", requestURI.Scheme, requestURI.Host)

		err = c.Visit(novelUrl)
		if err != nil {
			return &novelChapterApi, err
		}

	} else {
		for _, mchapter := range chapterInfoList {
			novelChapterApi.ChapterInfo = append(novelChapterApi.ChapterInfo, &model.ChapterInfo{
				ChapterName: mchapter.ChapterName,
				ChapterUrl:  mchapter.ChapterUrl,
			})
		}
		novelChapterApi.Name = novelName
		novelChapterApi.OriginUrl = novelUrl
		novelChapterApi.LinkPrefix = conf.RuleConfig.Rule[host]["link_prefix"].(string)
		novelChapterApi.Domain = fmt.Sprintf("%s://%s", requestURI.Scheme, requestURI.Host)
	}

	return &novelChapterApi, nil
}

func SearchChapterOfNovel(novelUrl string, novelName string) (*model.NovelChapter, error) {
	c := fetcher.NewNovelConnector()

	var novelChapter model.NovelChapter
	requestURI, err := url.ParseRequestURI(novelUrl)
	if err != nil {
		return &novelChapter, err
	}

	host := requestURI.Host
	chapterSelector, ok := conf.RuleConfig.Rule[host]["chapter_selector"].(string)
	if !ok {
		return &novelChapter, err
	}

	c.OnHTML(chapterSelector, func(element *colly.HTMLElement) {

		html, _ := element.DOM.Parent().Html()
		if err != nil {
			return
		}

		novelChapter.Content = template.HTML(html)
		novelChapter.Name = novelName
		novelChapter.OriginUrl = novelUrl
		novelChapter.LinkPrefix = conf.RuleConfig.Rule[host]["link_prefix"].(string)
		novelChapter.Domain = fmt.Sprintf("%s://%s", requestURI.Scheme, requestURI.Host)
	})

	err = c.Visit(novelUrl)
	if err != nil {
		return &novelChapter, err
	}

	return &novelChapter, nil
}
