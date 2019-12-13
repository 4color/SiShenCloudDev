<template>
    <d2-container type="full" class="page">

        <el-button slot="header" icon="el-icon-check" type="primary" style="margin-bottom: 5px; margin-left: 2px">保存</el-button>

        <div style="position: absolute; top:40px; left: 200px">
            <el-form ref="form" :model="form" label-width="120px">
                <el-form-item label="系统应用行政区">
                    <skl-xzq-tree></skl-xzq-tree>

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
    name: 'systeconfig',
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
