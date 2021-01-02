package dto

import (
	"time"
)

type RankReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Type     int `json:"type"`
}
type RankInfo struct {
	BookId   int64     `json:"book_id"`
	Weight   int       `json:"weight"`
	Title    string    `json:"title"`
	Category int       `json:"category"`
	Cover    string    `json:"cover"`
	Author   string    `json:"author"`
	Intro    string    `json:"intro"`
	Ctime    time.Time `json:"ctime"`
}

type RankResp struct {
	List []*RankInfo `json:"list"`
	//Count int          `json:"count"`
}
