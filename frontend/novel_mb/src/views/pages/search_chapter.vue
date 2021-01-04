<!--  -->
<template>
  <div>

    <div class="box_con">
      <div class="bread">
        <el-breadcrumb separator-class="el-icon-arrow-right" v-show="!listLoadingChapter">
          <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
          <el-breadcrumb-item>{{novel_name}}</el-breadcrumb-item>
        </el-breadcrumb>
      </div>
      <div id="sidebar" v-show="!this.listLoadingChapter">
        <div id="fmimg"><img alt="" src="@/assets/images/no_cover.jpg" width="120" height="150">

        </div>
      </div>
      <div id="maininfo" v-show="!this.listLoadingChapter">
        <div id="info">
          <h1>{{novel_name}}</h1>
          <!-- <p>作&nbsp;&nbsp;&nbsp;&nbsp;者：{{bookInfo.list.author}}</p>
          <p>最后更新：{{bookInfo.last_time|formateTime}}</p>
          <p>最新章节：<a href="javascript:void(0);" @click="chapterClick(bookInfo.last_chapter_id)">{{bookInfo.last_chapter_name}}</a></p> -->
        </div>

      </div>
      <div id="intro" v-show="!this.listLoadingChapter">
        <p style="text-align: center">
          <strong>注意！</strong>页面内容来自<a :href="origin_url">{{ this.origin_url }}</a>，本站不储存任何内容，为了更好的阅读体验进行在线解析，若有广告出现，请及时反馈。若您觉得侵犯了您的利益，请通知我们进行删除，然后访问
          <a :href="origin_url">原网页</a>
        </p>
        <!-- <p> {{bookInfo.list.intro}}</p> -->
      </div>
      <!-- <div class="backTop">
        <el-button type="warning" size="mini" plain @click="backTop">返回顶部</el-button>
      </div> -->
      <!-- <div class="goBottom">
        <el-button type="warning" size="mini" plain @click="goBottom">去到底部</el-button>
      </div> -->

      <!-- max-height="1000" -->
      <div id="chapter">
        <div class="title">正文</div>
        <div class="chapter_page" v-show="!this.listLoadingChapter">
          <el-table :data="tableData" ref="myTable" max-height="1000" style="width: 100%" :row-style="{height:'15px'}" :cell-style="{padding:'0 0'}" :show-header="false" id="TableDiv">

            <el-table-column>
              <template slot-scope="scope">
                <el-link type="primary" @click="searchChapterClick(scope.row.c1.url)">{{scope.row.c1.name}}</el-link>
              </template>
            </el-table-column>
            <el-table-column>
              <template slot-scope="scope">
                <el-link type="primary" @click="searchChapterClick(scope.row.c2.url)">{{scope.row.c2.name}}</el-link>
              </template>
            </el-table-column>
            <!-- <el-table-column>
              <template slot-scope="scope">
                <el-link type="primary" @click="chapterClick(scope.row.c3.id)">{{scope.row.c3.name}}</el-link>
              </template>
            </el-table-column> -->
          </el-table>
        </div>
        <!-- <pagination v-show="total > 0" :total="total" :page.sync="page" :limit.sync="page_size" @pagination="getChapterInfo" /> -->

      </div>
    </div>

  </div>
</template>

<script>
import { GetChapter } from '@/api/search'

