
const CompressionPlugin = require('compression-webpack-plugin')
module.exports = {
  // publicPath: process.env.NODE_ENV === 'production' ? '/mb' : '/',
  configureWebpack: config => {
    if (process.env.NODE_ENV === 'production') {
      return {
        plugins: [new CompressionPlugin({
          filename: '[path].gz[query]',
          test: new RegExp(
            '\\.(js|css)$'    //压缩 js 与 css
          ),
          threshold: 10240,
          minRatio: 0.8
        })]
      }
    }
  },
  devServer: {
    // host: "localhost", //要设置当前访问的ip 否则失效
    port: 8089,//当前web服务端口
    // open: false, //浏览器自动打开页面
    // publicPath: '/mb',
    proxy: {
      '/api': {
        target: 'http://localhost:3001/api/',

        ws: true,//是否代理websocket
        changeOrigin:
          true,//是否跨域
        pathRewrite: {
          '^/api': ''//url重写
        }
      },
      '/search_api': {
        target: 'http://localhost:8079/search_api/',//目标地址

        ws: true,//是否代理websocket
        changeOrigin:
          true,//是否跨域
        pathRewrite: {
          '^/search_api': ''//url重写
        }
      }
    }
  }
}