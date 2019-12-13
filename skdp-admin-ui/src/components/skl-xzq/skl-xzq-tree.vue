<template>


    <el-select placeholder="请选择行政区" v-model="form.xzqmc" style="width: 560px" ref="selectReport">
        <el-option  :value="form.xzqdm" :label="form.xzqmc" style="width: 560px;height:200px;overflow: auto;background-color:#fff">
            <el-tree :data="xzqdata" :props="defaultProps" @node-click="handleNodeClick" :auto-expand-parent="true"></el-tree>
        </el-option>
    </el-select>
</template>

<script>
  export default {
    name: 'skl-xzq-tree',
    data () {
      return {
        searchType: '3',
        form: {
          xzqdm:"",
          xzqmc:""
        },
        xzqdata: [
          {
            label: '全国',
            value: '00',
            children: [
              {
                label: '浙江省',
                value: '33'
              },
              {
                label: '福建省',
                value: '35'
              },
              {
                label: '广东省',
                value: '44',
                children: [
                  {
                    label: '广州市',
                    value: '4401'
                  },
                  {
                    label: '深圳市',
                    value: '4402'
                  }
                ]
              },
              {
                label: '浙江省',
                value: '33'
              }
            ]
          }
        ],
        defaultProps: {
          children: 'children',
          label: 'label'
        }
      }
    },
    mounted () {
    },
    methods: {

      handleNodeClick: function (node) {
        console.log(node, 'node')
        if (node.children) {

        } else {
          this.form.xzqdm = node.value;
          this.form.xzqmc = node.label;
          this.$refs.selectReport.blur()
        }
      }
    }
  }
</script>
