package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	//"net/http"
	"qing_novel/common/pool"
	"qing_novel/config"
	"qing_novel/handler"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logrus.Info(fmt.Sprintf("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri))
	}
}

func StartServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(LoggerToFile())
	apiPrefix := "/api"
	g := r.Group(apiPrefix)
	bookHandler := handler.NewBookHandler(pool.DefaultContent().Gorm)
	g.POST("/get_book_by_id", bookHandler.GetBookById)
	g.POST("/get_book_by_id_list", bookHandler.GetBookByIdList)
	g.POST("/get_book_by_name", bookHandler.GetBookByName)
	g.POST("/get_book_by_category", bookHandler.GetBookOrderByCategory)
	g.POST("/get_book_by_mtime", bookHandler.GetBookByMtime)
	g.POST("/get_book_by_ctime", bookHandler.GetBookByCtime)

	chapterHandler := handler.NewChapterHandler(pool.DefaultContent().Gorm)
	//根据category 和book mtime 查找 最新的
	//最新入库小说ctime查找 根据category 和book ctime
	g.POST("/get_chapter_info", chapterHandler.GetChapterNameByBookId)
	g.POST("/get_chapter_content", chapterHandler.GetChapterContentById)

	recommendHandler := handler.NewRecommendHandler(pool.DefaultContent().Gorm)

	g.POST("/get_recommend", recommendHandler.GetRecommend)
	rankHandler := handler.NewRankHandler(pool.DefaultContent().Gorm)

	g.POST("/get_rank", rankHandler.GetRank)
	r.Run(":" + config.Config.Server.Port)
}
