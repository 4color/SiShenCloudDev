<template>
    <d2-container type="full" class="page">

        <el-button slot="header" icon="el-icon-check" type="primary" style="margin-bottom: 5px; margin-left: 2px">保存</el-button>

        <div style="position: absolute; top:40px; left: 260px">
            <el-form ref="form" :model="form" label-width="120px">
                <el-form-item label="应用ID">
                    <el-input placeholder="请输入内容"></el-input>
                </el-form-item>
                <el-form-item label="应用名称">
                    <el-input placeholder="请输入内容"></el-input>
                </el-form-item>
                <el-form-item label="最新版本">
                    <el-input placeholder="0.1"></el-input>
                </el-form-item>
                <el-form-item label="接入地址">
                    <el-input placeholder="请输入内容"></el-input>
                </el-form-item>
                <el-form-item label="提交人">
                    <el-input placeholder="请输入内容"></el-input>
                </el-form-item>
                <el-form-item label="创建时间">
                    <el-date-picker
                            type="date"
                            placeholder="选择日期">
                    </el-date-picker>
                </el-form-item>
                <el-form-item label="是否有效">
                    <el-switch
                            active-text="有效"
                            inactive-text="无效">
                    </el-switch>
                </el-form-item>
                <el-form-item label="是否审核">
                    <el-switch
                            active-text="已审核"
                            inactive-text="未审核">
                    </el-switch>
                </el-form-item>
                <el-form-item label="可以使用的角色">
                    <el-checkbox-group v-model="checkRoles">
                        <el-checkbox  v-for="city in roles" :label="city" :key="city">{{city}}</el-checkbox>
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="依赖的应用">
                    <el-checkbox-group v-model="checkedCities" >
                        <el-checkbox  v-for="city in cities" :label="city" :key="city">{{city}}</el-checkbox>
                    </el-checkbox-group>
                </el-form-item>

                <el-form-item label="应用介绍">
                    <d2-quill
                            style="min-height: 200px; margin-bottom: 20px;"
                            v-model="form.intro"/>
                </el-form-item>

            </el-form>
        </div>


        <div style="position: absolute; top:40px; left: 60px">
            <el-upload
                    class="avatar-uploader"
                    action="https://jsonplaceholder.typicode.com/posts/"
                    :show-file-list="false"
                    :on-success="handleAvatarSuccess"
                    :before-upload="beforeAvatarUpload">
                <img v-if="imageUrl" :src="imageUrl" class="avatar">
                <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>
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
        form: {
          intro:''
        },
        imageUrl: '',
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
    },
    methods: {
      handleAvatarSuccess(res, file) {
        this.imageUrl = URL.createObjectURL(file.raw);
      },
      beforeAvatarUpload(file) {
        const isJPG = file.type === 'image/jpeg';
        const isLt2M = file.size / 1024 / 1024 < 2;

        if (!isJPG) {
          this.$message.error('上传头像图片只能是 JPG 格式!');
        }
        if (!isLt2M) {
          this.$message.error('上传头像图片大小不能超过 2MB!');
        }
        return isJPG && isLt2M;
      }
    }
  }
</script>

<style scoped>
    .avatar-uploader .el-upload {
        border: 1px dashed #d9d9d9;
        border-radius: 6px;
        cursor: pointer;
        position: relative;
        overflow: hidden;
    }
    .avatar-uploader .el-upload:hover {
        border-color: #409EFF;
    }
    .avatar-uploader-icon {
        font-size: 28px;
        color: #8c939d;
        width: 100px;
        height: 100px;
        line-height: 178px;
        text-align: center;
    }
    .avatar {
        width: 178px;
        height: 178px;
        display: block;
    }
</style>
