package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"qing_novel/common/crawl"

	"qing_novel/dto"
	"qing_novel/model"

	"time"
)

//根据书id获取所有章节
//根据章节id获取内容

type ChapterHandler struct {
	crawl          *crawl.CrawlBook
	book           *model.BookService
	chapterService *model.ChapterService
}

func NewChapterHandler(orm *gorm.DB) *ChapterHandler {
	return &ChapterHandler{
		crawl:          crawl.NewCrawlBook(),
		book:           model.NewBookService(orm),
		chapterService: model.NewChapterService(orm),
	}
}

func (ch *ChapterHandler) GetChapterNameByBookId(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			dto.RespFail(c, 100, err.Error(), nil)
			return
		}
	}()

	body, err := c.GetRawData()
	if err != nil {
		logrus.Warn(err)
		return
	}
	//
	req := dto.GetChapterNameByBookIdReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 || req.PageSize > 500 {
		req.PageSize = 20
	}

	if req.BookId <= 0 {
		err = errors.New("参数错误")
		return
	}
	//
	chapterList, err := ch.chapterService.GetChapterNameByBookIdList(req.BookId, req.Page, req.PageSize, req.Sort)
	if err != nil {
		logrus.Error(err)
		return
	}
	//fmt.Println(chapterList)
	count, err := ch.chapterService.GetChapterNameByBookIdCount(req.BookId)
	if err != nil {
		logrus.Error(err)
		return
	}
	//chapterData := []*dto.ChapterData{}
	resp := dto.GetChapterNameByBookIdResp{
		List: make([]*dto.ChapterData, 0),
	}
	for _, mchapter := range chapterList {
		resp.List = append(resp.List, &dto.ChapterData{
			ID:     mchapter.ID,
			Name:   mchapter.Name,
			BookId: mchapter.BookId,
			Ctime:  mchapter.Ctime,
			Mtime:  mchapter.Mtime,
		})
	}

	//resp.List = append(resp.List, chapterData...)
	resp.Count = count
	dto.RespSucess(c, resp)

	return
}

func (ch *ChapterHandler) GetChapterContentById(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			dto.RespFail(c, 100, err.Error(), nil)
			return
		}
	}()

	body, err := c.GetRawData()
	if err != nil {
		logrus.Warn(err)
		return
	}
	req := dto.GetChapterContentByIdReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}

	chapter, err := ch.chapterService.GetChapterContentById(req.Id, req.BookId)
	if err != nil {
		logrus.Error(err)
		return
	}
	book, err := ch.book.GetBookById(req.BookId)
	if err != nil {
		logrus.Error(err)
		return
	}

	preId, _ := ch.chapterService.GetPreChapter(chapter.ID, req.BookId)
	nextId, _ := ch.chapterService.GetNextChapter(chapter.ID, req.BookId)


	chapterInfo := dto.ChapterInfo{}
	url := ""
	if book.OssServer == 1 {
		url = "http://www.xbiquge.la" + chapter.Url
	} else {
		url = chapter.Url
	}
	resContent,_, err := ch.crawl.BiQuParseUrls(url)
	if err != nil {
		logrus.Warn(err)
		return
	}

	chapterInfo = dto.ChapterInfo{
		ID:       chapter.ID,
		Name:     chapter.Name,
		Content:  resContent,
		BookId:   chapter.BookId,
		BookName: book.Title,
		Category: book.Category,
		PreId:    preId,
		NextId:   nextId,
	}
	//logrus.Info(content)

	resp := dto.GetChapterContentByIdResp{
		List: make([]*dto.ChapterInfo, 0),
	}
	resp.List = append(resp.List, &chapterInfo)
	dto.RespSucess(c, resp)

	return
}

func (ch *ChapterHandler) getChapterContentByPost(url string, data interface{}, contentType string) ([]byte, error) {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
		return []byte{}, err
	}
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		logrus.Error(err)
		return []byte{}, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Error(err)
		return []byte{}, err
	}
	return result, nil
}
