<!--  -->
<template>
  <div>

    <div class="box_con">
      <div class="bread">
        <el-breadcrumb separator-class="el-icon-arrow-right" v-show="!listLoadingBook">
          <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
          <el-breadcrumb-item>{{bookInfo.list.category|filterCategory}}小说</el-breadcrumb-item>
          <el-breadcrumb-item>{{bookInfo.list.title}}</el-breadcrumb-item>
        </el-breadcrumb>
      </div>
      <div id="sidebar" v-show="!this.listLoadingBook">
        <div id="fmimg"><img alt="" :src="bookInfo.list.cover" width="120" height="150">

        </div>
      </div>
      <div id="maininfo" v-show="!this.listLoadingBook">
        <div id="info">
          <h1>{{bookInfo.list.title}}</h1>
          <p>作&nbsp;&nbsp;&nbsp;&nbsp;者：{{bookInfo.list.author}}</p>
          <p>最后更新：{{bookInfo.last_time|formateTime}}</p>
          <p>最新章节：<a href="javascript:void(0);" @click="chapterClick(bookInfo.last_chapter_id)">{{bookInfo.last_chapter_name}}</a></p>
        </div>

      </div>
      <div id="intro" v-show="!this.listLoadingBook">
        <p> {{bookInfo.list.intro}}</p>
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
                <el-link type="primary" @click="chapterClick(scope.row.c1.id)">{{scope.row.c1.name}}</el-link>
              </template>
            </el-table-column>
            <el-table-column>
              <template slot-scope="scope">
                <el-link type="primary" @click="chapterClick(scope.row.c2.id)">{{scope.row.c2.name}}</el-link>
              </template>
            </el-table-column>
            <!-- <el-table-column>
              <template slot-scope="scope">
                <el-link type="primary" @click="chapterClick(scope.row.c3.id)">{{scope.row.c3.name}}</el-link>
              </template>
            </el-table-column> -->
          </el-table>
        </div>
        <pagination v-show="total > 0" :total="total" :page.sync="page" :limit.sync="page_size" @pagination="getChapterInfo" />

      </div>
    </div>

  </div>
</template>

<script>
import { GetChapterInfo } from '@/api/chapter'
import { GetBookById } from '@/api/book'
import pagination from '@/components/Pagination'
export default {
  components: { pagination },
  methods: {
    getBookInfo() {
      GetBookById({ id: this.book_id }).then(response => {
        this.bookInfo = response.data
        this.listLoadingBook = false
      })

    },
    getChapterInfo() {
      GetChapterInfo({ book_id: this.book_id, page: this.page, page_size: this.page_size, sort: this.sort }).then(response => {
        this.chapterInfo = response.data.list
        this.total = response.data.count
        this.loadData()
        this.listLoadingChapter = false
      })
    },

    chapterClick(chapter_id) {
      this.$router.replace('/reader/' + this.book_id + "/" + chapter_id)
    },
    backTop() {
      this.$refs.myTable.bodyWrapper.scrollTop = 0;

    },
    goBottom() {
      this.$refs.myTable.bodyWrapper.scrollTop = 1000;
    },



    loadData() {
      this.tableData = []
      let maxGroup = Math.floor(this.chapterInfo.length / 2)
      let otherDataNums = this.chapterInfo.length - maxGroup * 2

      let data = this.chapterInfo.splice(0, maxGroup * 2)
      let precolumns = {}
      for (let i = 0; i < data.length; i++) {

        let temp = { id: data[i].id, name: data[i].name }
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
        let columns = { id: this.chapterInfo[this.chapterInfo.length - 1].id, name: this.chapterInfo[this.chapterInfo.length - 1].name }
        this.tableData.push({ c1: columns, c2: { id: 0, name: '' } });

      }

    },

  },
  watch: {
    $route: {
      handler: function (val) {
        if (this.$route.path.indexOf("chapter") !== -1) {
          this.book_id = Number(this.$route.params.book_id)
          this.page = Number(this.$route.params.page) || 1
          this.page_size = Number(this.$route.params.page_size) || 200
          this.sort = Number(this.$route.params.sort) || 1
          this.getBookInfo()
          this.getChapterInfo()
        }

      },
      deep: true
    }
  },
  data() {
    return {
      book_id: 0,
      bookInfo: {},
      chapterInfo: [],
      page: 1,
      sort: 1,
      total: 0,
      page_size: 200,
      tableData: [],
      // currentPage: 1,
      // pageSize: 333,
      // totalPage: 0,
      listLoading: true,
      listLoadingChapter: true
      // $refs: {
      //   myTable: 'any'
      // }
    };
  },




  mounted() {
    this.book_id = Number(this.$route.params.book_id)
    this.page = Number(this.$route.params.page) || 1
    this.page_size = Number(this.$route.params.page_size) || 200
    this.sort = Number(this.$route.params.sort) || 1
    this.getBookInfo()
    this.getChapterInfo()
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