<template>
    <d2-container type="full" class="page">

        <el-input slot="header" placeholder="请输入查询内容" style="width: 150px;margin-left: 5px"></el-input>
        <el-button slot="header" style="margin-bottom: 5px; margin-left: 2px">查询</el-button>

        <el-button slot="header" icon="el-icon-plus" type="success" style="margin-left: 2px"
                   @click="dialogVisible = true">新建顶级组
        </el-button>
        <el-button slot="header" icon="el-icon-plus" type="primary" style="margin-left: 2px"
                   @click="dialogVisible = true">新增字典
        </el-button>


        <SplitPane :min-percent='17' :default-percent='17' split="vertical">
            <template slot="paneL">
                <el-tree :data="dictdata" :props="defaultProps" :auto-expand-parent="true"
                         @node-contextmenu="rihgtClick" @node-click="nodeClick">
                    <span slot-scope="{ node, data }">
      <i :class="data.icon"></i>
      <span style="padding-left: 4px; font-size: 12px">{{node.label}}</span>
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
                            prop="sheng"
                            label="省级"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="xzqmc"
                            label="行政区名称"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            prop="xzqdm"
                            label="区县代码">
                    </el-table-column>
                    <el-table-column
                            prop="enable"
                            label="有效"
                            width="100"
                            :filters="[{ text: '有效', value: true }, { text: '无效', value: false }]"
                            :filter-method="filterTag"
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
                                    title="字典扩展"
                                    style="padding: 5px;"
                                    @click="handleDelete(scope.$index, scope.row)">
                            </el-button>

                        </template>
                    </el-table-column>
                </el-table>
            </template>
        </SplitPane>


        <!--菜单右键-->
        <div v-show="menuVisible" class="tree_menu" id="perTreeMenu" :style="{...rightMenu}">
            <ul id="menu">
                <li id="btnAddFolder"><i class="el-icon-folder"></i>
                    <el-link class="linkbtn" :underline="false" :disabled="linkbtnDisableFolder"
                             @click="showAddFolder=true">添加分组
                    </el-link>
                </li>
                <li id="btnDeleteFolder"><i class="el-icon-delete" style="color:red"></i>
                    <el-link class="linkbtn" :disabled="linkbtnDisableFolder"
                             :underline="false">删除分组
                    </el-link>
                </li>
                <el-divider></el-divider>
                <li tabindex="-1"><i class="el-icon-collection"></i>
                    <el-link class="linkbtn" :disabled="linkbtnDisableAddField"
                             :underline="false" @click="showAddField=true">添加字段
                    </el-link>
                </li>
                <li tabindex="-1"><i class="el-icon-minus" style="color:red"></i>
                    <el-link class="linkbtn" :disabled="linkbtnDisableDeleteField"
                             :underline="false">删除字段
                    </el-link>
                </li>
                <el-divider></el-divider>
                <li tabindex="-1"><i class="el-icon-notebook-2"></i>
                    <el-link class="linkbtn" :disabled="linkbtnDisableAddEnum"
                             :underline="false" @click="showAddEnum=true">添加字典
                    </el-link>
                </li>
                <li tabindex="-1"><i class="el-icon-delete-solid" style="color:red"></i>
                    <el-link class="linkbtn" :disabled="linkbtnDisableDeleteEnum"
                             :underline="false">删除字典
                    </el-link>
                </li>
            </ul>
        </div>
        <!--新增分组-->
        <el-dialog
                title="新增分组"
                :visible.sync="showAddFolder"
                width="30%"
        >
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="新增分组">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
    <el-button @click="showAddFolder = false">取 消</el-button>
    <el-button type="primary" @click="showAddFolder = false">确 定</el-button>
  </span>
        </el-dialog>

        <!--新增字段-->
        <el-dialog
                title="新增字段"
                :visible.sync="showAddField"
                width="30%"
        >
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="字段名称">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="字段代码">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
    <el-button @click="showAddField = false">取 消</el-button>
    <el-button type="primary" @click="showAddField = false">确 定</el-button>
  </span>
        </el-dialog>

        <!--新增字典-->
        <el-dialog
                title="新增字典"
                :visible.sync="showAddEnum"
                width="30%"
        >
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="字典名称">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="字典值">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="顺序">
                    <el-input-number  :min="1" :max="100"></el-input-number>
                </el-form-item>
                <el-form-item label="状态">
                    <el-switch
                            v-model="value"
                            active-color="#13ce66"
                            inactive-color="#EFEFEF">
                    </el-switch>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
    <el-button @click="showAddEnum = false">取 消</el-button>
    <el-button type="primary" @click="showAddEnum = false">确 定</el-button>
  </span>
        </el-dialog>
    </d2-container>
</template>

