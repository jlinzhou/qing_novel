<template>
  <div>
    <div id="header">
      <div class="select">
        <el-select v-model="fontsValue" placeholder="字体" size="mini" @change="setFontsValue()">
          <el-option v-for="(item,key) in fontsOptions" :key="key" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
        <el-select v-model="colorValue" placeholder="颜色" size="mini" @change="setColorValue()">
          <el-option v-for="(item,key) in colorOptions" :key="key" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
        <el-select v-model="sizeValue" placeholder="大小" size="mini" @change="setSizeValue()">
          <el-option v-for="(item,key) in sizeOptions" :key="key" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
        <el-select v-model="backGroundValue" placeholder="背景" size="mini" @change="changeBackColor()">
          <el-option v-for="(item,key) in backGroundOptions" :key="key" :label="item.label" :value="item.value">
          </el-option>
        </el-select>
        <el-select v-model="widthValue" placeholder="宽度" size="mini" @change="setWidthValue()">
          <el-option v-for="(item,key) in widthOptions" :key="key" :label="item.label" :value="item.value">
          </el-option>
        </el-select>

      </div>
      <div class="switch">
        <el-switch v-model="nightValue" active-text="夜间" @change="changeNight"> </el-switch>

      </div>
    </div>
    <div class="bookname" style="" v-if="!listLoading">
      <h1>{{this.chapter.title}}</h1>
      <div class="bottem1" style="">
        <a href="javascript:void(0);" @click="chapterClick(chapter.pre_chapter)" :style="{'color':pre_no_chapter_color}">上一章</a> ←
        <a href="javascript:void(0);" @click="bookClick()" style="">章节目录</a> →
        <a href="javascript:void(0);" @click="chapterClick(chapter.next_chapter)" :style="{'color':next_no_chapter_color}">下一章</a>
      </div>
    </div>
    <div v-if="!listLoading" v-html="this.chapter.content" id="chapter-content" :style="{'font-family': fontsValue,'color':colorValue,'font-size':sizeValue,'width':widthValue}">
    </div>
    <div v-if="!listLoading" class="bottem2" style="">
      <a href="javascript:void(0);" @click="chapterClick(chapter.pre_chapter)" :style="{'color':pre_no_chapter_color}">上一章</a> ←
      <a href="javascript:void(0);" @click="bookClick()" style="">章节目录</a> →
      <a href="javascript:void(0);" @click="chapterClick(chapter.next_chapter)" :style="{'color':next_no_chapter_color}">下一章</a>

    </div>

  </div>
</template>

<script>
import { GetContent } from '@/api/search'

