package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yournovel/yournovel/service/novel"
)

func novelContentApi(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			RespFail(c, 100, err.Error(), nil)
			return
		}
	}()

	contentUrl, exist := c.GetQuery("content_url")
	if !exist {
		err = errors.New("content_url不存在")
		return
	}
	novelName, exist := c.GetQuery("novel_name")
	if !exist {
		err = errors.New("novel_name不存在")
		return
	}

	fmt.Println("contentUrl:", contentUrl)

	content, err := novel.SearchContentOfNovel(contentUrl)
	if err != nil {
		err = errors.New("章节解析错误")
		return
	}
	content.NovelName = novelName
	content.ContentUrl = contentUrl

	data := map[string]interface{}{
		"content":    content,
		"novel_name": novelName,
	}
	RespSucess(c, data)

}

func novelContent(c *gin.Context) {

	chapterUrl, exist := c.GetQuery("chapter_url")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	contentUrl, exist := c.GetQuery("content_url")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, chapterUrl)
		return
	}
	novelName, exist := c.GetQuery("novel_name")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, chapterUrl)
		return
	}

	contentTitle, exist := c.GetQuery("content_title")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, chapterUrl)
		return
	}
	fmt.Println("contentUrl:", contentUrl)

	content, err := novel.SearchContentOfNovel(contentUrl)
	if err != nil {
		fmt.Println(content)
		//c.Redirect(http.StatusMovedPermanently, chapterUrl)
		return
	}
	content.NovelName = novelName
	content.ContentUrl = contentUrl

	c.HTML(http.StatusOK, "content_index.html", gin.H{
		"content":       content,
		"chapter_url":   chapterUrl,
		"novel_name":    novelName,
		"content_title": contentTitle,
		"head":          "content_head",
	})
}