<script>
  export default {
    name: 'system-dict',
    data () {
      return {
        menuVisible: false,
        searchType: '3',
        dialogVisible: false,
        //新增分组
        showAddFolder: false,
        //新增分组
        showAddField: false,
        //新增字典
        showAddEnum: false,
        linkbtnDisableFolder: true,
        linkbtnDisableAddField: true,
        linkbtnDisableDeleteField: true,
        linkbtnDisableAddEnum: true,
        linkbtnDisableDeleteEnum: true,
        rightMenu: {},
        form: {},
        data: [
          {
            sheng: '浙江省',
            xzqmc: '拱墅区',
            xzqdm: '330111',
            'enable': true
          },
          {
            sheng: '浙江省',
            xzqmc: '拱墅区',
            xzqdm: '330111',
            'enable': true
          },
          {
            sheng: '浙江省',
            xzqmc: '拱墅区',
            xzqdm: '330111',
            'enable': true
          },
          {
            sheng: '浙江省',
            xzqmc: '拱墅区',
            xzqdm: '330111',
            'enable': false
          }
        ],
        dictdata: [
          {
            lable: '通用字典',
            value: '33',
            type: 'folder',
            icon: 'el-icon-folder',
            children: [
              {
                lable: '土地用途',
                value: '4401',
                type: 'field',
                icon: 'el-icon-coin',
                children: [
                  {
                    lable: '工业用地',
                    value: '4401',
                    type: 'enum',
                    icon: 'el-icon-tickets'
                  },
                  {
                    lable: '住宿用地',
                    value: '4402',
                    type: 'enum',
                    icon: 'el-icon-tickets'
                  }
                ]
              },
              {
                lable: '深圳市',
                value: '4402',
                icon: 'el-icon-tickets\n'
              }
            ]
          },
          {
            lable: '福建省',
            value: '35',
            icon: 'el-icon-folder'
          },
          {
            lable: '广东省',
            value: '44',
            icon: 'el-icon-folder',
            children: [
              {
                lable: '广州市',
                value: '4401',
                icon: 'el-icon-tickets'
              },
              {
                lable: '深圳市',
                value: '4402',
                icon: 'el-icon-tickets'
              }
            ]
          },
          {
            lable: '浙江省',
            value: '33',
            icon: 'el-icon-folder'
          }
        ],
        defaultProps: {
          children: 'children',
          label: 'lable'
        }
      }
    },
    mounted () {
    },
    methods: {
      indexMethod (index) {
        return index + 1
      },
      handleEdit (index, row) {
        console.log(index, row)
      },
      handleDelete (index, row) {
        console.log(index, row)
      },
      jumpUrl (row) {
        this.$router.push('/department/xzqxz/' + row.xzqdm)
      },
      filterTag (value, row) {
        return row.tag === value
      },
      //树的右键
      rihgtClick (e, object, value, element) {

        //判断数据类型
        //禁用menu下面所有的组件
        this.linkbtnDisableFolder = true
        this.linkbtnDisableAddField = true
        this.linkbtnDisableDeleteField = true
        this.linkbtnDisableAddEnum = true
        this.linkbtnDisableDeleteEnum = true

        if (object.type === 'folder') {
          this.linkbtnDisableFolder = false
          this.linkbtnDisableAddField = false
        }

        if (object.type === 'field') {
          this.linkbtnDisableDeleteField = false
          this.linkbtnDisableAddEnum = false
        }

        if (object.type === 'enum') {
          this.linkbtnDisableAddEnum = false
          this.linkbtnDisableDeleteEnum = false
        }

        this.rightMenu = { top: e.pageY + 'px', left: e.pageX + 'px' }
        this.menuVisible = true
        const self = this
        document.onclick = function (ev) {
          if (ev.target !== document.getElementById('perTreeMenu')) {
            self.menuVisible = false
          }
        }
      },
      nodeClick (e, object, value, element) {
        this.menuVisible = false
      }
    }
  }
</script>

<style lang="scss" scoped>
    .tree_menu {
        position: fixed;
        display: block;
        width: 180px;
        z-index: 20000;
        background-color: #fff;
        padding: 5px 0;
        border: 1px solid #ebeef5;
        border-radius: 4px;
        box-shadow: 0 2px 12px 0 rgba(0, 0, 0, .1);

        span {
            font-size: 12px;
        }

    }

    ul {
        margin: 0;
        padding: 0;
    }

    ul li {
        list-style: none;
        margin: 0;
        padding: 0 15px;
        font-size: 14px;
        line-height: 30px;
    }

    ul li:hover {
        background-color: #ebeef5
    }

    .el-divider--horizontal {
        margin: 5px 0px !important;
    }

    .linkbtn {
        padding-left: 3px;
    }
</style>
