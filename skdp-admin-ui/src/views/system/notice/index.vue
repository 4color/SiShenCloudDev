<template>

    <d2-container>
        <el-input slot="header" placeholder="请输入查询内容" style="width: 150px;margin-left: 5px"></el-input>
        <el-button slot="header" style="margin-bottom: 5px; margin-left: 2px">查询</el-button>

        <el-button slot="header" type="primary" icon="el-icon-circle-plus-outline"
                   style="margin-left: 2px;margin-bottom: 5px;"
                   @click="addnotice()">新增公告
        </el-button>
        <el-table
                :data="data"
                stripe
                style="width: 100%">
            <el-table-column
                    type="index"
                    :index="indexMethod">
            </el-table-column>
            <el-table-column
                    prop="pubtime"
                    label="发布时间"
                    width="180">
            </el-table-column>
            <el-table-column
                    prop="title"
                    label="公告标题"
                   >
                <template slot-scope="scope">
                    <el-button type="text" @click="toDetail(scope.row)">{{scope.row.title}}</el-button>
                </template>
            </el-table-column>
            <el-table-column
                    prop="user"
                    label="发布人"
                    width="180">

            </el-table-column>
            <el-table-column
                    prop="enable"
                    label="有效"
                    width="100"
                    :filters="[{ text: '有效', value: true }, { text: '无效', value: false }]"
                    :filter-method="filterTag"
                    filter-placement="bottom-end">
                <template slot-scope="scope">
                    <el-tag
                            :type="scope.row.enable ? 'primary' : 'info'"
                            disable-transitions>{{scope.row.enable}}
                    </el-tag>
                </template>
            </el-table-column>
            <el-table-column
                    align="right"
                    width="200">
                <template slot-scope="scope">

                    <el-button
                            size="mini" icon="el-icon-edit"
                            title="编辑"
                            style="margin-right: 5px;padding: 5px;"
                            @click="handleEdit(scope.$index, scope.row)">
                    </el-button>
                    <el-popconfirm
                            title="这是一段内容确定删除吗？"
                    >
                        <el-button
                                size="mini"
                                type="danger"
                                icon="el-icon-delete"
                                title="删除"
                                slot="reference"
                                style="margin-right: 5px;padding: 5px;"
                                @click="handleDelete(scope.$index, scope.row)">
                        </el-button>

                    </el-popconfirm>

                    <el-button
                            size="mini"
                            type="warning plain"
                            icon="el-icon-check"
                            title="启用"
                            style="padding: 5px;"
                            v-if="scope.row.enable===false"
                            @click="handleDelete(scope.$index, scope.row)">
                    </el-button>

                    <el-button
                            size="mini"
                            type="success plain"
                            icon="el-icon-close"
                            title="禁用"
                            style="padding: 5px;"
                            v-if="scope.row.enable"
                            @click="handleDelete(scope.$index, scope.row)">
                    </el-button>

                </template>
            </el-table-column>
        </el-table>

        <div class="block" style="text-align: center; padding: 20px">
            <el-pagination
                    :page-size="100"
                    layout="total, prev, pager, next"
                    :total="1000">
            </el-pagination>
        </div>
    </d2-container>
</template>

<script>
  export default {
    name: 'noticeindex',
    data () {
      return {
        data: [
          {
            articleid: 'test',
            pubtime: 'test',
            title: '拱墅区',
            user: '浙江省',
            'enable': true
          },
          {
            articleid: 'test',
            username: '浙江省',
            xzqmc: '拱墅区',
            xzqdm: '330111',
            'enable': true
          },
          {
            sheng: '浙江省',
            xzqmc: '拱墅区',
            xzqdm: '330111',
            'enable': true
          },
          {
            sheng: '浙江省',
            xzqmc: '拱墅区',
            xzqdm: '330111',
            'enable': false
          }
        ],
      }
    },
    methods: {
      toDetail (row) {
        this.$router.push('/system/noticedetail/' + row.articleid)
      }
    }
  }
</script>

<style scoped>

</style>
