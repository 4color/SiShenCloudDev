import { mapActions } from 'vuex'

export default {
    name: 'departmentxzq',
    data () {
        return {
            nowDate: new Date().getFullYear(),
            dialogVisible: false,
            form: { where: '', level: -1 },
            addxzqform: {
                xzqvalue: '',
                xzqmc: '',
                parentid: '',
                xzqnf: this.nowDate,
                changetime: '',
                xzqlevel: '1',
                xzqid: ''
            },
            data: [],
            defaultProps: {
                children: 'children',
                label: 'xzqmc'
            },
            pageindex: 1,
            pagesize: 15,
            total: 0,
            parentid: '',
            parentxzqmc: '全国'
        }
    },
    mounted () {
        this.loadXzqQuery()
    },
    methods: {
        ...mapActions('skdp/xzq', [
            'GetLoadXzqList', 'PutAddXzq', 'GetXzqInfo', 'DeleteXzq','PutEnableXzq'
        ]),
        indexMethod (index) {
            return ((this.pageindex - 1) * this.pagesize) + index + 1
        },
        jumpUrl (row) {
            this.$router.push('/department/xzqxz/' + row.xzqdm)
        },
        filterTag (value, row) {
            return row.tag === value
        },
        //  行政区列表
        loadXzqQuery () {
            var model = {
                'where': this.form.where,
                'level': this.form.level,
                'pageindex': this.pageindex,
                'pagesize': this.pagesize,
                'parentid': this.parentid,
                'parentlikeorequal': 'like'
            }
            this.GetLoadXzqList(model).then((res) => {
                if (res != null) {
                    if (res.status === 200) {
                        this.total = res.data.count
                        this.data = res.data.data
                    } else {
                        this.$message.error(res.message)
                    }
                }
            })
        },
        //  行政区列表
        loadXzqTree (node, resolve) {
            if (node.level === 0) {
                return resolve([{
                    xzqmc: '全国',
                    xzqid: ''
                }])
            }
            var pParentid = node.data.xzqid
            this.parentxzqmc = node.data.xzqmc
            var model = {
                'where': '',
                'level': -1,
                'pageindex': 1,
                'pagesize': 99,
                'parentid': pParentid,
                'parentlikeorequal': 'equal'
            }
            this.GetLoadXzqList(model).then((res) => {
                if (res != null) {
                    if (res.status === 200) {
                        this.total = res.data.count
                        resolve(res.data.data)
                    } else {
                        this.$message.error(res.message)
                    }
                }
            })
        },
        //  行政树被点点
        treeClick (data, node, event) {
            this.parentid = data.xzqid
            this.parentxzqmc = data.xzqmc
            this.loadXzqQuery()
        },
        handleCurrentChange (val) {
            this.pageindex = val
            this.loadXzqQuery()
        },
        showAddxzq () {
            this.addxzqform.parentid = this.parentid
            this.addxzqform.changetime = new Date()
            if (this.parentid.length === 0) {
                this.addxzqform.xzqlevel = '1'
            }
            if (this.parentid.length === 2) {
                this.addxzqform.xzqlevel = '2'
            }
            if (this.parentid.length === 4) {
                this.addxzqform.xzqlevel = '3'
            }
            if (this.parentid.length === 6) {
                this.addxzqform.xzqlevel = '4'
            }

            this.dialogVisible = true
        },
        //  新增或编辑行政区
        handleAddxzq () {
            this.addxzqform.xzqlevel = parseInt(this.addxzqform.xzqlevel)
            this.PutAddXzq(this.addxzqform).then((res) => {
                if (res != null) {
                    if (res.status === 200) {
                        this.$message.success(res.message)
                        this.loadXzqQuery()
                        this.dialogVisible = false
                    } else {
                        this.$message.error(res.message)
                    }
                }
            })
        },
        //  编辑行政区
        showEdit (index, row) {
            this.GetXzqInfo(row.xzqid).then((res) => {
                if (res != null) {
                    if (res.status === 200) {
                        this.addxzqform = res.data
                        this.addxzqform.xzqlevel = this.addxzqform.xzqlevel.toString()
                        this.dialogVisible = true
                    } else {
                        this.$message.error(res.message)
                    }
                }
            })
        },
        handleDelete (index, row) {
            this.DeleteXzq(row.xzqid).then((res) => {
                if (res != null) {
                    if (res.status === 200) {
                        this.$message.success(res.message)
                        this.loadXzqQuery()
                    } else {
                        this.$message.error(res.message)
                    }
                }
            })
        },
        //  开启或禁用行政区
        EnabledOrDisableXzq (row, isdis) {
            var model = { 'xzqid': row.xzqid, 'enable': isdis }
            this.PutEnableXzq(model).then((res) => {
                if (res != null) {
                    if (res.status === 200) {
                        this.$message.success(res.message)
                        this.loadXzqQuery()
                    } else {
                        this.$message.error(res.message)
                    }
                }
            })
        }
    }
}
