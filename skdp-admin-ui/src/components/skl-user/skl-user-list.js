import { mapActions } from 'vuex'
import util from '@/libs/util.js'

export default {
    name: 'skluserlist',
    data () {
        return {
            data: [],
            form: { where: '' },
            list: {
                pageindex: 1,
                pagesize: 15,
                total: 0
            }
        }
    },
    mounted () {
        this.loadList()
    },
    methods: {
        ...mapActions('skdp/users', [
            'MethodLoadUserList'
        ]),
        indexMethod (index) {
            return ((this.list.pageindex - 1) * this.list.pagesize) + index + 1
        },
        toDetail (row) {
            this.$router.push('/department/userdetail/' + row.userid)
        },
        ProcessTime (data) {
            return util.ProcessTime(data)
        },
        //  用户列表
        loadList () {
            var model = {
                'where': this.form.where,
                'pageindex': this.list.pageindex,
                'pagesize': this.list.pagesize
            }
            this.MethodLoadUserList(model).then((res) => {
                if (res != null) {
                    if (res.status === 200) {
                        this.total = res.data.count
                        this.data = res.data.data
                    } else {
                        this.$message.error(res.message)
                    }
                }
            })
        }
    }
}
