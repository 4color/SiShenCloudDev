<template>
    <d2-container type="full" class="page">

        <el-select slot="header" placeholder="请选择" style="width: 140px">
            <el-option label="游客" value="guest">
            </el-option>
            <el-option label="一般用户" value="nomarl">
            </el-option>
            <el-option label="VIP用户" value="vip">
            </el-option>
        </el-select>
        <el-input slot="header" placeholder="请输入查询内容" style="width: 150px;margin-left: 5px"></el-input>
        <el-button slot="header" style="margin-bottom: 5px; margin-left: 2px">查询</el-button>
        <el-button slot="header" style="margin-left: 2px; margin-bottom: 5px" @click="dialogVisible = true">新增
        </el-button>

        <!--用户列表-->
        <skl-user-list></skl-user-list>

        <!--新增-->
        <el-dialog
                title="新增角色"
                :visible.sync="dialogVisible"
                width="30%"
                >
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="角色名称">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="角色代码">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
    <el-button @click="dialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
  </span>
        </el-dialog>
    </d2-container>
</template>

<script>
  export default {
    name: 'users',
    data () {
      return {
        dialogVisible: false,
        form: {},

        roledata: [
          {
            lable: '游客',
            value: '33'
          },
          {
            lable: '一般用户',
            value: '35'
          },
          {
            lable: 'VIP用户',
            value: '44'
          },
          {
            lable: '尊贵VIP用户',
            value: '33'
          }
        ],
        defaultProps: {
          children: 'children',
          label: 'lable'
        }
      }
    },
    mounted () {
    },
    methods: {
      indexMethod (index) {
        return index + 1
      },
      handleEdit (index, row) {
        console.log(index, row)
      },
      handleDelete (index, row) {
        console.log(index, row)
      },
      jumpUrl (row) {
        this.$router.push('/department/xzqxz/' + row.xzqdm)
      },
      filterTag (value, row) {
        return row.tag === value
      },
    }
  }
</script>

<style lang="scss" scoped>
    .page {
        .vue-grid-layout {
            background-color: $color-bg;
            border-radius: 4px;
            margin: -10px;

            .page_card {
                height: 100%;
                @extend %unable-select;
            }

            .vue-resizable-handle {
                opacity: .3;

                &:hover {
                    opacity: 1;
                }
            }
        }
    }
</style>
