import { RequestLoadUserList } from '@api/users'

export default {
    namespaced: true,
    actions: {
        //  获取用户列表
        MethodLoadUserList ({ dispatch }, model) {
            return new Promise((resolve, reject) => {
                // 开始请求登录接口
                RequestLoadUserList(model)
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
