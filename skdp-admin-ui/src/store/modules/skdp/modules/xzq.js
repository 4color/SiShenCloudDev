import { LoadXzqList, AddXzq, XzqInfo, DeleteXzq, PutEnableXzq } from '@api/xzq'

export default {
    namespaced: true,
    actions: {
        //  获取行政区
        GetLoadXzqList ({ dispatch }, model) {
            return new Promise((resolve, reject) => {
                // 开始请求登录接口
                LoadXzqList(model)
                    .then(async res => {
                        // 结束
                        resolve(res)
                    })
                    .catch(err => {
                        console.log('err: ', err)
                        reject(err)
                    })
            })
        },
        //  修改行政区
        PutAddXzq ({ dispatch }, model) {
            return new Promise((resolve, reject) => {
                // 开始请求登录接口
                AddXzq(model)
                    .then(async res => {
                        // 结束
                        resolve(res)
                    })
                    .catch(err => {
                        console.log('err: ', err)
                        reject(err)
                    })
            })
        },
        //  获取行政区
        GetXzqInfo ({ dispatch }, model) {
            return new Promise((resolve, reject) => {
                // 开始请求登录接口
                XzqInfo(model)
                    .then(async res => {
                        // 结束
                        resolve(res)
                    })
                    .catch(err => {
                        console.log('err: ', err)
                        reject(err)
                    })
            })
        },
        //  获取行政区
        DeleteXzq ({ dispatch }, model) {
            return new Promise((resolve, reject) => {
                // 开始请求登录接口
                DeleteXzq(model)
                    .then(async res => {
                        // 结束
                        resolve(res)
                    })
                    .catch(err => {
                        console.log('err: ', err)
                        reject(err)
                    })
            })
        },
        //  开启禁用行政区
        PutEnableXzq ({ dispatch }, model) {
            return new Promise((resolve, reject) => {
                // 开始请求登录接口
                PutEnableXzq(model)
                    .then(async res => {
                        // 结束
                        resolve(res)
                    })
                    .catch(err => {
                        console.log('err: ', err)
                        reject(err)
                    })
            })
        }
    }
}
