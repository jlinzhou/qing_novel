package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"qing_novel/dto"
	"qing_novel/model"
)

type RecommendHandler struct {
	bookService      *model.BookService
	recommendService *model.RecommendService
}

func NewRecommendHandler(orm *gorm.DB) *RecommendHandler {
	return &RecommendHandler{
		bookService:      model.NewBookService(orm),
		recommendService: model.NewRecommendService(orm),
	}
}

func (rh *RecommendHandler) GetRecommend(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			dto.RespFail(c, 100, err.Error(), nil)
			return
		}
	}()
	body, _ := c.GetRawData()
	req := dto.RecommendReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}
	resp := dto.RecommendResp{
		List: make([]*model.Book, 0),
	}

	recommendList, err := rh.recommendService.GetBookId(req.Type, req.Count, req.Category)
	if err != nil {
		logrus.Error(err)
		return
	}
	for _, recommend := range recommendList {
		book, err := rh.bookService.GetBookById(recommend.BookId)
		if err != nil {
			logrus.Error(err)
			return
		}
		resp.List = append(resp.List, book)
	}
	dto.RespSucess(c, resp)

	return
}
