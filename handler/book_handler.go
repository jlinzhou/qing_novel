package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"qing_novel/dto"
	"qing_novel/model"
)

const MaxPageSize = 500

type BookHandler struct {
	bookService    *model.BookService
	chapterService *model.ChapterService
}

func NewBookHandler(orm *gorm.DB) *BookHandler {
	return &BookHandler{
		bookService:    model.NewBookService(orm),
		chapterService: model.NewChapterService(orm),
	}
}
func (bh *BookHandler) GetBookById(c *gin.Context) {
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
	req := dto.GetBookByIdReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}
	if req.ID == 0 {
		err = errors.New("参数错误")
		return
	}

	book, err := bh.bookService.GetBookById(req.ID)
	if err != nil {
		logrus.Error(err)
		return
	}
	lastChapter, err := bh.chapterService.GetLastChapterByBookIdList(req.ID)
	if err != nil {
		logrus.Error(err)
		return
	}
	//count, err := bh.bookService.GetBookByNameCount(req.Name)
	//if err != nil {
	//	logrus.Error(err)
	//	return
	//}
	resp := dto.GetBookByIdResp{
		List:            book,
		LastChapterName: lastChapter.Name,
		LastChapterID:   lastChapter.ID,
		LastTime:        book.Mtime.Unix(),
	}

	dto.RespSucess(c, resp)

	return
}
func (bh *BookHandler) GetBookByIdList(c *gin.Context) {
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
	req := dto.GetBookByIdListReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}
	if len(req.IdList) == 0 {
		return
	}

	bookList, err := bh.bookService.GetBookByIdList(req.IdList)
	if err != nil {
		logrus.Error(err)
		return
	}
	resp := dto.GetBookByIdListResp{
		List: make([]*model.Book, 0),
	}
	resp.List = append(resp.List, bookList...)

	dto.RespSucess(c, resp)

	return
}
func (bh *BookHandler) GetBookByName(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			dto.RespFail(c, 100, err.Error(), nil)
			return
		}
	}()
	body, _ := c.GetRawData()
	req := dto.GetBookByNameReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}
	if req.Name == "" {
		err = errors.New("参数错误")
		return
	}
	if req.PageSize > MaxPageSize || req.PageSize <= 0 {
		req.PageSize = MaxPageSize
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	bookList, err := bh.bookService.GetBookByNameList(req.Name, req.Page, req.PageSize)
	if err != nil {
		logrus.Error(err)
		return
	}
	count, err := bh.bookService.GetBookByNameCount(req.Name)
	if err != nil {
		logrus.Error(err)
		return
	}
	resp := dto.GetBookByNameResp{
		List: make([]*dto.BookInfo, 0),
	}
	for _, book := range bookList {
		chapter, err := bh.chapterService.GetLastChapterByBookIdList(book.ID)
		if err != nil {
			logrus.Warn(err)
			continue
		}
		resp.List = append(resp.List, &dto.BookInfo{
			ID:              book.ID,
			Title:           book.Title,
			Cover:           book.Cover,
			Category:        book.Category,
			Author:          book.Author,
			LastChapterId:   chapter.ID,
			LastChapterName: chapter.Name,
			Ctime:           book.Ctime.Unix(),
			Mtime:           chapter.Mtime.Unix(),
		})

	}
	//resp := dto.GetBookByNameResp{
	//	List: make([]*dto.BookInfo, 0),
	//}
	//resp.List = append(resp.List, bookList...)
	resp.Count = count
	resp.PageTotal = resp.Count / req.PageSize
	dto.RespSucess(c, resp)

	return
}

func (bh *BookHandler) GetBookOrderByCategory(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			dto.RespFail(c, 100, err.Error(), nil)
			return
		}
	}()
	body, _ := c.GetRawData()
	req := dto.GetBookByCategoryReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}
	if req.PageSize > MaxPageSize || req.PageSize <= 0 {
		req.PageSize = MaxPageSize
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	bookList, err := bh.bookService.GetBookOrderByCategoryList(req.Category, req.Page, req.PageSize)
	if err != nil {
		logrus.Error(err)
		return
	}
	count, err := bh.bookService.GetBookOrderByCategoryCount(req.Category)
	if err != nil {
		logrus.Error(err)
		return
	}
	resp := dto.GetBookByCategoryResp{
		List: make([]*model.Book, 0),
	}
	resp.List = append(resp.List, bookList...)
	resp.Count = count
	resp.PageTotal = resp.Count / req.PageSize
	dto.RespSucess(c, resp)

	return
}

//get_book_by_mtime
func (bh *BookHandler) GetBookByMtime(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			dto.RespFail(c, 100, err.Error(), nil)
			return
		}
	}()
	body, _ := c.GetRawData()
	req := dto.GetBookByMtimeReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}
	if req.PageSize > MaxPageSize || req.PageSize <= 0 {
		req.PageSize = MaxPageSize
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	bookList, err := bh.bookService.GetBookOrderByMtime(req.Category, req.Page, req.PageSize)
	if err != nil {
		logrus.Error(err)
		return
	}
	//count, err := bh.bookService.GetBookOrderByCategoryCount(req.Category)
	//if err != nil {
	//	logrus.Error(err)
	//	return
	//}
	data := dto.GetBookByMtimeRes{
		List: make([]dto.GetBookByMtime, 0),
	}
	for _, mbook := range bookList {
		chapter := &model.Chapter{}
		if mbook.LastChapterId != 0 {
			chapter, err = bh.chapterService.GetChapterContentById(mbook.LastChapterId, mbook.ID)
			if err != nil {
				logrus.Warn(err)
				continue
			}
		} else {
			chapter, err = bh.chapterService.GetLastChapterByBookIdList(mbook.ID)
			if err != nil {
				logrus.Warn(err)
				continue
			}
			err = bh.bookService.UpdateLatestChapter(mbook.ID, chapter.ID)
			if err != nil {
				logrus.Warn(err)

			}
		}

		bookByMtime := dto.GetBookByMtime{}
		bookByMtime.Book = mbook
		bookByMtime.ChapterId = chapter.ID
		bookByMtime.ChapterName = chapter.Name

		data.List = append(data.List, bookByMtime)
	}

	//data.Count = count
	data.PageTotal = data.Count / req.PageSize
	dto.RespSucess(c, data)
}

func (bh *BookHandler) GetBookByCtime(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			dto.RespFail(c, 100, err.Error(), nil)
			return
		}
	}()
	body, _ := c.GetRawData()
	req := dto.GetBookByCtimeReq{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		logrus.Error(err)
		return
	}
	if req.PageSize > MaxPageSize || req.PageSize <= 0 {
		req.PageSize = MaxPageSize
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	bookList, err := bh.bookService.GetBookOrderByCtime(req.Category, req.Page, req.PageSize)
	if err != nil {
		logrus.Error(err)
		return
	}
	//count, err := bh.bookService.GetBookOrderByCategoryCount(req.Category)
	//if err != nil {
	//	logrus.Error(err)
	//	return
	//}
	resp := dto.GetBookByCtimeRes{
		List: make([]*model.Book, 0),
	}
	resp.List = append(resp.List, bookList...)
	//resp.Count = count
	resp.PageTotal = resp.Count / req.PageSize
	dto.RespSucess(c, resp)
}
