<template>
    <d2-container type="full" class="page">

        <el-button slot="header" style="margin-left: 2px; margin-bottom: 5px" @click="dialogVisible = true">新增</el-button>

        <SplitPane :min-percent='17' :default-percent='17' split="vertical">
            <template slot="paneL">
                <el-tree :data="roledata" :props="defaultProps" :auto-expand-parent="true"></el-tree>
            </template>
            <template slot="paneR" >
                <el-tabs v-model="activeName" style="padding: 5px">
                    <el-tab-pane label="用户列表" name="user">
                        <!--用户列表-->
                        <skl-user-list></skl-user-list>
                    </el-tab-pane>
                    <el-tab-pane label="角色权限" name="private">角色权限</el-tab-pane>
                </el-tabs>

            </template>
        </SplitPane>

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
    name: 'role',
    data () {
      return {
        activeName:"user",
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
      filterTag(value, row) {
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
