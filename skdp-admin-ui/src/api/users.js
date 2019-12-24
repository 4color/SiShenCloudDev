import request from '@/plugin/axios'

export function RequestLoadUserList (data) {
    return request({
        url: '/users/list',
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

export function DeleteXzq (data) {
    return request({
        url: '/xzq/' + data,
        method: 'delete'
    })
}

export function PutEnableXzq (data) {
    return request({
        url: '/xzq/enable/' + data.xzqid + '/' + data.enable,
        method: 'put'
    })
}