export default {
  // components: { pagination },

  methods: {
    getChapter() {
      const load = this.$loading({
        lock: true,
        text: '请耐心等待...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      });
      GetChapter({ url: this.novel_url, novel_name: this.novel_name }).then(response => {

        this.chapterInfo = response.data.chapter.chapter_info
        let link_prefix = response.data.chapter.link_prefix
        this.origin_url = response.data.chapter.origin_url
        if (link_prefix == '-1') {
          this.prefix = response.data.chapter.origin_url;
        } else if (link_prefix == '1') {
          this.prefix = '';
        } else if (link_prefix == '0') {
          this.prefix = response.data.chapter.domain;
        }

        this.loadData()
        this.listLoadingChapter = false
        load.close();
      })
    },
    searchChapterClick(url) {
      this.$router.replace('/global_search/search_reader/' + encodeURIComponent(url) + "/" + encodeURIComponent(this.novel_url) + "/" + encodeURIComponent(this.novel_name))
    },

    loadData() {
      this.tableData = []
      let maxGroup = Math.floor(this.chapterInfo.length / 2)
      let otherDataNums = this.chapterInfo.length - maxGroup * 2

      let data = this.chapterInfo.splice(0, maxGroup * 2)
      let precolumns = {}
      for (let i = 0; i < data.length; i++) {

        let temp = { url: this.prefix + '/' + data[i].chapter_url, name: data[i].chapter_name }
        if (i % 2 === 0) {
          precolumns = {}
          precolumns["c1"] = temp
        }
        if (i % 2 === 1) {
          precolumns["c2"] = temp
          this.tableData.push(precolumns);
        }


      }

      if (otherDataNums === 1) {
        let columns = { url: this.prefix + '/' + this.chapterInfo[this.chapterInfo.length - 1].chapter_url, name: this.chapterInfo[this.chapterInfo.length - 1].chapter_name }
        this.tableData.push({ c1: columns, c2: { url: '', name: '' } });

      }

    },
    backTop() {
      this.$refs.myTable.bodyWrapper.scrollTop = 0;

    },
    goBottom() {
      this.$refs.myTable.bodyWrapper.scrollTop = 1000;
    },

  },
  data() {
    return {
      novel_url: '',
      novel_name: '',
      prefix: '',
      origin_url: '',
      // content_title: '',
      // chapter_url: '',
      chapterInfo: [],
      tableData: [],
      listLoadingChapter: true

    };
  },


  mounted() {
    this.novel_url = this.$route.params.novel_url
    this.novel_name = this.$route.params.novel_name

    this.getChapter()
  },

}
</script>
<style   lang="scss" scoped>
.box_con {
  // border: 4px solid #aab8d8;
  overflow: hidden;
  // width: 976px;
  margin: 10px auto;
}

.bread {
  width: 100%;

  // height: 30px;
  display: inline-block;
  overflow: hidden;
  height: 20px;
  font-size: 20px;
  float: left;
  background: #cfe1ea;
}
.el-breadcrumb {
  margin-top: 2px;
}

#maininfo {
  display: inline-block;
  float: right;
  width: 60%;
  // width: 1000px;
}
// a {
//   float: left;
// }
.con_top #bdshare {
  float: right;
  line-height: 20px;
  padding-right: 20px;
  padding-top: 9px;
  text-align: right;
}
#bdshare img {
  border: 0;
  margin: 0;
  padding: 0;
  cursor: pointer;
}
#bdshare a,
#bdshare_s a,
#bdshare_pop a {
  text-decoration: none;
  cursor: pointer;
}
#info {
  padding: 10px;
  margin: 10px;
  font-size: 14px;
}
#info h1 {
  // font-family: ºÚÌå;
  font-size: 16px;
  font-weight: 700;
  overflow: hidden;
  // margin: auto;
  // float: left;
  padding: 1px;
  margin-top: 5px;
  margin-left: 20px;

  text-align: left;
}
h1 {
  display: block;
  font-size: 16px;
  margin-block-start: 0.67em;
  margin-block-end: 0.67em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  font-weight: bold;
}
#info p {
  height: 17px;
  line-height: 17px;
  padding-top: 2px;
  // width: 550px;
  margin-left: 20px;
  overflow: hidden;
  display: block;
  text-align: left;

  // float: left;
}
#intro {
  width: 96%;
  overflow: hidden;
  // float: left;
  line-height: 150%;
  border-top: 1px dashed #88c6e5;
  padding: 10px;
  font-size: 13px;
  text-align: left;
}
#intro p {
  text-indent: 2em;
  margin-top: 10px;
  text-align: left;
}
#sidebar {
  float: left;
  width: 140px;
  text-align: left;
}
#fmimg {
  background-color: #e1eced;
  float: left;
  width: 126px;
  margin: 12px;
  padding: 12px;
  position: relative;
}
#fmimg img {
  border: medium none;
  height: 150px;
  width: 120px;
  margin: 3px;
}
#fmimg .b {
  background-position: 0 -294px;
}
#fmimg span {
  top: 8px;
  right: 8px;
  width: 88px;
  height: 88px;
  position: absolute;
  display: block;
  z-index: 999;
}
#chapter {
  border: 4px solid #aab8d8;

  overflow: auto;
  // background-color: red !important;
}
.backTop {
  position: fixed;
  margin-left: 75%;
  z-index: 9999999;
}
.goBottom {
  position: fixed;
  margin-left: 75%;
  margin-top: 33px;
  z-index: 9999999;
}

