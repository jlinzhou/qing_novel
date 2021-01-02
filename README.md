# 基于golang+vue搭建小说书城---青书网

## 功能介绍

```
1：已经爬取了笔趣阁[http://www.xbiquge.la/]所有小说的链接，当请求阅读内容的时候再爬取章节。
2：首页，分类，最近阅读的页面展示
3：整合了小说搜索引擎[https://github.com/DemonFengYuXiang/YourNovel]这个项目.
```

## 部分页面展示

展示链接:http://49.234.71.174/

<img src="https://novel-zjl2.oss-cn-beijing.aliyuncs.com/images/QQ20210102-165059%402x.png?Expires=1609584086&OSSAccessKeyId=TMP.3Kj6BfLVMDFF6C7s3v3sppC8m6zyRcXzoU2CDYZXeRaNWCW3UqhWuXXT87Ymg925TnLr8VZwkShYsQESaBzLr4c6TRVza3&Signature=p2l8Vsp3fYgMd1f5zU8Pt8%2FVXIY%3D" alt="图片替换文本" width="140" height="300" align="left"  />

<img src="https://novel-zjl2.oss-cn-beijing.aliyuncs.com/images/QQ20210102-165153%402x.png?Expires=1609584403&OSSAccessKeyId=TMP.3Kj6BfLVMDFF6C7s3v3sppC8m6zyRcXzoU2CDYZXeRaNWCW3UqhWuXXT87Ymg925TnLr8VZwkShYsQESaBzLr4c6TRVza3&Signature=ExO0axmRN%2FNv6oLczXqzSNgnmxA%3D" alt="图片替换文本" width="140" height="300" align="left" />

<img src="https://novel-zjl2.oss-cn-beijing.aliyuncs.com/images/QQ20210102-165228%402x.png?Expires=1609584383&OSSAccessKeyId=TMP.3Kj6BfLVMDFF6C7s3v3sppC8m6zyRcXzoU2CDYZXeRaNWCW3UqhWuXXT87Ymg925TnLr8VZwkShYsQESaBzLr4c6TRVza3&Signature=kYJ6rl7QUcFn%2BD3e7aDi5X%2BWymE%3D" alt="图片替换文本" width="140" height="300" align="left" />

<img src="https://novel-zjl2.oss-cn-beijing.aliyuncs.com/images/QQ20210102-165309%402x.png?Expires=1609584116&OSSAccessKeyId=TMP.3KhaJ8vi5rY5SHPLkhbHGwT1x3o7u9EmJoYM5tTZei6Dd4D95Hw5FqNjCLJYUZENWz4Fwha8p4f8KkM7aTHZhXCJ9KKbvP&Signature=Q6wf%2FgcCreJsg01XLWVPJQs2HKE%3D" alt="图片替换文本" width="140" height="300" align="left" />


<br/>


## 环境配置

**后端两个进程，书城进程和小说搜索引擎进程。**

**前端基于vue，得使用node**

#### 1. 书城进程

* 需要mysql数据库，运行data/book.sql,创建并插入初始数据。[只有一本小说数据(牧神记),全本小说数据下载链接：]
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