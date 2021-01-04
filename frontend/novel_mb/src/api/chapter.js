import request from '@/utils/request'


export function GetChapterInfo(data) {
  return request({
    url: '/api/get_chapter_info',
    method: 'post',
    data
  })
}

export function GetChapterContent(data) {
  return request({
    url: '/api/get_chapter_content',
    method: 'post',
    data
  })
}
