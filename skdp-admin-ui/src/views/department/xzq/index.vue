<template>
    <d2-container type="full" class="page">
        <el-radio-group v-model="form.level" slot="header" size="mini">
            <el-radio :label=1>省</el-radio>
            <el-radio :label=2>市</el-radio>
            <el-radio :label=3>县</el-radio>
            <el-radio :label=4>乡镇</el-radio>
            <el-radio :label=5>村</el-radio>
        </el-radio-group>
        <el-input slot="header" placeholder="请输入查询内容" style="width: 150px;margin-left: 5px"
                  v-model="form.where"></el-input>
        <el-button slot="header" style="margin-bottom: 5px; margin-left: 2px" @click="loadXzqQuery">查询</el-button>
        <el-button slot="header" style="margin-left: 2px" @click="showAddxzq">新增</el-button>


        <SplitPane :min-percent='17' :default-percent='17' split="vertical">
            <template slot="paneL">
                <el-tree :props="defaultProps" :auto-expand-parent="true" :load="loadXzqTree" @node-click="treeClick"
                         lazy>
                      <span slot-scope="{ node, data }">
      <i class="el-icon-folder"></i>
      <span style="padding-left: 4px; font-size: 12px">{{data.xzqmc}}</span>
    </span>

                </el-tree>
            </template>
            <template slot="paneR">
                <el-table
                        :data="data"
                        stripe
                        style="width: 100%">
                    <el-table-column
                            type="index"
                            :index="indexMethod">
                    </el-table-column>
                    <el-table-column
                            prop="xzqmc"
                            label="行政区名称"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="xzqvalue"
                            label="行政区代码">
                    </el-table-column>
                    <el-table-column
                            prop="enable"
                            label="有效"
                            width="100"
                            :filters="[{ text: '有效', value: 1 }, { text: '无效', value: 0 }]"
                            :filter-method="filterTag"
                            filter-placement="bottom-end">
                        <template slot-scope="scope">
                            <el-tag
                                    :type="scope.row.enable=='1' ? 'primary' : 'info'"
                                    disable-transitions>{{scope.row.enable=='1'?'有效':'无效'}}
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
                                    @click="showEdit(scope.$index, scope.row)">
                            </el-button>
                            <el-popconfirm
                                    title="确定删除吗？"
                                    icon="el-icon-info"
                                    @onConfirm="handleDelete(scope.$index, scope.row)"
                            >
                                <el-button
                                        size="mini"
                                        type="danger"
                                        icon="el-icon-delete"
                                        title="删除"
                                        slot="reference"
                                        style="margin-right: 5px;padding: 5px;"
                                >
                                </el-button>

                            </el-popconfirm>

                            <el-button
                                    size="mini"
                                    type="success plain"
                                    icon="el-icon-open"
                                    title="启用"
                                    style="padding: 5px;"
                                    v-if="scope.row.enable===0"
                                    @click="EnabledOrDisableXzq( scope.row,1)">
                            </el-button>

                            <el-button
                                    size="mini"
                                    type="Info plain"
                                    icon="el-icon-turn-off"
                                    title="禁用"
                                    style="padding: 5px;"
                                    v-if="scope.row.enable===1"
                                    @click="EnabledOrDisableXzq( scope.row,0)">
                            </el-button>


                            <el-button
                                    size="mini"
                                    type="primary"
                                    icon="el-icon-s-tools"
                                    title="行政区扩展"
                                    style="padding: 5px;"
                                   >
                            </el-button>

                        </template>
                    </el-table-column>
                </el-table>

                <div class="block" style="text-align: center; padding: 20px">
                    <el-pagination
                            :page-size="pagesize"
                            layout="total, prev, pager, next"
                            :total="total"
                            @current-change="handleCurrentChange">
                    </el-pagination>
                </div>
            </template>
        </SplitPane>

        <!--新增-->
        <el-dialog
                title="新增行政区"
                :visible.sync="dialogVisible"
                width="30%"
        >
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="父级行政区">
                    {{this.parentxzqmc}}
                </el-form-item>
                <el-form-item label="行政区名称">
                    <el-input v-model="addxzqform.xzqmc"></el-input>
                </el-form-item>
                <el-form-item label="行政区代码">
                    <el-input v-model="addxzqform.xzqvalue"></el-input>
                </el-form-item>
                <el-form-item label="年份">
                    <el-select v-model="addxzqform.xzqnf" placeholder="请选择">
                        <el-option
                                v-for="n in nowDate"
                                :key="n"
                                :label="n"
                                :value="n" v-if="n>2015">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="层级">
                    <el-select v-model="addxzqform.xzqlevel" placeholder="请选择">
                        <el-option value="1" label="省级"></el-option>
                        <el-option value="2" label="市级"></el-option>
                        <el-option value="3" label="区县级"></el-option>
                        <el-option value="4" label="乡镇级"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
    <el-button @click="dialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="handleAddxzq">确 定</el-button>
  </span>
        </el-dialog>
    </d2-container>
</template>

<script src="./index.js"></script>

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
