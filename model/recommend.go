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
type Recommend struct {
	ID       int       `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT" `
	BookId   int64     `json:"book_id" gorm:"not null;index:book_id_idx"`
	Category int       `json:"category" gorm:"not null;column:category"`
	Weight   int       `json:"weight"  gorm:"not null;column:weight"`
	Type     int       `json:"type" gorm:"not null;column:type"` //1总推荐 2上周强推 3分类推荐
	Ctime    time.Time `json:"ctime"  gorm:"not null;column:ctime;default:CURRENT_TIMESTAMP;"`
	Mtime    time.Time `json:"mtime"  gorm:"not null;column:mtime;"`
}

func (Recommend) TableName() string {
	return "recommend"
}

type RecommendService struct {
	orm *gorm.DB
}

func NewRecommendService(orm *gorm.DB) *RecommendService {
	return &RecommendService{orm: orm}
}

//excel插入

func (rs *RecommendService) GetBookId(mtype int, count int, category int) ([]*Recommend, error) {
	recommendList := []*Recommend{}

	if mtype == 1 || mtype == 2 {
		if err := rs.orm.Where("type=?", mtype).Order("weight").Limit(count).Find(&recommendList).Error; err != nil {
			return nil, err
		}
		return recommendList, nil
	}
	if err := rs.orm.Where("type=? and category=?", mtype, category).Order("weight").Limit(count).Find(&recommendList).Error; err != nil {
		return nil, err
	}
	return recommendList, nil
}
