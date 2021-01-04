import request from '@/utils/request'


export function GetRecommend(data) {
  return request({
    url: '/api/get_recommend',
    method: 'post',
    data
  })
}
export function GetRank(data) {
  return request({
    url: '/api/get_rank',
    method: 'post',
    data
  })
}


export function GetBookById(data) {
  return request({
    url: '/api/get_book_by_id',
    method: 'post',
    data
  })
}
export function GetBookByIdList(data) {
  return request({
    url: '/api/get_book_by_id_list',
    method: 'post',
    data
  })
}

export function GetBookByCategory(data) {
  return request({
    url: '/api/get_book_by_category',
    method: 'post',
    data
  })
}

export function GetBookByMtime(data) {
  return request({
    url: '/api/get_book_by_mtime',
    method: 'post',
    data
  })
}
export function GetBookByCtime(data) {
  return request({
    url: '/api/get_book_by_ctime',
    method: 'post',
    data
  })
}

// get_book_by_name
export function GetBookByName(data) {
  return request({
    url: '/api/get_book_by_name',
    method: 'post',
    data
  })
}