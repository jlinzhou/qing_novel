package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"qing_novel/dto"
	"qing_novel/model"
)

type RankHandler struct {
	bookService *model.BookService
	rankService *model.RankService
}

func NewRankHandler(orm *gorm.DB) *RankHandler {
	return &RankHandler{
		bookService: model.NewBookService(orm),
		rankService: model.NewRankService(orm),
	}
}

func (rh *RankHandler) GetRank(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			dto.RespFail(c, 100, err.Error(), nil)
			return
		}
	}()
	body, _ := c.GetRawData()
	req := dto.RankReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}
	resp := dto.RankResp{
		List: make([]*dto.RankInfo, 0),
	}
	if req.PageSize > MaxPageSize || req.PageSize <= 0 {
		req.PageSize = MaxPageSize
	}
	if req.Page <= 0 {
		req.Page = 1
	}

	rankList, err := rh.rankService.GetRankList(req.Type, req.Page, req.PageSize)
	if err != nil {
		logrus.Error(err)
		return
	}
	book_id_list := make([]int64, 0)
	for _, rank := range rankList {
		book_id_list = append(book_id_list, rank.BookId)
	}
	book_list, err := rh.bookService.GetBookByIdList(book_id_list)
	if err != nil {
		logrus.Error(err)
		return
	}
	for i, book := range book_list {
		resp.List = append(resp.List, &dto.RankInfo{
			BookId:   book.ID,
			Weight:   i + 1,
			Title:    book.Title,
			Category: book.Category,
			Intro:    book.Intro,
			Author:   book.Author,
			Cover:    book.Cover,
			Ctime:    book.Ctime,
		})
	}

	dto.RespSucess(c, resp)

	return
}
