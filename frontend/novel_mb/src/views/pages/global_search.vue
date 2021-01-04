<!--  -->
<template>
  <div id="search">
    <div id="header">
      <el-input class="search-input" v-model="currSearchKey" placeholder="请输入小说名"></el-input>

      <el-button type="primary" icon="el-icon-search" @click="getSearchData">搜索</el-button>

    </div>
    <div class="read_book">
      <span class="font-tip">简介：这是一个基于其他搜索引擎构建的垂直小说搜索引擎，可以搜索到全网的免费小说。</span>
      <span class="font-tip"> 备注:第一次搜索比较慢，需要等待10秒后重新刷新再尝试一下。</span>

      <p style="color: #868383;font-size: 13px; ">找到 {{ this.count }} 条结果（用时{{ this.elapsedTime }}s）</p>
      <div class="list-unstyled item">

        <div class="result-item" v-for="(item ,index) in bookList" :key="index">
          <li>
            <a href="javascript:void(0);" @click="chapterClick(item.href)">{{ item.title }}</a>
            <div class="netloc">
              <i>{{ item.host }}</i>
              <i><a href="javascript:void(0);" @click="clickHost(item.href)">查看源网址</a></i>
              <span class="label-primary" v-if="item.is_parse">已解析</span>
              <span class="label-danger" v-if="!item.is_parse">未解析</span>
            </div>
            <div class="tags">

            </div>
          </li>

        </div>
      </div>

    </div>

  </div>
</template>

<script>
import { GetSearch } from '@/api/search'

export default {
  components: {},
  methods: {
    chapterClick(href) {
      this.$router.replace('/global_search/search_chapter/' + encodeURIComponent(href) + '/' + encodeURIComponent(this.currSearchKey))
    },
    clickHost(href) {
      window.location.href = href
    },
    getSearchData() {
      const load = this.$loading({
        lock: true,
        text: '第一次搜索比较慢...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      });
      this.listLoading = true
      if (this.currSearchKey != '') {
        localStorage.setItem('global_search', this.currSearchKey);

        GetSearch({ wd: this.currSearchKey }).then(response => {
          this.count = response.data.count
          this.elapsedTime = response.data.elapsedTime
          this.bookList = response.data.list
          this.listLoading = false
          load.close();
        })
      }

    }
  },
  // localStorage.getItem('latest_view')
  data() {
    return {
      listLoading: true,
      count: 0,
      elapsedTime: 0,
      currSearchKey: '',
      bookList: [],
    };
  },
  mounted() {
    this.currSearchKey = localStorage.getItem('global_search')
  }
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
.search-input {
  width: 60%;
}
.font-tip {
  display: block;
  color: #0000008f;
  font-size: 10px;
  text-align: left;
  margin-left: 0px;
}
.result-item {
  margin-left: 20%;
  width: 70%;
  height: 75px;
  margin-top: 20px;
  text-align: left;
}
li {
  list-style-type: none;
}
.netloc {
  display: block;
  color: #006621;
}
.label-danger {
  background-color: #d9534f;
  display: inline;
  padding: 0.2em 0.6em 0.3em;
  font-size: 75%;
  font-weight: 700;
  line-height: 1;
  color: #fff;
  text-align: center;
  white-space: nowrap;
  vertical-align: baseline;
  border-radius: 0.25em;
}
.label-primary {
  background-color: #337ab7;
  display: inline;
  padding: 0.2em 0.6em 0.3em;
  font-size: 75%;
  font-weight: 700;
  line-height: 1;
  color: #fff;
  text-align: center;
  white-space: nowrap;
  vertical-align: baseline;
  border-radius: 0.25em;
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
  color: #337ab7;
  text-decoration: none;
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