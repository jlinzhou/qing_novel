package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//总推荐 上周强推
//分类推荐
//分类左下推荐
//搜索框远程搜索
//
type Rank struct {
	ID       int       `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT" `
	BookId   int64     `json:"book_id" gorm:"not null;index:book_id_idx"`
	Mtype    int       `json:"mtype" gorm:"not null;column:mtype"`
	Weight   int       `json:"weight"  gorm:"not null;column:weight"`
	BookName string    `json:"book_name"   gorm:"not null;column:book_name"`
	Author   string    `json:"author"   gorm:"not null;column:author"`
	Type     int       `json:"type" gorm:"not null;column:type"` //点击周11，月12，总13；推荐周21，月22，完23；完本周31，月32，完33
	Ctime    time.Time `json:"ctime"  gorm:"not null;column:ctime;default:CURRENT_TIMESTAMP;"`
}

func (Rank) TableName() string {
	return "rank"
}

type RankService struct {
	orm *gorm.DB
}

func NewRankService(orm *gorm.DB) *RankService {
	return &RankService{orm: orm}
}

//excel插入

func (rs *RankService) GetRankList(mtype int, page int, pageSize int) ([]*Rank, error) {
	rankList := []*Rank{}

	if err := rs.orm.Where("type=?", mtype).Order("weight").Limit(pageSize).Offset((page - 1) * pageSize).Find(&rankList).Error; err != nil {
		return nil, err
	}
	return rankList, nil
}
