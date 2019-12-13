<template>
    <d2-container>
        <el-select slot="header" v-model="search.type" placeholder="请选择" style="width: 140px">
            <el-option label="平台类" value="guest">
            </el-option>
            <el-option label="地图类" value="nomarl">
            </el-option>
            <el-option label="辅助类" value="vip">
            </el-option>
        </el-select>
        <el-input slot="header" placeholder="请输入查询内容" style="width: 150px;margin-left: 5px"></el-input>
        <el-button slot="header" style="margin-bottom: 5px; margin-left: 2px">查询</el-button>

        <el-button slot="header" type="primary" icon="el-icon-circle-plus-outline" style="margin-left: 2px;margin-bottom: 5px;"
                   @click="addapp()">新增应用
        </el-button>

        <div style="padding:10px">
            <el-row :gutter="12">
                <el-col :span="8">
                    <el-card shadow="hover" class="el-alert el-alert--success is-dark myalert">
                        待审核应用数：90
                    </el-card>
                </el-col>
                <el-col :span="8">
                    <el-card shadow="hover" class="el-alert el-alert--info is-dark myalert">
                        拒绝的应用数：10
                    </el-card>
                </el-col>
                <el-col :span="8">
                    <el-card shadow="hover" class="el-alert el-alert--warning is-dark myalert">
                        推荐的应用：5
                    </el-card>
                </el-col>
            </el-row>

            <el-table
                    :data="data"
                    stripe
                    style="width: 100%">
                <el-table-column
                        type="index"
                        :index="indexMethod">
                </el-table-column>
                <el-table-column
                        prop="appname"
                        label="应用名称"
                        width="180">
                    <template slot-scope="scope">
                        <el-button type="text" @click="toDetail(scope.row)">{{scope.row.appname}}</el-button>
                    </template>
                </el-table-column>
                <el-table-column
                        prop="type"
                        label="分类"
                        width="180">
                </el-table-column>
                <el-table-column
                        prop="version"
                        label="最新版本">
                </el-table-column>
                <el-table-column
                        prop="author"
                        label="提交者">
                </el-table-column>
                <el-table-column
                        prop="createtime"
                        label="提交时间">
                </el-table-column>
                <el-table-column
                        prop="checkstatus"
                        label="审核状态">
                </el-table-column>
                <el-table-column
                        prop="enable"
                        label="有效"
                        width="100"
                        :filters="[{ text: '有效', value: true }, { text: '无效', value: false }]"
                        filter-placement="bottom-end">
                    <template slot-scope="scope">
                        <el-tag
                                :type="scope.row.enable ? 'primary' : 'info'"
                                disable-transitions>{{scope.row.enable}}
                        </el-tag>
                    </template>
                </el-table-column>
                <el-table-column
                        align="right">
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
            <div class="block">
                <el-pagination
                        :page-size="100"
                        layout="total, prev, pager, next"
                        :total="1000">
                </el-pagination>
            </div>
        </div>
    </d2-container>
</template>

<script>
  export default {
    name: 'market',
    data () {
      return {
        currentDate: new Date(),
        data: [
          {
            appid: 'test',
            appname: '四神考试系统',
            version: '0.1',
            type: '考试类',
            author: '四神',
            'enable': true
          },
          {
            appid: 'test',
            appname: '四神BI系统',
            version: '0.1',
            type: '考试类',
            author: '四神',
            'enable': true
          },
          {
            appname: '四神地图(WEBGIS)平台',
            version: '0.1',
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
        search:{
          type : "guest"
        }
      }
    }, methods: {
      indexMethod (index) {
        return index + 1
      },
      toDetail (row) {
        this.$router.push('/market/detail/' + row.appid)
      },
      addapp ()      {
        this.$router.push('/market/addapp/')
      }
    }
  }
</script>

<style scoped>
    .myalert {
        font-size: 20px;
    }
</style>
