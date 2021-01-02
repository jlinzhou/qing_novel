package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yournovel/yournovel/service/novel"
	"yournovel/yournovel/tool"
)

func novelChapterApi(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			RespFail(c, 100, err.Error(), nil)
			return
		}
	}()

	webSiteUrl, exist := c.GetQuery("url")
	if !exist {
		err = errors.New("源网址不存在")
		return
	}
	novelName, exist := c.GetQuery("novel_name")
	if !exist {
		err = errors.New("小说名不存在")
		return
	}

	novelChapter, err := novel.SearchChapterOfNovelApi(webSiteUrl, novelName)
	if err != nil {
		fmt.Println(err)
		err = errors.New("章节获取失败")
		return
	}
	data := map[string]interface{}{
		"chapter": novelChapter,
		"head":    "chapter_head",
	}
	RespSucess(c, data)
}

func novelChapter(c *gin.Context) {

	webSiteUrl, exist := c.GetQuery("url")
	if !exist {
		tool.ErrorResponse(c, "源网址不存在", webSiteUrl)
		return
	}
	novelName, exist := c.GetQuery("novel_name")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, webSiteUrl)
		return
	}

	novelChapter, err := novel.SearchChapterOfNovel(webSiteUrl, novelName)
	if err != nil {
		//c.Redirect(http.StatusMovedPermanently, webSiteUrl)
		return
	}

	c.HTML(http.StatusOK, "chapter_index.html", gin.H{
		"chapter": novelChapter,
		"head":    "chapter_head",
	})
}