export default {
  components: {},
  methods: {
    bookClick(id) {
      this.$router.replace('/global_search/search_chapter/' + encodeURIComponent(this.novel_url) + '/' + encodeURIComponent(this.novel_name))
    },
    chapterClick(url) {
      if (url !== '') {
        this.$router.replace('/global_search/search_reader/' + encodeURIComponent(url) + "/" + encodeURIComponent(this.novel_url) + "/" + encodeURIComponent(this.novel_name))

        this.$router.go(0)
      }
    },
    getContent() {
      this.listLoading = true;
      // console.log("this.chapter:", this.chapter)
      GetContent({ content_url: this.content_url, novel_name: this.novel_name }).then(response => {
        this.chapter = response.data.content
        this.content_title = this.chapter.title
        // console.log("this.chapter2:", this.chapter)

        // this.chapter[0].content = this.chapter[0].content.replace(/\n/g, "<br/>")
        // this.chapter[0].content = this.chapter[0].content.replace(" ", "&nbsp;")
        // this.chapter[0].content = this.chapter[0].content.replace(/\r/g, "<br/>")
        // // this.chapter[0].content = this.chapter[0].content.replace("　　", "<br/>")
        // this.chapter[0].content = this.chapter[0].content.replace(/([^\x00-\xff])\s\s([^\x00-\xff])/ig, "$1<br/>&nbsp;&nbsp;&nbsp;&nbsp;$2")
        if (this.chapter.pre_chapter === '') {
          this.pre_no_chapter_color = "#E6D6D6"
        } else {
          this.pre_no_chapter_color = "#000000"
        }
        if (this.chapter.next_chapter === '') {
          this.next_no_chapter_color = "#E6D6D6"
        } else {
          this.next_no_chapter_color = "#000000"
        }
        this.listLoading = false;
        this.saveLastestChapter(this.content_url, this.novel_url, this.novel_name, this.content_title)
      })
    },

    saveLastestChapter(content_url, novel_url, novel_name, content_title) {
      let data = JSON.parse(localStorage.getItem('latest_view'));
      let temp = [];
      let find_in_local = false
      let timestamp = new Date().getTime()
      let newdata = { type: 2, novel_url: novel_url, content_url: content_url, content_title: content_title, novel_name: novel_name, timestamp: timestamp }
      if (data == null) {
        temp.push(newdata)
      } else {
        data.forEach(item => {
          if (item['novel_url'] === novel_url) {
            temp.push(newdata)
            find_in_local = true
          } else {
            temp.push(item)
          }
        })
        if (!find_in_local) {
          temp.push(newdata)
        }
      }

      // console.log(temp)
      if (temp.length > 40) {
        temp = temp.slice(0, 40)
      }

      localStorage.setItem('latest_view', JSON.stringify(temp.sort(this.compare('timestamp'))));
    },
    compare(property) {
      return function (a, b) {
        var value1 = a[property];
        var value2 = b[property];
        return value2 - value1;
      }
    },
    changeBackColor() {
      this.$store.dispatch("book/saveBbg", this.backGroundValue)
      this.$cookies.set("content_backcolor", this.backGroundValue, -1);
    },
    setFontsValue() {
      this.$cookies.set("content_fonts", this.fontsValue, -1);
    },
    setColorValue() {
      this.$cookies.set("content_color", this.colorValue, -1);
    },
    setSizeValue() {
      this.$cookies.set("content_size", this.sizeValue, -1);
    },
    setWidthValue() {
      this.$cookies.set("content_width", this.widthValue, -1);
    },
    initSetting() {
      this.backGroundValue = this.$cookies.get("content_backcolor") === '' ? this.backGroundValue : this.$cookies.get("content_backcolor")
      this.fontsValue = this.$cookies.get("content_fonts") === '' ? this.fontsValue : this.$cookies.get("content_fonts");
      this.colorValue = this.$cookies.get("content_color") === '' ? this.colorValue : this.$cookies.get("content_color");
      this.sizeValue = this.$cookies.get("content_size") === '' ? this.sizeValue : this.$cookies.get("content_size");
      this.widthValue = this.$cookies.get("content_width") === '' ? this.widthValue : this.$cookies.get("content_width");
      this.$store.dispatch("book/saveBbg", this.backGroundValue)
    },
    changeNight() {
      if (this.nightValue) {
        this.backGroundValue = "#454545"
        this.$store.dispatch("book/saveBbg", this.backGroundValue)
        this.$cookies.set("content_backcolor", this.backGroundValue, -1);
      } else {
        this.backGroundValue = '#f0f9eb'
        this.$store.dispatch("book/saveBbg", this.backGroundValue)
        this.$cookies.set("content_backcolor", this.backGroundValue, -1);
      }

    },
  },

  watch: {

    $route: {
      handler: function (val) {
        if (this.$route.path.indexOf("search_reader") !== -1) {
          this.novel_url = this.$route.params.novel_url
          this.content_url = this.$route.params.content_url
          this.novel_name = this.$route.params.novel_name
          this.getContent()
        }
      },
      deep: true
    }
  },
  data() {
    return {
      // chapter_id: 0,
      pre_no_chapter_color: '#333',
      next_no_chapter_color: '#333',
      // book_id: 0,
      pre_chapter: '',
      next_chapter: '',
      chapter: {},
      novel_url: '',
      content_url: '',
      novel_name: '',
      content_title: '',
      listLoading: true,
      fontsOptions: [
        { value: 'SimSun', label: '默认' },
        { value: 'SimSun', label: '字体' },
        { value: 'SimHei', label: '黑体' },
        { value: 'KaiTi', label: '楷体' },
        { value: 'Microsoft YaHei', label: '雅黑' },
        { value: 'FangSong', label: '仿宋' },
        { value: 'NSimSun', label: '新宋体' },
      ],
      fontsValue: 'SimSun',
      colorOptions: [
        { value: '#000000', label: '默认' },
        { value: '#000000', label: '颜色' },
        { value: '#9370DB', label: '暗紫' },
        { value: '#2E8B57', label: '藻绿' },
        { value: '#2F4F4F', label: '深灰' },
        { value: '#77889F', label: '青灰' },
        { value: '#800000', label: '栗色' },
        { value: '#6A5ACD', label: '青蓝' },
        { value: '#BC8F8F', label: '玫褐' },
        { value: '#F4A460', label: '黄褐' },
        { value: '#F5F5DC', label: '米色' },
        { value: '#F5F5F5', label: '雾白' },
      ],
      colorValue: '#000000',
      sizeOptions: [
        { value: '12pt', label: '12pt' },
        { value: '14pt', label: '14pt' },
        { value: '16pt', label: '默认' },
        { value: '16pt', label: '大小' },
        { value: '18pt', label: '18pt' },
        { value: '20pt', label: '20pt' },
        { value: '22pt', label: '22pt' },
        { value: '25pt', label: '25pt' },
        { value: '30pt', label: '30pt' },
      ],
      sizeValue: '16pt',
      backGroundOptions: [
        { value: '#f0f9eb', label: '默认' },
        { value: '#f0f9eb', label: '背景' },
        { value: '#FFFFFF', label: '白雪' },
        { value: '#000000', label: '漆黑' },
        { value: '#FFFFED', label: '明黄' },
        { value: '#EEFAEE', label: '淡绿' },
        { value: '#CCE8CF', label: '草绿' },
        { value: '#FCEFFF', label: '红粉' },
        { value: '#EFEFEF', label: '深灰' },
        { value: '#D2B48C', label: '茶色' },
        { value: '#E7F4FE', label: '银色' },
        { value: "#454545", label: '夜间' },
      ],
      backGroundValue: '#f0f9eb',
      widthOptions: [
        { value: '95%', label: '默认' },
        { value: '95%', label: '宽度' },
        { value: '85%', label: '85%' },
        { value: '76%', label: '75%' },
        { value: '67%', label: '65%' },
        { value: '53%', label: '50%' },
        { value: '41%', label: '40%' },
      ],
      widthValue: '95%',
      nightValue: false
    };
  },
  mounted() {
    this.novel_url = this.$route.params.novel_url
    this.content_url = this.$route.params.content_url
    this.novel_name = this.$route.params.novel_name
    this.getContent()
    this.$store.dispatch("book/saveHeader", true)
    this.initSetting()


  },
  destroyed() {
    this.$store.dispatch("book/saveHeader", false)

  }
}

