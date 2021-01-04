<template>

  <div :class="{'hidden':hidden}" class="pagination-container">

    <el-pagination :background="background" small :current-page.sync="currentPage" :page-size.sync="pageSize" :layout="layout" :page-sizes="pageSizes" :total="total" v-bind="$attrs" :pager-count="3" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
  </div>
</template>

<script>
export default {
  name: 'Pagination',
  props: {
    total: {
      required: true,
      type: Number
    },
    page: {
      type: Number,
      default: 1
    },
    limit: {
      type: Number,
      default: 200
    },
    pageSizes: {
      type: Array,
      default() {
        return [100, 200, 300, 500]
      }
    },
    layout: {
      type: String,
      default: 'total, sizes, prev, pager, next'
    },
    background: {
      type: Boolean,
      default: true
    },
    autoScroll: {
      type: Boolean,
      default: true
    },
    hidden: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    currentPage: {
      get() {
        return this.page
      },
      set(val) {
        this.$emit('update:page', val)
      }
    },
    pageSize: {
      get() {
        return this.limit
      },
      set(val) {
        this.$emit('update:limit', val)
      }
    }
  },
  methods: {
    handleSizeChange(val) {
      this.$emit('pagination', { page: this.currentPage, limit: val })
    },
    handleCurrentChange(val) {
      this.$emit('pagination', { page: val, limit: this.pageSize })
    }
  }
}
</script>

<style scoped>
.pagination-container {
  margin-left: -20px;
  /* width: 90%; */
  background: #fff;
  padding: 32px 16px;
}
.pagination-container.hidden {
  display: none;
}
</style>
