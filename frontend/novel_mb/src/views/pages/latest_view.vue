<!--  -->
<template>
  <div>
    <div class="article">
      <h2 class="title">
        <span>最近阅读</span>
      </h2>
      <div class="block">
        <div v-show="lastViewList.length==0">
          暂无最近阅读
        </div>
        <div style="margin-top: 20px">
          <el-radio-group v-model="check_type" size="small">
            <el-radio-button label="1">本站书库历史</el-radio-button>
            <el-radio-button label="2">全网搜历史</el-radio-button>
          </el-radio-group>
        </div>
        <!-- <span class="font-tip">备注:通过全网搜索的这里不会有记录</span> -->
        <div v-for="(item,index) in lastViewList" :key="index" v-show="!listLoadingTop && check_type==='1'">
          <div class="block_img"><a href="javascript:void(0);" @click="bookClick(item.id)"><img height="100" width="80" v-lazy="item.cover"></a></div>
          <div class="block_txt">
            <p><a href="javascript:void(0);" @click="bookClick(item.id)"></a></p>
            <h2><a href="javascript:void(0);" @click="bookClick(item.id)">{{item.title}}</a></h2>
            <!-- <p></p> -->
            <p>作者：{{item.author}}</p>
            <p class="the-time">{{item.timestamp|share_data_time}}</p>
            <p class="my_tag"><span>书签:</span><a href="javascript:void(0);" @click="chapterClick(item.id,item.chapter_id)">{{item.chapter_name}}</a></p>
            <br>
            <p><a class="intro" href="javascript:void(0);" @click="bookClick(item.id)"> {{item.intro}}</a></p>
          </div>
          <div style="clear:both"></div>
        </div>
        <div v-for="(item,index) in global_search_list" :key="index" v-show=" check_type==='2'">
          <div class="block_img"><a href="javascript:void(0);" @click="searchBookClick(item.novel_url,item.novel_name)"><img height="100" width="80" src="@/assets/images/no_cover.jpg"></a></div>
          <div class="block_txt2">
            <!-- <p><a href="javascript:void(0);" @click="searchBookClick(item.id)"></a></p> -->
            <h2><a href="javascript:void(0);" @click="searchBookClick(item.novel_url,item.novel_name)">{{item.novel_name}}</a></h2>
            <!-- <p></p> -->
            <p>来源{{item.novel_url}}</p>
            <p class="the-time2">{{item.timestamp|share_data_time}}</p>
            <p class="my_tag2"><span>书签:</span><a href="javascript:void(0);" @click="searchChapterClick(item.novel_url,item.content_url,item.novel_name)">{{item.content_title}}</a></p>
            <br>
            <!-- <p><a class="intro" href="javascript:void(0);" @click="bookClick(item.id)"> {{item.intro}}</a></p> -->
          </div>
          <div style="clear:both"></div>
        </div>
      </div>

    </div>

  </div>
</template>

<script>
import { GetBookByIdList } from '@/api/book'
// import { ChapterClick, BookClick } from '@/utils/common'
// type chapter_url book_url book_name 
// this.$router.replace('/global_search/search_chapter/' + encodeURIComponent(href) + '/' + encodeURIComponent(this.currSearchKey))
// this.$router.replace('/global_search/search_reader/' + encodeURIComponent(this.chapter_url) + "/" + encodeURIComponent(url) + "/" + encodeURIComponent(this.novel_name) + "/" + "name")

