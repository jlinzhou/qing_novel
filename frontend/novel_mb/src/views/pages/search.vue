<!--  -->
<template>
  <div id="search">
    <div id="header">
      <span>搜索“{{this.$store.getters.searchKey}}”结果</span>
    </div>
    <div class="read_book">
      <div class="search-tip" v-if="showSearch" @click="clickSearch">搜索不到，试试全网搜索</div>
      <div class="block" v-for="(item ,index) in bookList" :key="index">
        <div class="block_img"><a href="javascript:void(0);" @click="bookClick(item.id)"><img height="100" width="80" :src="item.cover"></a></div>
        <div class="block_txt">
          <p><a href="javascript:void(0);" @click="bookClick(item.id)"></a></p>
          <h2><a a href="javascript:void(0);" @click="bookClick(item.id)">{{item.title}}</a></h2>
          <p></p>
          <p>分类：{{item.category|filterCategory}}小说</p>
          <p>作者：{{item.author}}</p>
          <p>最新章节：<a a href="javascript:void(0);" @click="chapterClick(item.last_chapter_id)">{{item.last_chapter_name}}</a></p>
        </div>
        <div style="clear:both"></div>
      </div>

    </div>
    <!-- 
      <div id="header">
      <span>搜索“{{this.$store.getters.searchKey}}”结果</span>
    </div>
    <div>
      <el-table :data="bookList" style="width: 100%">
        <el-table-column label="文章名称">
          <template slot-scope="scope">
            <el-link type="primary" @click="bookClick(scope.row.id)">{{scope.row.title}}</el-link>
          </template>
        </el-table-column>
        <el-table-column label="最新章节">
          <template slot-scope="scope">
            <el-link type="primary" @click="chapterClick(scope.row.id,scope.row.last_chapter_id)">{{scope.row.last_chapter_name}}</el-link>
          </template>
        </el-table-column>

        <el-table-column prop="author" label="作者">
        </el-table-column>
        <el-table-column prop="mtime" label="更新">
          <template slot-scope="scope">
            {{scope.row.mtime|formateTime}}
          </template>
        </el-table-column>
      </el-table>
    </div> -->

  </div>
</template>

<script>
import { GetBookByName } from '@/api/book'

export default {
  components: {},
  methods: {
    bookClick(id) {
      this.$router.replace('/chapter/' + id)
    },
    chapterClick(book_id, chapter_id) {
      this.$router.replace('/reader/' + book_id + "/" + chapter_id)
    },
    clickSearch() {
      this.$router.replace('/global_search')
    },
    getSearchData() {
      this.showSearch = false
      if (this.$store.getters.searchKey != '') {
        GetBookByName({ name: this.$store.getters.searchKey }).then(response => {
          this.bookList = response.data.list
          console.log("this.bookList.length:", this.bookList.length)
          if (this.bookList.length === 0) {
            this.showSearch = true
          }
          console.log("this.showSearch:", this.showSearch)
          this.listLoading = false

        })
      }

    }
  },
  data() {
    return {
      listLoading: true,
      showSearch: false,
      bookList: [],
    };
  },
  watch: {
    "$store.getters.searchKey": {
      handler: function (val) {
        this.getSearchData()
      },
      deep: true
    }
  },
  mounted() {
    this.getSearchData()

  },
}
</script>
<style lang="scss" scoped>
#search {
  margin-top: 0px;
  // border: 4px solid #aab8d8;
}
#header {
  border-bottom: 1px solid rgb(136, 198, 229);
  text-align: center;
  padding: 0px 10px;
  line-height: 40px;
  background-color: rgb(225, 236, 237);
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
a {
  color: #000;
  text-decoration: none;
}
.search-tip {
  color: blue;
}
.block_txt {
  border: 0;
  height: 100px;
  overflow: hidden;
  line-height: 20px;
  padding-left: 10px;
  padding-top: 8px;
}
.block_txt p,
.block_txt2 p {
  text-align: left;
  height: auto;
  line-height: 12px;
  font-size: 12px;
}
.block_txt h2,
.block_txt2 h2 {
  text-align: left;
  font-size: 16px;
  line-height: 14px;
  height: auto;
}
</style>