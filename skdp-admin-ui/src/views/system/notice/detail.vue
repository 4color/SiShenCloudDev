<template>
    <d2-container type="full" class="page">

        <el-button slot="header" icon="el-icon-document" style="margin-bottom: 5px; margin-left: 2px">返回</el-button>

        <el-form ref="form" :model="form" label-width="120px" style="padding: 30px">
            <el-form-item label="公告标题">
                zhuyf
            </el-form-item>
            <el-form-item label="发布时间">
                VIP用户
            </el-form-item>
            <el-form-item label="浏览次数">
                1234
            </el-form-item>
            <el-form-item label="内容">
                <div v-html="form.content"></div>
            </el-form-item>


        </el-form>

    </d2-container>


</template>

<script>

  export default {
    name: 'userdetail',
    data () {
      return {
        form: {
          content:'<table width="600" border="0" cellspacing="1"  style="text-align:left;"><tbody><tr><td colspan="4"><b>公告信息：</b></td></tr><tr><td class="title" width="128">采购项目名称</td><td colspan="3" width="430">中华人民共和国舟山海关空调系统维修改造采购项目招标文件更改公告</td></tr><tr><td class="title">品目</td><td colspan="3"><p><p>工程/修缮工程/房屋修缮</p></p></td></tr><tr><td class="title">采购单位</td><td colspan="3">中华人民共和国舟山海关</td></tr><tr><td class="title">行政区域</td><td width="168">舟山市</td><td class="title" width="128">公告时间</td><td width="168">2019年12月09日  15:14</td></tr><tr><td class="title">首次公告日期</td><td width="168">2019年08月13日</td><td class="title" width="128">更正日期</td><td width="168">2019年08月13日</td></tr><tr><td colspan="4"><b>联系人及联系方式：</b></td></tr><tr><td class="title">项目联系人</td><td colspan="3">褚先生</td></tr><tr><td class="title">项目联系电话</td><td colspan="3">0580-8133766</td></tr><tr><td class="title" width="128">采购单位</td><td width="430" colspan="3">中华人民共和国舟山海关</td></tr><tr><td class="title">采购单位地址</td><td colspan="3">舟山市临城海景道599号</td></tr><tr><td class="title">采购单位联系方式</td><td colspan="3">褚先生，0580-8133766</td></tr><tr><td class="title">代理机构名称</td><td colspan="3">深圳市国信招标有限公司舟山分公司</td></tr><tr><td class="title">代理机构地址</td><td colspan="3">浙江省舟山市昌国路232号中楼202</td></tr><tr><td class="title">代理机构联系方式</td><td colspan="3">朱小姐，任小姐；13857236444，13567673203</td></tr></tbody></table>'
        }
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
