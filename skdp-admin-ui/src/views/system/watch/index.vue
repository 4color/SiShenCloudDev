<template>
    <d2-container type="full" class="page">
        <div style="margin: 10px">

            <el-row :gutter="20">
                <el-col :span="8">

                    <el-card class="box-card" shadow="hover">
                        <div slot="header" class="clearfix">
                            <span>CPU监控</span>
                        </div>
                        <ve-liquidfill :data="chartData"></ve-liquidfill>
                        <el-alert
                                title="成功提示的文案"
                                type="success"
                                :closable="false">
                        </el-alert>
                    </el-card>

                </el-col>
                <el-col :span="8">

                    <el-card class="box-card" shadow="hover">
                        <div slot="header" class="clearfix">
                            <span>内存监控</span>
                        </div>
                        <ve-liquidfill :data="chartData"></ve-liquidfill>
                        <el-alert
                                title="8G/32G"
                                type="success"
                                :closable="false">
                        </el-alert>
                    </el-card>
                </el-col>
                <el-col :span="8">
                    <el-card class="box-card" shadow="hover">
                        <div slot="header" class="clearfix">
                            <span>磁盘监控</span>
                        </div>
                        <ve-liquidfill :data="chartData"></ve-liquidfill>
                        <el-alert
                                title="100M/2048M"
                                type="warning"
                                :closable="false">
                        </el-alert>
                    </el-card>
                </el-col>
            </el-row>

        </div>
    </d2-container>
</template>

<script>

    export default {
        name: 'watchindex',
        data () {
            return {
                chartData: {
                    columns: ['city', 'percent'],
                    rows: [{
                        city: 'CPU',
                        percent: 0.6
                    }]
                }
            }
        }
    }
</script>

<style lang="scss" scoped>
    .el-alert__title {
        font-size: 20px !important;
    }
</style>