</script>

<style lang="scss" scoped>
// body {
//   background-color: rgb(16, 29, 29);
// }
#content {
  // width: 1200px;
  // height: 900px;
  margin-top: 15px;
  border: 4px solid #aab8d8;
}
#header {
  border-bottom: 1px solid rgb(136, 198, 229);
  text-align: left;
  padding: 0px 10px;
  line-height: 40px;
  // height: 40px;
  background-color: rgb(225, 236, 237);
}
.bread {
  width: 500px;
  // height: 30px;

  display: inline-block;
  overflow: hidden;
}
.select {
  // width: 400px;
  width: 100%;
  // height: 30px;
  display: inline-block;
  // margin-left: 100px;
}
.el-select {
  width: 20%;
  display: inline-block;
}
.switch {
  margin-top: 10px;
  display: inline-block;
  float: right;
}
.el-switch {
  display: inline-block;
  float: right;
}
.bookname {
  border-bottom: 1px dashed rgb(136, 198, 229);
  line-height: 30px;
  padding-top: 10px;
  margin-bottom: 10px;
}
.bottem a,
.bottem1 a,
.bottem2 a {
  color: rgb(8, 83, 8);
  font-size: 14px;
  margin-left: 10px;
  margin-right: 10px;
}
.bottem2 {
  border-top: 1px dashed rgb(136, 198, 229);
  // text-align: center;
  // width: 900px;
  // // margin: auto 20px;
  // // padding: 15px;
  // clear: both !important;
}
#chapter-content {
  font-size: 19pt;
  letter-spacing: 0.2em;
  line-height: 150%;
  padding-top: 15px;
  width: 85%;
  margin: auto;
  // font-family: 宋体;
  width: 95%;
  color: rgb(0, 0, 0);
  display: block;
  text-align: left;
}
</style>