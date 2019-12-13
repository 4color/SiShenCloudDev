<template>
    <d2-container>
        <el-select slot="header" v-model="year" style="float: right; padding: 3px 0">
            <el-option
                    v-for="item in 2019" v-if="item>2017"
                    :key="item"
                    :label="item"
                    :value="item">
            </el-option>
        </el-select>

        <div class="inner">
            <ve-histogram :data="chartData" :settings="chartSettings" height="700px"></ve-histogram>
        </div>
    </d2-container>
</template>

<script>
  export default {
    name: 'reportindex',
    data () {
      this.chartSettings = {
        showLine: ['退款']
      }
      return {
        year: 2019,
        chartData: {
          columns: ['月份', '收入', '退款'],
          rows: [
            { '月份': '1月', '收入': 1393, '退款': 12 },
            { '月份': '2月', '收入': 3530, '退款': 13 },
            { '月份': '3月', '收入': 2923, '退款': 15 },
            { '月份': '4月', '收入': 1723, '退款': 20 },
            { '月份': '5月', '收入': 3792, '退款': 69 },
            { '月份': '6月', '收入': 4593, '退款': 23 },
            { '月份': '7月', '收入': 4593, '退款': 34 },
            { '月份': '8月', '收入': 4593, '退款': 21 },
            { '月份': '9月', '收入': 4593, '退款': 12 },
            { '月份': '10月', '收入': 4593, '退款': 0 },
            { '月份': '11月', '收入': 6593, '退款': 330 },
            { '月份': '12月', '收入': 4593, '退款': 7 }

          ]
        }
      }
    }
  }
</script>

<style lang="scss" scoped>
    .inner {
        position: absolute;
        top: 20px;
        right: 20px;
        bottom: 20px;
        left: 20px;
    }
</style>