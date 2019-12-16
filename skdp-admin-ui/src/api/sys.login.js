import request from '@/plugin/axios'

export function AccountLogin (data) {
    return request({
        url: '/manager/login',
        method: 'post',
        data
    })
}

export function CheckCode () {
    return request({
        url: '/checkcode',
        method: 'get'
    })
}
