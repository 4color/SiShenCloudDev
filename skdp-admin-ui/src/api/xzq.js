import request from '@/plugin/axios'

export function LoadXzqList (data) {
    return request({
        url: '/xzq/list',
        method: 'post',
        data
    })
}

export function AddXzq (data) {
    return request({
        url: '/xzq/add',
        method: 'post',
        data
    })
}

export function XzqInfo (data) {
    return request({
        url: '/xzq/id/' + data,
        method: 'get'
    })
}
