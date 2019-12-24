<template>
    <el-table
            :data="data"
            stripe
            style="width: 100%">
        <el-table-column
                type="index"
                :index="indexMethod">
        </el-table-column>
        <el-table-column
                label="行政区"
                width="180">
        </el-table-column>
        <el-table-column
                prop="nickname"
                label="昵称"
                width="180">
        </el-table-column>
        <el-table-column
                prop="email"
                label="电子邮件"
                width="180">
        </el-table-column>
        <el-table-column
                prop="createdat"
                label="注册时间">
            <template slot-scope="scope">
                {{ProcessTime(scope.row.createdat)}}
            </template>
        </el-table-column>
        <el-table-column
                prop="source"
                label="注册来源">
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


                <el-button
                        size="mini"
                        type="primary"
                        icon="el-icon-s-tools"
                        title="用户扩展"
                        style="padding: 5px;"
                        @click="handleDelete(scope.$index, scope.row)">
                </el-button>

            </template>
        </el-table-column>
    </el-table>
</template>

<script src="./skl-user-list.js">
</script>

<style scoped>

</style>
