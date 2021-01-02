package dto

import (
	"time"
)

//根据书id获取所有章节
//根据章节id获取内容
type GetChapterNameByBookIdReq struct {
	BookId   int64 `json:"book_id"`
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Sort     int   `json:"sort"` //1正序 2到序
}
type ChapterData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	//BookName string    `json:"book_name"`
	BookId int64 `json:"book_id"`
	//Category int       `json:"category"`
	Ctime time.Time `json:"ctime"`
	Mtime time.Time `json:"mtime"`
}

type GetChapterNameByBookIdResp struct {
	List  []*ChapterData `json:"list"`
	Count int            `json:"count"`
}

type GetChapterContentByIdReq struct {
	Id          int    `json:"id"`
	ChapterName string `json:"chapter_name"`
	BookId      int64  `json:"book_id"`
}

//Count    int `json:"count"`
type ChapterInfo struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Content  string    `json:"content"`
	BookName string    `json:"book_name"`
	BookId   int64     `json:"book_id"`
	Category int       `json:"category"`
	PreId    int       `json:"pre_id"`
	NextId   int       `json:"next_id"`
	Ctime    time.Time `json:"ctime"`
	Mtime    time.Time `json:"mtime"`
}
type GetChapterContentByIdResp struct {
	List []*ChapterInfo `json:"list"`
	//PageTotal int             `json:"page_total"`
	//Count     int             `json:"count"`
}