// .box_con {
//   border: 4px solid #aab8d8;
//   overflow: hidden;
//   // width: 976px;
//   margin: 10px auto;
// }
// .bread {
//   width: 100%;

//   // height: 30px;
//   display: inline-block;
//   overflow: hidden;
//   height: 20px;
//   font-size: 20px;
//   float: left;
//   background: #cfe1ea;
// }
// .el-breadcrumb {
//   margin-top: 2px;
// }
// #maininfo {
//   float: right;
//   width: 1000px;
// }
// // a {
// //   float: left;
// // }
// .con_top #bdshare {
//   float: right;
//   line-height: 20px;
//   padding-right: 20px;
//   padding-top: 9px;
//   text-align: right;
// }
// #bdshare img {
//   border: 0;
//   margin: 0;
//   padding: 0;
//   cursor: pointer;
// }
// #bdshare a,
// #bdshare_s a,
// #bdshare_pop a {
//   text-decoration: none;
//   cursor: pointer;
// }
// #info {
//   padding: 10px;
//   margin: 10px;
//   font-size: 15px;
// }
// #info h1 {
//   font-family: ºÚÌå;
//   font-size: 28px;
//   font-weight: 700;
//   overflow: hidden;
//   // margin: auto;
//   // float: left;
//   padding: 1px;
//   margin-top: 5px;
//   margin-left: 20px;

//   text-align: left;
// }
// h1 {
//   display: block;
//   font-size: 2em;
//   margin-block-start: 0.67em;
//   margin-block-end: 0.67em;
//   margin-inline-start: 0px;
//   margin-inline-end: 0px;
//   font-weight: bold;
// }
// #info p {
//   height: 25px;
//   line-height: 25px;
//   padding-top: 2px;
//   width: 550px;
//   margin-left: 20px;
//   overflow: hidden;
//   display: block;
//   text-align: left;

//   // float: left;
// }
// #intro {
//   width: 96%;
//   overflow: hidden;
//   line-height: 150%;
//   border-top: 1px dashed #88c6e5;
//   padding: 10px;
//   font-size: 13px;
// }
// #intro p {
//   text-indent: 2em;
//   margin-top: 10px;
// }
// #sidebar {
//   float: left;
//   width: 140px;
//   text-align: left;
// }
// #fmimg {
//   background-color: #e1eced;
//   float: left;
//   width: 126px;
//   margin: 12px;
//   padding: 12px;
//   position: relative;
// }
// #fmimg img {
//   border: medium none;
//   height: 150px;
//   width: 120px;
//   margin: 3px;
// }
// #fmimg .b {
//   background-position: 0 -294px;
// }
// #fmimg span {
//   top: 8px;
//   right: 8px;
//   width: 88px;
//   height: 88px;
//   position: absolute;
//   display: block;
//   z-index: 999;
// }
// #chapter {
//   border: 4px solid #aab8d8;

//   overflow: auto;
//   // background-color: red !important;
// }
// #TableDiv {
//   background-color: red !important;
// }
// .el-table td {
//   background: red !important;
// }
// .el-table th {
//   background: red !important;
// }
// .el-table .success-row {
//   background: red !important;
// }
</style>