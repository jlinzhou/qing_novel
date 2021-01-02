package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math"
	"time"
)

const (
	NoContent   = "1"
	ChapterName = "chapter_"
	MaxNovel    = 100
)

type Chapter struct {
	ID   int    `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT" `
	Name string `json:"name" gorm:"not null;index:title_idx"`
	//Content   string    `json:"content" gorm:"not null;column:content"`
	//BookName  string    `json:"book_name"  gorm:"not null;column:book_name"`
	BookId    int64     `json:"book_id"  gorm:"not null;column:book_id"`
	Url       string    `json:"url"   gorm:"not null;column:url"`
	LocalPath string    `json:"local_path"   gorm:"not null;column:local_path"`
	Ctime     time.Time `json:"ctime"  gorm:"not null;column:ctime;"`
	Mtime     time.Time `json:"mtime"  gorm:"not null;column:mtime;"`
}

//"id","name","content","book_name","book_id","ctime","mtime"
//func (Chapter) TableName() string {
//	return "Chapter"
//}

//func (m *Book) BeforeUpdate(scope *gorm.Scope) error {
//	m.Mtime = time.Now()
//	return nil
//}

type ChapterService struct {
	orm *gorm.DB
}

func NewChapterService(orm *gorm.DB) *ChapterService {
	return &ChapterService{orm: orm}
}
func (cs *ChapterService) TableName(bookId int64) string {
	tableName := ChapterName + fmt.Sprintf("%v", math.Ceil(float64(bookId)/float64(MaxNovel)))
	return tableName
}

func (cs *ChapterService) GetChapterNameByBookIdList(bookId int64, page int, pageSize int, sort int) ([]*Chapter, error) {
	chapterList := []*Chapter{}
	sortStr := ""
	if sort == 2 {
		sortStr = "desc"
	} else {
		sortStr = "asc"
	}

	if err := cs.orm.Table(cs.TableName(bookId)).Where("book_id=?", bookId).Order("id " + sortStr).Limit(pageSize).Offset((page - 1) * pageSize).Find(&chapterList).Error; err != nil {
		return nil, err
	}
	return chapterList, nil
}
func (cs *ChapterService) GetChapterContentByName(name string, bookId int64) ([]*Chapter, error) {
	chapterList := []*Chapter{}
	if err := cs.orm.Table(cs.TableName(bookId)).Where("book_id=? and name=?", bookId, name).Find(&chapterList).Error; err != nil {
		return nil, err
	}
	return chapterList, nil
}

func (cs *ChapterService) GetChapterNameByBookIdCount(bookId int64) (int, error) {
	var count int
	if err := cs.orm.Table(cs.TableName(bookId)).Where("book_id=?", bookId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (cs *ChapterService) GetChapterContentById(id int, bookId int64) (*Chapter, error) {
	chapter := Chapter{}
	if err := cs.orm.Table(cs.TableName(bookId)).Where("id=?", id).First(&chapter).Error; err != nil {
		return nil, err
	}
	return &chapter, nil
}

func (cs *ChapterService) GetLastChapterByBookIdList(bookId int64) (*Chapter, error) {
	chapter := Chapter{}
	if err := cs.orm.Table(cs.TableName(bookId)).Where("book_id=?", bookId).Order("id desc").Limit(1).First(&chapter).Error; err != nil {
		return nil, err
	}
	return &chapter, nil
}

func (cs *ChapterService) GetMinIdMaxId(bookId int64) (int, int, error) {
	minIdSlice := []int{}
	maxIdSlice := []int{}
	tableName := cs.TableName(bookId)
	//fmt.Println("tableName:", tableName)
	if err := cs.orm.Raw(fmt.Sprintf("SELECT min(id) as id FROM %s WHERE book_id = ?", tableName), bookId).Pluck("id", &minIdSlice).Error; err != nil {
		return 0, 0, err
	}
	if err := cs.orm.Raw(fmt.Sprintf("SELECT max(id) as id FROM %s WHERE book_id = ?", tableName), bookId).Pluck("id", &maxIdSlice).Error; err != nil {
		return 0, 0, err
	}
	return minIdSlice[0], maxIdSlice[0], nil
}
func (cs *ChapterService) GetPreChapter(id int, bookId int64) (int, error) {
	idSlice := []int{}
	tableName := cs.TableName(bookId)
	//fmt.Println("tableName:", tableName)
	if err := cs.orm.Raw(fmt.Sprintf("SELECT id FROM %s WHERE book_id = ? and id<? order by id desc limit 1", tableName), bookId, id).Pluck("id", &idSlice).Error; err != nil {
		return 0, err
	}
	if len(idSlice) > 0 {
		return idSlice[0], nil
	} else {
		return 0, nil
	}

}
func (cs *ChapterService) GetNextChapter(id int, bookId int64) (int, error) {
	idSlice := []int{}
	tableName := cs.TableName(bookId)
	//fmt.Println("tableName:", tableName)
	if err := cs.orm.Raw(fmt.Sprintf("SELECT id FROM %s WHERE book_id = ? and id>? order by id  limit 1", tableName), bookId, id).Pluck("id", &idSlice).Error; err != nil {
		return 0, err
	}
	if len(idSlice) > 0 {
		return idSlice[0], nil
	} else {
		return 0, nil
	}
}
func (cs *ChapterService) UpdateLocalPath(bookId int64, chapterId int, localPath string) error {

	params := make(map[string]string, 0)
	params["local_path"] = localPath
	if err := cs.orm.Table(cs.TableName(bookId)).Where("book_id=? and id", bookId, chapterId).Update(params).Error; err != nil {
		return err
	}
	return nil
}
