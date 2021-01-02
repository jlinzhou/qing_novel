# 基于golang+vue搭建小说书城---青书网

## 功能介绍

```
1：已经爬取了笔趣阁[http://www.xbiquge.la/]所有小说的链接，当请求阅读内容的时候再爬取章节。
2：首页，分类，最近阅读的页面展示
3：整合了小说搜索引擎[https://github.com/DemonFengYuXiang/YourNovel]这个项目.
```

## 部分页面展示


![avatar](https://novel-zjl2.oss-cn-beijing.aliyuncs.com/images/QQ20210102-165059%402x.png?Expires=1609592030&OSSAccessKeyId=TMP.3Kj6BfLVMDFF6C7s3v3sppC8m6zyRcXzoU2CDYZXeRaNWCW3UqhWuXXT87Ymg925TnLr8VZwkShYsQESaBzLr4c6TRVza3&Signature=fQ9M%2BesdE3pamsazmuJMNjdYhbY%3D)


<br/>
<br/>

## 环境配置

**后端两个进程，书城进程和小说搜索引擎进程。**

**前端基于vue，得使用node**

#### 1. 书城进程

* 需要mysql数据库，运行data/book.sql,创建并插入初始数据。[只有一本小说数据(牧神记),全部小说数据下载链接：[data.sql]https://novel-zjl2.oss-cn-beijing.aliyuncs.com/images/book2.sql?Expires=1609591883&OSSAccessKeyId=TMP.3Kj6BfLVMDFF6C7s3v3sppC8m6zyRcXzoU2CDYZXeRaNWCW3UqhWuXXT87Ymg925TnLr8VZwkShYsQESaBzLr4c6TRVza3&Signature=0Gxijvy46qX%2F2OqpgZZKJG6cMTo%3D]
* config/config.yaml 的数据库参数需要配置
* ```go run main.go```

#### 2. 小说搜索引擎进程

* 需要redis，请配置YourNovel/yournovel/conf/server.go 中的redis配置
* ```cd YourNovel```
* ```go run main.go```

#### 3.前端

```sh
cd frontend/nove_mb
cnpm install
yarn run serve

```



## 最后
**有问题欢迎留言，觉得有用的话点个star，不甚感激**