<!--  -->
<template>
  <div>
    <div class="article">
      <h2 class="title">
        <div class="title_cate">
          <span>点击排行榜</span>
          <el-radio-group class="radio" v-model="type" size="small" @change="getRank()">
            <el-radio-button label="11">周榜</el-radio-button>
            <el-radio-button label="12">月榜</el-radio-button>
            <el-radio-button label="13">总榜</el-radio-button>
          </el-radio-group>
        </div>
        <div class="title_cate">
          <span>推荐排行榜</span>
          <el-radio-group class="radio" v-model="type" size="small" @change="getRank()">
            <el-radio-button label="21">周榜</el-radio-button>
            <el-radio-button label="22">月榜</el-radio-button>
            <el-radio-button label="23">总榜</el-radio-button>
          </el-radio-group>
        </div>
        <div class="title_cate">
          <span>完本排行榜</span>
          <el-radio-group class="radio" v-model="type" size="small" @change="getRank()">
            <el-radio-button label="31">周榜</el-radio-button>
            <el-radio-button label="32">月榜</el-radio-button>
            <el-radio-button label="33">总榜</el-radio-button>
          </el-radio-group>
        </div>
      </h2>
      <div class="block">
        <div v-for="(item,index) in rankList" :key="index" v-show="!listLoadingTop">

          <div class="block_rank_id">
            <el-button type="warning" circle>{{item.weight}}</el-button>
          </div>
          <!-- :src="item.cover" -->
          <div class="block_img"><a href="javascript:void(0);" @click="bookClick(item.book_id)"><img height="100" width="80" v-lazy="item.cover"></a></div>
          <div class="block_txt">
            <p><a href="javascript:void(0);" @click="bookClick(item.book_id)"></a></p>
            <h2><a href="javascript:void(0);" @click="bookClick(item.book_id)"><span class="s1">[{{item.category|filterCategory}}]</span>{{item.title}}</a></h2>
            <p></p>
            <p>作者：{{item.author}}</p>
            <p><a class="intro" href="javascript:void(0);" @click="bookClick(item.book_id)"> {{item.intro}}</a></p>
          </div>
          <div style="clear:both"></div>
        </div>
      </div>

    </div>

  </div>
</template>

<script>
import { GetRank } from '@/api/book'
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
    getRank() {
      this.listLoadingTop = true
      GetRank({ type: Number(this.type) }).then(response => {
        this.rankList = response.data.list
        this.listLoadingTop = false
      })
    },
  },
  data() {
    return {
      listLoadingTop: true,
      rankList: [],
      type: 11,

    };
  },
  watch: {
    $route: {
      handler: function (val) {
        if (this.$route.path.indexOf("rank") !== -1) {
          this.type = Number(this.$route.params.type) || 11
          this.getRank()

        }

      },
      deep: true
    }
  },

  mounted() {
    this.type = Number(this.$route.params.type) || 11
    this.getRank()

  }
}
</script>
<style   lang="scss" scoped>
.article {
  margin: 10px auto 10px auto;
}
.title {
  height: 110px;
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
.title span,
.title .radio {
  float: left;
  text-align: left;
  display: inline-block;
}
.radio {
  margin-left: 15px;
}
.title_cate {
  height: 35px;
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
.block_rank_id {
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
.block_txt p,
.block_txt2 p {
  height: auto;
  font-size: 14px;
}
.el-button.is-circle {
  border-radius: 50%;
  padding: 6px;
  width: 1.5rem;
  height: 1.5rem;
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
.my_tag,
.my_tag a {
  color: rgb(182, 105, 105);
}
</style>