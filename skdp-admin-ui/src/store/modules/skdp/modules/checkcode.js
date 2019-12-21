import { CheckCode } from '@api/sys.login'

export default {
    namespaced: true,
    actions: {
        //  获取验证码
        GetCheckCode ({ dispatch }) {
           return new Promise((resolve, reject) => {
               // 开始请求登录接口
               CheckCode()
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
