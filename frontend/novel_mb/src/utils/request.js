import axios from 'axios'
//import store from '@/store'
import { Message } from 'element-ui'
// create an axios instance
const service = axios.create({
  // baseURL: 'http://localhost:8888/api/private/v1/', // api 的 base_url
  // baseURL: '/api',
  withCredentials: true, // 跨域请求时发送 cookies
  timeout: 10000 // request timeout
})

// request interceptor
//service.interceptors.request.use()
// request interceptor
service.interceptors.request.use(
  config => {
    // Do something before request is sent
    // if (store.getters.token) {
    //   // 让每个请求携带token-- ['X-Token']为自定义key 请根据实际情况自行修改
    //   config.headers['X-Token'] = getToken()
    // }
    return config
  },
  error => {
    // Do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)


service.interceptors.response.use(

  response => {
    // console.log(JSON.parse(response.data))
    // response = JSON.parse(response)
    // console.log("response1:", response)
    const res = response.data
    // console.log(res)
    // res.data["message"]=
    if (res.code !== 0) {
      // Message({
      //   message: res.msg || 'error',
      //   type: 'error',
      //   duration: 5 * 1000
      // })
    } else {
      return res
    }
  },
  error => {
    console.log('err' + error) // for debug
    // Message({
    //   message: error.message,
    //   type: 'error',
    //   duration: 5 * 1000
    // })
    return Promise.reject(error)
  }
)
export default service
