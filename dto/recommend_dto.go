package dto

import "qing_novel/model"

//type category book_id
//mtype int, count int, category
type RecommendReq struct {
	Count    int `json:"count"`
	Type     int `json:"type"`
	Category int `json:"Category"`
}

type RecommendResp struct {
	List []*model.Book `json:"list"`
	//Count int          `json:"count"`
}
