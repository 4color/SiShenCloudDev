<template>
    <d2-container type="full" class="page">

        <el-button slot="header" icon="el-icon-circle-plus-outline" type="primary" style="margin-bottom: 5px; margin-left: 2px">添加新版本</el-button>

        <div style="position: absolute; top:40px; left: 200px">
            <el-form ref="form" :model="form" label-width="120px">
                <el-form-item label="应用ID">
                    app-bi
                </el-form-item>
                <el-form-item label="应用名称">
                    四神BI系统
                </el-form-item>
                <el-form-item label="最新版本">
                    V0.2
                </el-form-item>
                <el-form-item label="接入地址">
                    http://www.baidu.com
                </el-form-item>
                <el-form-item label="提交人">
                    四神
                </el-form-item>
                <el-form-item label="创建时间">
                    2019-11-12
                </el-form-item>
                <el-form-item label="是否有效">
                    有效
                </el-form-item>
                <el-form-item label="是否审核">
                    已审核
                </el-form-item>
                <el-form-item label="可以使用的角色">
                    <el-checkbox-group v-model="checkRoles">
                        <el-checkbox disabled v-for="city in roles" :label="city" :key="city">{{city}}</el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="依赖的应用">
                    <el-checkbox-group v-model="checkedCities" >
                        <el-checkbox disabled v-for="city in cities" :label="city" :key="city">{{city}}</el-checkbox>
                    </el-checkbox-group>
                </el-form-item>

            </el-form>
        </div>


        <div style="position: absolute; top:40px; left: 100px">
            <el-avatar :size="80" :src="circleUrl"></el-avatar>
        </div>
    </d2-container>


</template>

<script>
  const cityOptions = ['上海', '北京', '广州', '深圳']
  const allrows = ['游客', '一般角色', '包年VIP', '永久VIP']
  export default {
    name: 'appdetail',
    data () {
      return {
        dialogTableVisible: false,
        checkedCities: ['上海', '北京'],
        cities: cityOptions,
        roles: allrows,
        checkRoles: ['包年VIP', '永久VIP'],
        form: {},
        circleUrl: 'https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png',
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
    .myalert {
        font-size: 30px;
    }
</style>
