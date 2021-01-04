//state为单一状态树
const state = {
  bodyBackground: '#f0f9eb',
  searchKey: '',
  hideHeader: false,
}
//更改store中state状态的唯一方法
const mutations = {
  SET_BBG: (state, bodyBackground) => {
    state.bodyBackground = bodyBackground
  },
  SET_SEARCH: (state, key) => {
    state.searchKey = key
  },
  SET_HEADER: (state, header) => {
    state.hideHeader = header
  },
}

const actions = {
  // user login
  saveBbg({ commit }, data) {
    commit('SET_BBG', data)
  },
  saveSearch({ commit }, data) {
    commit('SET_SEARCH', data)
  },
  saveHeader({ commit }, data) {
    commit('SET_HEADER', data)
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
