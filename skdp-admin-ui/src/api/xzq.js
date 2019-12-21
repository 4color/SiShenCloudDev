import request from '@/plugin/axios'

export function LoadXzqList (data) {
    return request({
        url: '/xzq/list',
        method: 'post',
        data
    })
}
