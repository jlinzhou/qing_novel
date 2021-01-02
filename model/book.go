package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

//if "玄幻": 1,"修真": 2,"都市": 3,"穿越 历史": 4,"网游" : 5,"科幻":6,"女频":7,"二次元":8
//const{}
type Book struct {
	ID            int64     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT" `
	Title         string    `json:"title" gorm:"not null;index:title_idx"`
	Category      int       `json:"category" gorm:"not null;column:category"`
	Cover         string    `json:"cover"  gorm:"not null;column:cover"`
	Author        string    `json:"author"  gorm:"not null;column:author"`
	Intro         string    `json:"intro"  gorm:"not null;column:intro"`
	Tag           string    `json:"tag"  gorm:"not null;column:tag"`
	LastChapterId int       `json:"last_chapter_id"  gorm:"not null;column:last_chapter_id"`
	OssServer     int       `json:"oss_server"   gorm:"not null;column:oss_server"` //1 2
	Ctime         time.Time `json:"ctime"  gorm:"not null;column:ctime;default:CURRENT_TIMESTAMP;"`
	Mtime         time.Time `json:"mtime"  gorm:"not null;column:mtime;"`
}

//# 未分类 0
//# 玄幻1 修真2  都市3 穿越4 网游5 科幻6 名著7
func (Book) TableName() string {
	return "book"
}

func (m *Book) BeforeUpdate(scope *gorm.Scope) error {
	m.Mtime = time.Now()
	return nil
}

type BookService struct {
	orm *gorm.DB
}

func NewBookService(orm *gorm.DB) *BookService {
	return &BookService{orm: orm}
}

func (itf *BookService) Insert(data *Book) error {
	if err := itf.orm.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (itf *BookService) GetBookById(id int64) (*Book, error) {
	book := Book{}
	if err := itf.orm.Where("id=?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (itf *BookService) GetBookByIdList(id []int64) ([]*Book, error) {
	book := []*Book{}
	if err := itf.orm.Where("dtime is null and id in (?)", id).Find(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (itf *BookService) GetBookByNameList(name string, page int, pageSize int) ([]*Book, error) {
	bookList := []*Book{}
	if err := itf.orm.Where(fmt.Sprintf("dtime is null and ( title like '%%%s%%' or author  like '%%%s%%' )", name, name)).Limit(pageSize).Offset((page - 1) * pageSize).Find(&bookList).Error; err != nil {
		return nil, err
	}
	return bookList, nil
}
func (itf *BookService) GetBookByNameCount(name string) (int, error) {
	var count int
	if err := itf.orm.Model(&Book{}).Where(fmt.Sprintf("dtime is null and (title like '%%%s%%' or author  like '%%%s%%' )", name, name)).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (itf *BookService) GetBookOrderByCategoryList(catagory int, page int, pageSize int) ([]*Book, error) {
	bookList := []*Book{}
	if catagory <= 0 {
		if err := itf.orm.Where("dtime is null and cover !=''").Order("id desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&bookList).Error; err != nil {
			return nil, err
		}
	} else {
		if err := itf.orm.Where("dtime is null and cover !='' and category =?", catagory).Order("id desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&bookList).Error; err != nil {
			return nil, err
		}
	}

	return bookList, nil
}

func (itf *BookService) GetBookOrderByCategoryCount(catagory int) (int, error) {
	var count int
	if catagory <= 0 {
		if err := itf.orm.Model(&Book{}).Where("dtime is null").Count(&count).Error; err != nil {
			return 0, err
		}
	} else {
		if err := itf.orm.Model(&Book{}).Where("dtime is null and category =?", catagory).Count(&count).Error; err != nil {
			return 0, err
		}
	}

	return count, nil
}

func (itf *BookService) GetBookOrderByCtime(catagory int, page int, pageSize int) ([]*Book, error) {
	bookList := []*Book{}
	if catagory <= 0 {
		if err := itf.orm.Where("dtime is null").Order("ctime desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&bookList).Error; err != nil {
			return nil, err
		}
	} else {
		if err := itf.orm.Where("dtime is null and  category =? ", catagory).Order("ctime desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&bookList).Error; err != nil {
			return nil, err
		}
	}

	return bookList, nil
}

func (itf *BookService) GetBookOrderByMtime(catagory int, page int, pageSize int) ([]*Book, error) {
	bookList := []*Book{}
	if catagory <= 0 {
		if err := itf.orm.Where("dtime is null").Order("mtime desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&bookList).Error; err != nil {
			return nil, err
		}
	} else {
		if err := itf.orm.Where("dtime is null and category =? ", catagory).Order("mtime desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&bookList).Error; err != nil {
			return nil, err
		}
	}

	return bookList, nil
}
func (itf *BookService) UpdateLatestChapter(book_id int64, chapter_id int) error {
	params := make(map[string]interface{}, 0)
	params["last_chapter_id"] = chapter_id
	if err := itf.orm.Model(&Book{}).Where("id = ?", book_id).Update(params).Error; err != nil {
		return err
	}

	return nil
}

//func (itf *BookService) GetBookFront(catagory int, page int, pageSize int) ([]*Book, error) {
//	bookList := []*Book{}
//	if catagory <= 0 {
//		if err := itf.orm.Order("mtime desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&bookList).Error; err != nil {
//			return nil, err
//		}
//	} else {
//		if err := itf.orm.Where("category =?", catagory).Order("mtime desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&bookList).Error; err != nil {
//			return nil, err
//		}
//	}
//	return bookList, nil
//}
