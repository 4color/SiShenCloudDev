<template>
    <d2-container type="full" class="page">


        <el-button slot="header" icon="el-icon-printer" style="margin-bottom: 5px; margin-left: 2px">打印</el-button>
        <el-button slot="header" icon="el-icon-money" @click="dialogTableVisible = true" type="primary"
                   style="margin-bottom: 5px; margin-left: 2px">消费记录
        </el-button>
        <el-button slot="header" icon="el-icon-document" style="margin-bottom: 5px; margin-left: 2px">操作日志</el-button>

        <el-form ref="form" :model="form" label-width="120px" style="padding: 30px">
            <el-form-item label="用户名">
                zhuyf
            </el-form-item>
            <el-form-item label="用户类型">
                VIP用户
            </el-form-item>
            <el-form-item label="行政区">
                浙江省
            </el-form-item>
            <el-form-item label="公司名称">
                浙江臻善科技有限公司
            </el-form-item>
            <el-form-item label="手机号">
                18738888333
            </el-form-item>
            <el-form-item label="Email">
                18738888333@qq.com
            </el-form-item>
            <el-form-item label="通信地址">
               某某省 某某市 经二路
            </el-form-item>
            <el-form-item label="注册时间">
                2019-11-11 12:12:12
            </el-form-item>
            <el-form-item label="已开通功能">
                <el-checkbox-group v-model="checkedCities" @change="handleCheckedCitiesChange">
                    <el-checkbox disabled v-for="city in cities" :label="city" :key="city">{{city}}</el-checkbox>
                </el-checkbox-group>
            </el-form-item>

        </el-form>


        <el-dialog title="消费记录" :visible.sync="dialogTableVisible">
            <el-table :data="gridData">

                <el-table-column property="date" label="消费日期" width="100"></el-table-column>
                <el-table-column property="orderno" label="订单编号" width="150"></el-table-column>
                <el-table-column property="name" label="项目名称" width="200"></el-table-column>
                <el-table-column property="money" label="金额" width="100"></el-table-column>
                <el-table-column property="status" label="状态"></el-table-column>
            </el-table>
        </el-dialog>


        <div style="position: absolute; top:40px; left: 400px" >
            <el-avatar :size="80" :src="circleUrl"></el-avatar>
        </div>
        <div style="position: absolute; top:40px; right: 30px" class="myalert">
            <div class="el-alert el-alert--success is-light">

                已消费500元
            </div>
        </div>
    </d2-container>


</template>

<script>
  const cityOptions = ['上海', '北京', '广州', '深圳']
  export default {
    name: 'userdetail',
    data () {
      return {
        dialogTableVisible: false,
        checkedCities: ['上海', '北京'],
        cities: cityOptions,
        form: {},
        circleUrl: "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png",
        gridData: [{
          date: '2016-05-02',
          orderno: '201902222222',
          name: '开通VIP',
          money: '+500',
          status: '已结算'
        }, {
          date: '2016-05-04',
          orderno: '201902222222',
          name: '开通商城',
          money: '+500',
          status: '已结算'
        }, {
          date: '2016-05-01',
          orderno: '201902222222',
          name: '开通GIS功能',
          money: '+500',
          status: '有退款'
        }, {
          date: '2016-05-03',
          name: '退款GIS功能',
          money: '-500',
          status: '已结算'
        }],
      }
    }
  }
</script>

<style scoped>
 .myalert
 {
     font-size: 30px;
 }
</style>