export default {
  components: {},
  filters: {
    share_data_time(timestamp) {
      let str = ''
      let today = new Date();
      let timeold = (today.getTime() - timestamp);
      let secondsold = Math.floor(timeold / 1000);
      let e_daysold = timeold / (24 * 60 * 60 * 1000);

      let daysold = Math.floor(e_daysold);
      if (daysold > 0) {
        str = daysold + "天"
      }
      let e_hrsold = (e_daysold - daysold) * 24;
      let hrsold = Math.floor(e_hrsold);
      if (hrsold > 0) {
        str += hrsold + "小时"
      }
      let e_minsold = (e_hrsold - hrsold) * 60;
      let minsold = Math.floor(e_minsold)
      if (minsold > 0) {
        str += minsold + "分"
      }
      let seconds = Math.floor((e_minsold - minsold) * 60);
      if (seconds > 0) {
        str += seconds + "秒前"
      }
      return str
    },
  },
  methods: {
    bookClick(id) {
      this.$router.push({ path: '/chapter/' + id })
    },
    searchBookClick(novel_url, book_name) {
      this.$router.replace('/global_search/search_chapter/' + encodeURIComponent(novel_url) + '/' + encodeURIComponent(book_name))
    },
    searchChapterClick(novel_url, content_url, novel_name, content_title) {
      this.$router.replace('/global_search/search_reader/' + encodeURIComponent(content_url) + "/" + encodeURIComponent(novel_url) + "/" + encodeURIComponent(novel_name) + "/")
    },
    chapterClick(book_id, chapter_id) {
      this.$router.push({ path: '/reader/' + book_id + "/" + chapter_id })
    },
    getlastestBookList() {
      let data = JSON.parse(localStorage.getItem('latest_view'));
      let id_list = [];
      this.global_search_list = []
      let chapter_dict = {};
      if (data == null) {
        return
      }
      for (let i = 0; i < data.length; i++) {
        if (data[i]['type'] === 2) {
          this.global_search_list.push(data[i])
        } else {
          id_list.push(data[i]['book_id'])
          chapter_dict[data[i]['book_id']] = { 'chapter_name': data[i]['chapter_name'], 'chapter_id': data[i]['chapter_id'], 'timestamp': data[i]['timestamp'] }
        }

      }
      if (id_list.length !== 0) {
        this.getBookByIdList(id_list, chapter_dict)
      }
      // console.log("global_search_list", this.global_search_list)
      if (this.global_search_list.length !== 0) {
        this.global_search_list = this.global_search_list.sort(this.compare('timestamp'))

      }


    },
    compare(property) {
      return function (a, b) {
        var value1 = a[property];
        var value2 = b[property];
        return value2 - value1;
      }
    },

    getBookByIdList(id_list_m, chapter_dict) {
      this.lastViewList = []
      GetBookByIdList({ id_list: id_list_m }).then(response => {
        let temp = []
        temp = response.data.list
        temp.forEach(item => {
          if (chapter_dict.hasOwnProperty(item['id'])) {
            this.lastViewList.push({ ...item, 'chapter_name': chapter_dict[item['id']]['chapter_name'], 'chapter_id': chapter_dict[item['id']]['chapter_id'], 'timestamp': chapter_dict[item['id']]['timestamp'] })

          }
        });
        // console.log("this.lastViewList :", this.lastViewList)
        this.lastViewList = this.lastViewList.sort(this.compare('timestamp'))
        this.listLoadingTop = false
      })
    },
  },
  data() {
    return {
      listLoadingTop: true,
      global_search_list: [],
      check_type: "1",
      lastViewList: [],

    };
  },

  mounted() {
    this.getlastestBookList()

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
.font-tip {
  display: block;
  color: #0000008f;
  font-size: 10px;
  text-align: left;
  margin-left: 0px;
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
  line-height: 14px;
  font-size: 13px;
  padding-left: 10px;
}
.block_txt span {
  float: left;
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
.block_txt p {
  height: auto;
  font-size: 14px;
}
.block_txt2 {
  border: 0;
  margin-left: 15%;
  float: left;
  height: 120px;
  overflow: hidden;
  line-height: 12px;
  font-size: 13px;
  padding-left: 10px;
}
.block_txt2 p {
  height: auto;
  font-size: 12px;
  text-align: left;
}
p {
  display: block;
  margin-block-start: 1em;
  margin-block-end: 1em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
}
.the-time {
  display: inline-block;
  float: right;
  margin-top: -25px;
  color: #b7317a;
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
.my_tag,
.my_tag a {
  color: rgb(182, 105, 105);
}
</style>