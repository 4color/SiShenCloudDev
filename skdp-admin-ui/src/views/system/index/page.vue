<template>
    <d2-container type="full" class="page">
        <div style="margin: 10px">

            <el-row :gutter="20">
                <el-col :span="24">

                    <el-card class="box-card" shadow="never" :body-style="{ paddingBottom: '5px' }">
                        <el-row>
                            <el-col :span="6">
                                <div style="text-align: center; padding-top: 50px">
                                    <div class="countup">
                                        <d2-count-up :end="100" :decimals="0"/>
                                    </div>
                                    <div>今日系统访问量</div>
                                </div>
                            </el-col>
                            <el-col :span="12">
                                <ve-line :data="chartDataVistor" :settings="chartSettingsVistor"
                                         height="200px" :legend-visible="false"
                                         :grid={height:180,top:10}></ve-line>
                            </el-col>
                            <el-col :span="6">
                                <div style="text-align: center; padding-top: 50px">
                                    <div class="countup">
                                        <d2-count-up :end="510" :decimals="0"/>
                                    </div>
                                    <div>本月系统访问量</div>
                                </div>
                            </el-col>
                        </el-row>
                    </el-card>

                </el-col>
            </el-row>
            <!--第二行-->
            <el-row :gutter="20">
                <el-col :span="16">

                    <el-card class="box-card" shadow="hover" :body-style="{ paddingBottom: '5px' }">
                        <div slot="header" class="clearfix">
                            <span>消费数据</span>
                        </div>
                        <div class="item">
                            <ve-histogram :data="chartData"></ve-histogram>
                        </div>
                    </el-card>

                </el-col>
                <el-col :span="8">
                    <el-card class="box-card" shadow="hover" :body-style="{ paddingBottom: '5px' }">
                        <div slot="header" class="clearfix">
                            <span>用户构成</span>
                        </div>
                        <div class="item">
                            <ve-pie :data="chartDataUser"></ve-pie>
                        </div>
                    </el-card>
                </el-col>
            </el-row>
            <el-row :gutter="20">
                <el-col :span="8">

                    <el-card class="box-card" shadow="hover">
                        <div slot="header" class="clearfix">
                            <span>留言评价</span>
                        </div>
                        <div class="item"></div>
                    </el-card>

                </el-col>
                <el-col :span="8">

                    <el-card class="box-card" shadow="hover">
                        <div slot="header" class="clearfix">
                            <span>内存监控</span>
                        </div>
                        <div class="item"></div>
                    </el-card>
                </el-col>
                <el-col :span="8">
                    <el-card class="box-card" shadow="hover">
                        <div slot="header" class="clearfix">
                            <span>磁盘监控</span>
                        </div>
                        <div class="item"></div>
                    </el-card>
                </el-col>
            </el-row>

        </div>
    </d2-container>
</template>

<script>

  export default {
    name: 'index',
    data () {
      return {
        chartData: {
          columns: ['日期', '访问用户', '下单用户', '下单率'],
          rows: [
            { '日期': '1/1', '访问用户': 1393, '下单用户': 1093, '下单率': 0.32 },
            { '日期': '1/2', '访问用户': 3530, '下单用户': 3230, '下单率': 0.26 },
            { '日期': '1/3', '访问用户': 2923, '下单用户': 2623, '下单率': 0.76 },
            { '日期': '1/4', '访问用户': 1723, '下单用户': 1423, '下单率': 0.49 },
            { '日期': '1/5', '访问用户': 3792, '下单用户': 3492, '下单率': 0.323 },
            { '日期': '1/6', '访问用户': 4593, '下单用户': 4293, '下单率': 0.78 }
          ]
        },
        chartDataUser: {
          columns: ['日期', '访问用户'],
          rows: [
            { '日期': '1/1', '访问用户': 1393 },
            { '日期': '1/2', '访问用户': 3530 },
            { '日期': '1/3', '访问用户': 2923 },
            { '日期': '1/4', '访问用户': 1723 },
            { '日期': '1/5', '访问用户': 3792 },
            { '日期': '1/6', '访问用户': 4593 }
          ]
        },
        chartSettingsVistor: {
          dimension: ['日期']
        },
        chartDataVistor: {
          columns: ['日期', '访问用户'],
          rows: [
            { '日期': '1', '访问用户': 1393 },
            { '日期': '2', '访问用户': 3530 },
            { '日期': '3', '访问用户': 2923 },
            { '日期': '4', '访问用户': 1723},
            { '日期': '5', '访问用户': 3792},
            { '日期': '6', '访问用户': 4593 }
          ]
        }
      }
    }, mounted () {
      // 加载完成后显示提示
      //this.showInfo()
    }, methods: {}
  }
</script>

<style lang="scss" scoped>

    .el-row {
        margin-bottom: 20px;

        &:last-child {
            margin-bottom: 0;
        }
    }

    .countup {
        font-size: 40px;
    }


</style>