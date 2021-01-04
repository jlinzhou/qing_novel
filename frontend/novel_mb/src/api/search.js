import request from '@/utils/request'


export function GetChapter(data) {
  return request({
    url: '/search_api/chapter',
    method: 'get',
    params: data,
  })
}

export function GetSearch(data) {
  return request({
    url: '/search_api/search',
    method: 'get',
    params: data,
  })
}


export function GetContent(data) {
  return request({
    url: '/search_api/content',
    method: 'get',
    params: data,
  })
}
