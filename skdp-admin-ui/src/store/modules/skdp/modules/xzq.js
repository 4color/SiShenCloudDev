import { LoadXzqList } from '@api/xzq'

export default {
    namespaced: true,
    actions: {
        //  获取验证码
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
        }

    }
}
