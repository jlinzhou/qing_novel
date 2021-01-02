package dto

import (
	"qing_novel/model"
)

//根据书名进行模糊搜索()(返回id，书名...)
//获取热门图书(count(数量) category（分类）)(返回id，书名...)
type GetBookByIdReq struct {
	ID int64 `json:"id"`
}
type GetBookByIdResp struct {
	List            *model.Book `json:"list"`
	LastTime        int64       `json:"last_time"`
	LastChapterName string      `json:"last_chapter_name"`
	LastChapterID   int         `json:"last_chapter_id"`
}

type GetBookByIdListReq struct {
	IdList []int64 `json:"id_list"`
}
type GetBookByIdListResp struct {
	List []*model.Book `json:"list"`
	//LastTime        int64       `json:"last_time"`
	//LastChapterName string      `json:"last_chapter_name"`
	//LastChapterID   int         `json:"last_chapter_id"`
}
type GetBookByNameReq struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Name     string `json:"name"`
}
type BookInfo struct {
	ID              int64  `json:"id" `
	Title           string `json:"title"`
	Cover           string `json:"cover"`
	Category        int    `json:"category"`
	Author          string `json:"author" `
	LastChapterId   int    `json:"last_chapter_id"`
	LastChapterName string `json:"last_chapter_name"`
	Ctime           int64  `json:"ctime"`
	Mtime           int64  `json:"mtime"`
}

type GetBookByNameResp struct {
	List      []*BookInfo `json:"list"`
	PageTotal int         `json:"page_total"`
	Count     int         `json:"count"`
}

type GetBookByCategoryReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Category int `json:"category"`
}

//Count    int `json:"count"`

type GetBookByCategoryResp struct {
	List      []*model.Book `json:"list"`
	PageTotal int           `json:"page_total"`
	Count     int           `json:"count"`
}

type GetBookByCtimeReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Category int `json:"category"`
}

type GetBookByCtimeRes struct {
	List      []*model.Book `json:"list"`
	PageTotal int           `json:"page_total"`
	Count     int           `json:"count"`
}

type GetBookByMtimeReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Category int `json:"category"`
}
type GetBookByMtime struct {
	Book        *model.Book `json:"book"`
	ChapterId   int         `json:"chapter_id"`
	ChapterName string      `json:"chapter_name"`
}
type GetBookByMtimeRes struct {
	List      []GetBookByMtime `json:"list"`
	PageTotal int              `json:"page_total"`
	Count     int              `json:"count"`
}
