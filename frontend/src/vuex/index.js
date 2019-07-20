import Vue from 'vue'
import Vuex from 'vuex'
import Cookies from 'js-cookie'
import routerData from './modules/routerData'
import role from './modules/role'
import layout from './modules/layout/index'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    token: Cookies.get('token'),
    nickname: Cookies.get('nickname')
  },
  mutations: {
    setToken (state, token) {
      state.token = token
      Cookies.set('token', token ,{ expires: 1/24 });
    },
    setNickname (state, nickname) {
      state.nickname = nickname
      Cookies.set('nickname', nickname ,{ expires: 1/24 });
    }
  },
  actions: {
    setToken ({commit}, token) {
      return new Promise((resolve, reject) => {
        commit('setToken', token)
        resolve()
      })
    },
    setNickname ({commit}, nickname) {
      return new Promise((resolve, reject) => {
        commit('setNickname', nickname)
        resolve()
      })
    }
  },
  getters: {
    addRouters: state => state.routerData.addRouters,
    token: state => state.token,
    nickname: state => state.nickname,
    info: state => state.role.info,
    routers: state => state.routerData.routers,
    logoShow: state => state.layout.logoShow,
    isCollapse: state => state.layout.isCollapse,
    uniquerouter: state => state.layout.uniquerouter,
    tabnavBox: state => state.layout.tabnavBox,
    visible: state => state.layout.visible,
    left: state => state.layout.left,
    top: state => state.layout.top,
    rightNav: state => state.layout.rightNav,
  },
  modules: {
    routerData,
    role,
    layout
  },
})

export default store
