<!--  -->
<template>
  <div>
    <div class="article">
      <h2 class="title">
        <span>封面推荐</span>
      </h2>
      <div class="block">
        <div v-for="(item,index) in topBook" :key="index" v-show="!listLoadingTop">
          <div class="block_img"><a href="javascript:void(0);" @click="bookClick(item.id)"><img height="100" width="80" :src="item.cover"></a></div>
          <div class="block_txt">
            <p><a href="javascript:void(0);" @click="bookClick(item.id)"></a></p>
            <h2><a href="javascript:void(0);" @click="bookClick(item.id)">{{item.title}}</a></h2>
            <p></p>
            <p>作者：{{item.author}}</p>
            <p><a class="intro" href="javascript:void(0);" @click="bookClick(item.id)"> {{item.intro}}</a></p>
          </div>
          <div style="clear:both"></div>
        </div>

        <!-- <ul>
          <li><a href="/wapbook/951.html" class="blue">伏天氏</a>/净无痕</li>
          <li><a href="/wapbook/26874.html" class="blue">沧元图</a>/我吃西红柿</li>
          <li><a href="/wapbook/14930.html" class="blue">元尊</a>/天蚕土豆</li>
          <li><a href="/wapbook/10.html" class="blue">武炼巅峰</a>/莫默</li>
          <li><a href="/wapbook/27256.html" class="blue">临渊行</a>/宅猪</li>
        </ul> -->
      </div>
      <div class="bottom-book">
        <div class="bottom-book-left">
          <div class="l">
            <h2 class="title">
              <span>最近更新小说列表</span>
            </h2>
            <ul>
              <li v-for="(item,idx) in lastUpdateBook" :key="idx" v-show="!listLoadingUpdate">
                <span class="s1">[{{item.book.category|filterCategory}}]</span>
                <span class="s2"><a href="javascript:void(0);" @click="bookClick(item.book.id)">{{item.book.title}}</a></span>
                <span class="s3"><a href="javascript:void(0);" @click="chapterClick(item.book.id,item.chapter_id)">{{item.chapter_name}}</a></span>
                <!-- <span class="s4">{{item.book.author}}</span> -->
              </li>

            </ul>
          </div>
        </div>
        <div class="bottom-book-right">
          <div class="r">
            <!-- <h2>最新入库小说</h2> -->
            <h2 class="title">
              <span>最新入库小说</span>
            </h2>
            <!-- <span>最新入库小说</span> -->
            <ul>
              <li v-for="(item,idx) in lastCreateBook" :key="idx" v-show="!listLoadingCreate">
                <span class="s1">[{{item.category|filterCategory}}]</span>
                <span class="s2"><a href="javascript:void(0);" @click="bookClick(item.id)">{{item.title}}</a></span>
                <span class="s5">{{item.author}}</span>
              </li>

            </ul>

          </div>
        </div>

      </div>
    </div>

  </div>
</template>

<script>
import { GetRecommend, GetBookByCategory, GetBookByMtime, GetBookByCtime } from '@/api/book'
// import { ChapterClick, BookClick } from '@/utils/common'


