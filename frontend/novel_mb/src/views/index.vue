<!--  -->
<template>
  <!-- :style="{'background-color': this.$store.getters.bodyBackground}" -->
  <div class="body" :style="{'background-color': this.$store.getters.bodyBackground}">
    <el-container>
      <!-- <el-header class="header" style="height:30px"></el-header> -->

      <el-main class="main">
        <div class="header">
          <span class="header-img"><img src="@/assets/images/favicon.png" alt="青书网" width="38" height="38">

          </span>
          <span class="back-i" @click="back_step">
            <i class="el-icon-caret-left ">返回</i>
          </span>
          <span>
            <b>青书网</b>
          </span>
          <el-button class="gohome" type="success" size="mini" v-if="this.$store.getters.hideHeader" @click="gohome">首页</el-button>
        </div>

        <div class="search-main" v-if="!this.$store.getters.hideHeader">
          <div class="search-input">
            <el-input placeholder="请输入小说名/作者" v-model="searchKey" @change="searchData"></el-input>
            <el-button @click="searchData">搜索</el-button>
          </div>

        </div>
        <div class="menu" v-if="!this.$store.getters.hideHeader">
          <el-menu class="el-menu" router style="min-height:100%" mode="horizontal" @select="handleSelect">
            <el-menu-item index="/home" :class="[activeIndex===1?'active':'','link-category-index']">
              <span class="g-span">首页</span>
            </el-menu-item>
            <el-menu-item index="/category/xuanhuan" :class="[activeIndex===2?'active':'','link-category-index']">
              <span class="g-span"> 分类</span>
            </el-menu-item>
            <el-menu-item index="/rank" :class="[activeIndex===3?'active':'','link-category-index']">
              <span class="g-span">排行榜</span>
            </el-menu-item>
            <el-menu-item index="/latest_view" :class="[activeIndex===4?'active':'','link-category-index']">
              <span class="g-span">最近读</span>
            </el-menu-item>
            <el-menu-item index="/global_search" :class="[activeIndex===5?'active':'','link-category-index']">
              <span class="g-span">全网搜</span>
            </el-menu-item>
          </el-menu>
        </div>
        <div class="index-middle">
          <router-view></router-view>
        </div>
      </el-main>
      <el-footer>
        <div class="footer">

          <div class="foot-tips">
            <br>
            <span>&nbsp;&nbsp;&nbsp;&nbsp;</span>
            <div>本站所有小说为转载作品，所有章节均由网友上传，转载至本站只是为了宣传本书让更多读者欣赏。</div>

          </div>
          <div class="foot-div">

            <el-button class="foot-home" @click="gohome"> 首页</el-button>
            <!-- <a href="./static/file/上传文件字段要求.docx" download="下载时文件名.docx">点击显示名</a>
 -->

          </div>
          <!-- <div>
                Copyright © 2015-2022 小说网 All Rights Reserved.
              </div> -->
          <div>

          </div>

        </div>

      </el-footer>
      <!-- <div class="middle">
        <div class="index-middle">
          <router-view></router-view>
        </div>
      </div> -->

    </el-container>

  </div>
</template>

<script>

export default {
  data() {
    return {
      activeIndex: 0,
      searchKey: '',
    };
  },
  watch: {
    $route: {
      handler: function (val) {
        if (this.$route.path.indexOf("reader") === -1) {
          this.$store.dispatch("book/saveBbg", '#f0f9eb')
        }
        if (this.$route.path.indexOf("/home") !== -1) {
          this.activeIndex = 1
        }
        if (this.$route.path.indexOf("/category") !== -1) {
          this.activeIndex = 2
        }
        if (this.$route.path.indexOf("/rank") !== -1) {
          this.activeIndex = 3
        }
        if (this.$route.path.indexOf("/latest_view") !== -1) {
          this.activeIndex = 4
        }
        if (this.$route.path.indexOf("/global_search") !== -1) {
          this.activeIndex = 5
        }

      },
      deep: true
    }
  },
  methods: {
    handleSelect(index) {
      this.activeIndex = index
    },
    back_step() {
      this.$router.back(-1)
    },
    filterCategory(code) {
      let empty = '';
      this.categorySelect.forEach((item) => {
        if (item.id === code) {
          empty = item.name;
        }
      });
      return empty;
    },
    gohome() {
      this.$router.replace('/home')
    },


    searchData() {
      if (this.searchKey != '') {
        this.$store.dispatch("book/saveSearch", this.searchKey)
        this.$router.replace('/search')
        this.searchKey = ''
      }
    }
  }

}
</script>
<style  lang="scss" >
.body {
  background-color: #f0f9eb;
  margin: 0;
  width: 100%;
  height: 100%;
}
.main {
  padding: 0px;
  /* width: 1200px;
  height: 900px;
  margin: auto;
  margin-top: -33px; */
}

.header {
  width: 100%;
  background: #d4deae;
  display: block;
  position: relative;
  font-size: 16px;
  height: 50px;
  line-height: 50px;
}
.header-img {
  float: left;
  margin-left: 5px;
  margin-top: 5px;
}
.back-i {
  display: inline-block;
  float: left;
  color: #776223;
  font-size: 18px;
}
.search-input {
  margin-top: 5px;
  width: 100%;
  /* margin-left: -70px; */
}
.search-input > .el-input {
  width: 60%;
  margin-left: 5px;
}
.search-input > .el-button {
  margin: 2px;
  border: 1px solid black;
}
/* .el-menu-item.is-active {
  background-color: rgb(231, 235, 240) !important;
} */
.main {
  width: 100%;
  height: 900px;
  /* margin: auto; */
  /* margin-top: 10px; */
}
.el-menu {
  width: 100%;
  height: 40px;
  background-color: #f3f9f5e3;
}
.el-menu-item {
  width: 12%;
  height: 40px;
  line-height: 40px;
  margin: auto;
  /* font-size: 40%; */
  text-align: center;
  /* padding-bottom: 10px; */
}
.el-menu--horizontal > .el-menu-item {
  height: 40px;
  line-height: 40px;
}
.gohome {
  display: inline-block;
  float: right;
  margin-top: 10px;
  margin-right: 10px;
}
.el-footer {
  padding: 0 0px;
}
.footer {
  font-size: 12px;
}
.foot-div {
  width: 100%;

  background-color: white;
}
.foot-home {
  float: left;
  margin-left: 15%;
  font-size: 8px;
}
.foot-pc {
  // float: right;
  font-size: 8px;
  // margin-right: -234px;
}
.foot-tips {
  background-color: white;
  width: 100%;
  margin-top: -8px;
  text-align: left;
  height: 60px;
}
.foot-tips div {
  // margin-top: 7px;
}
.link-category-index {
  width: 19%;
  // height: 20px;
  font-size: 18px;

  font-family: monospace;
}
/* .body {
  background-color: #f0f9eb;
  width: 100%;
  height: 100%;
}
.header {
  background-color: #e4e7ed;
  font-size: 20px;
  line-height: 20px;
}
.main {
  width: 1200px;
  height: 900px;
  margin: auto;
  margin-top: -33px;
}
.middle {
  width: 1200px;
  height: 900px;
  margin: auto;
}
.search-main {
  width: 1200px;
  height: 80px;
  margin: auto;
}


.g-span {
  color: brown;
}
.active {
  background-color: rgb(231, 235, 240) !important;
}
.link-category {
  width: 165px;
  font-size: 17px;
  font-family: monospace;
}

.footer {
  width: 750px;
  margin: auto;
  border-top: 1px solid black;
}
.index-middle {
  overflow: auto;
} */
</style>