export default {
  components: {},

  methods: {
    bookClick(id) {
      this.$router.push({ path: '/chapter/' + id })
    },
    chapterClick(book_id, chapter_id) {
      this.$router.push({ path: '/reader/' + book_id + "/" + chapter_id })
    },
    getTopBook() {
      GetRecommend({ count: 4, type: 1 }).then(response => {
        this.topBook = response.data.list
        this.topBook.forEach(item => {
          if (item.title.length > 6) {
            item.title = item.title.substring(0, 5) + "..."
          }
        });
        // console.log(response)
        // this.alldir = response.data.items
        // this.totaldir = response.data.total
        this.listLoadingTop = false
      })
    },
    // getRecommendBook() {
    //   this.recommendBook = []

    //   GetRecommend({ count: 14, type: 2 }).then(response => {
    //     // this.topBook = response.data.list

    //     for (let i = 0; i < response.data.list.length; i++) {
    //       if (i >= 4) {
    //         this.recommendBook.push(response.data.list[i])
    //       }
    //     }

    //     this.listLoading = false
    //   })
    // },
    getCategoryBook() {
      this.categoryList = []
      for (let i = 1; i <= 6; i++) {
        GetRecommend({ category: i, count: 13, type: 3 }).then(response => {
          this.categoryList.push(response.data.list)
        })
      }
    },
    getLastUpdateBook() {
      GetBookByMtime({ page_size: 10 }).then(response => {
        this.lastUpdateBook = response.data.list
        this.listLoadingUpdate = false
      })
    },
    getLastCreateBook() {
      GetBookByCtime({ page_size: 10 }).then(response => {
        this.lastCreateBook = response.data.list
        this.listLoadingCreate = false
      })
    },

  },
  data() {
    return {
      listLoadingTop: true,
      listLoadingUpdate: true,
      listLoadingCreate: true,
      topBook: [],
      // recommendBook: [],
      categoryList: [],
      lastUpdateBook: [],
      lastCreateBook: []
      // xiuzhenList: [],
      // dushiList: [],
      // chuanyueList: [],
      // wangyouList: [],
      // kehuanList: [],

    };
  },

  mounted() {
    this.getTopBook()
    // this.getRecommendBook()
    this.getCategoryBook()
    this.getLastUpdateBook()
    this.getLastCreateBook()
  }
}
</script>
<style   lang="scss" scoped>
.article {
  margin: 10px auto 10px auto;
}
.title {
  height: 35px;
  background: #ecf0f0;
  border-bottom: 1px solid #007bb1;
  color: #000;
  font-size: 16px;
  padding-left: 10px;
  line-height: 35px;
  font-weight: 400;
  font-weight: 700;
}
ul {
  margin: 0;
  padding: 0;
}
// .title a {
//   float: right;
//   padding-right: 10px;
//   font-size: 14px;
//   font-weight: 400;
// }
.title span {
  text-align: left;
  display: block;
}
.block {
  padding-left: 10px;
  padding-bottom: 10px;
}
.block_img {
  height: auto;
  border: 0;
  overflow: hidden;
  padding-top: 10px;
  padding-bottom: 10px;
  float: left;
}
.block_txt {
  border: 0;
  height: 120px;
  overflow: hidden;
  line-height: 13px;
  font-size: 13px;
  padding-left: 10px;
}
a {
  text-decoration: none;
}
.block_txt a {
  float: left;
  // text-align: left;
  color: #000;
  line-height: 16px;
  text-decoration: none;
}
.block_txt p,
.block_txt2 p {
  height: auto;
  font-size: 14px;
}
p {
  display: block;
  margin-block-start: 1em;
  margin-block-end: 1em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
}
ul {
  display: block;
  list-style-type: disc;
  margin-block-start: 1em;
  margin-block-end: 1em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  // padding-inline-start: 40px;
}
.block li {
  line-height: 35px;

  display: block;
  text-align: left;
  border-bottom: 2px solid #dccece;
}
.block li a {
  color: #5864c3;
  // float: left;
}
.intro {
  text-align: left;
}
.bottom-book {
  width: 100%;
}
.bottom-book-left {
  width: 100%;
}
.bottom-book-right {
  width: 100%;
}
.bottom-book-left .l {
  float: left;
  width: 100%;
}

.bottom-book-left ul {
  display: block;
  list-style-type: disc;
  margin-block-start: 1em;
  margin-block-end: 1em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  // padding-inline-start: 40px;
}
.bottom-book-left .l li {
  padding: 5px 0px 0px 0px;
  border-bottom: solid 1px #dddddd;
  height: 20px;
  line-height: 19px;
  list-style: none outside none;
  display: list-item;
  text-align: -webkit-match-parent;
  overflow: hidden;
}
.bottom-book-left .l li .s1 {
  // width: 75px;
  width: 50px;
  overflow: hidden;
}
.bottom-book-left .l li .s2 {
  // width: 165px;
  width: 100px;
  margin-right: 10px;
  overflow: hidden;
}
.bottom-book-left .l li .s3 {
  // width: 300px;290+325=615
  // width: 85px;
  // overflow: hidden;
  float: right;
  text-align: right;
  overflow: hidden;
}

.bottom-book-left .l li span {
  float: left;
  display: inline-block;
}
.bottom-book-left .l li span a {
  float: left;
}

.bottom-book-right > .r {
  width: 100%;
  // border: 3px solid #88c6e5;
  float: left;
  // width: 695px;
  // background: #e1eced;
}
// .bottom-book-right h2 {
//   display: block;
//   font-size: 1.5em;
//   margin-block-start: 0.83em;
//   margin-block-end: 0.83em;
//   margin-inline-start: 0px;
//   margin-inline-end: 0px;
//   font-weight: bold;
// }
.bottom-book-right ul {
  display: block;
  list-style-type: disc;
  margin-block-start: 1em;
  margin-block-end: 1em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  padding-inline-start: 40px;
}
.bottom-book-right .r li {
  padding: 5px 0px 0px 0px;
  border-bottom: solid 1px #dddddd;
  height: 20px;
  line-height: 19px;
  margin-left: -50px;
  list-style: none outside none;
  display: list-item;
  text-align: -webkit-match-parent;
  overflow: hidden;
}
.bottom-book-right .r li .s1 {
  width: 68px;
}

.bottom-book-right .r li .s2 a {
  float: right;
}
.bottom-book-right .r li span {
  float: left;
  display: inline-block;
}
.bottom-book-right .r li .s5 {
  float: right;
  text-align: right;
}
</style>