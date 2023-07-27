<template>
  <div class="app-container">
    <div class="filter-container" align="right" style="margin-bottom: 5px;">
      <el-input v-model="search" placeholder="请输入关键字" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-button class="filter-item" style="margin-left: 5px;" type="primary" icon="el-icon-circle-plus-outline" @click="userCreateDialog">
        增加用户
      </el-button>
    </div>
    <el-table
      v-loading="listLoading"
      :data="searchUserData"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
      height="850"
    >
      <el-table-column align="center" label="序号" width="60">
        <template slot-scope="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="用户名" align="center" width="200">
        <template slot-scope="scope">
          <el-link type="primary" :underline="false">{{ scope.row.username }}</el-link>
        </template>
      </el-table-column>
      <el-table-column label="用户中文名" align="center" width="200">
        <template slot-scope="scope">
          <el-link type="primary" :underline="false">{{ scope.row.cnname }}</el-link>
        </template>
      </el-table-column>
      <el-table-column align="center" label="创建时间">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.CreatedAt|dateToStr }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="更新时间">
        <template slot-scope="scope">
          <i class="el-icon-time" />
          <span>{{ scope.row.UpdatedAt|dateToStr }}</span>
        </template>
      </el-table-column>
      <el-table-column label="当前状态" align="center" width="200" class-name="small-padding fixed-width">
        <template slot-scope="{row}">
          <el-tooltip class="item" effect="dark" content="点击修改状态" placement="top">
            <el-button v-if="row.status =='2'" size="mini" type="success" @click="changeUserClick(row, '1')">
              启用
            </el-button>
          </el-tooltip>
          <el-tooltip class="item" effect="dark" content="点击修改状态" placement="top">
            <el-button v-if="row.status =='1'" size="mini" type="info" @click="changeUserClick(row, '2')">
              禁用
            </el-button>
          </el-tooltip>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="420" class-name="small-padding fixed-width">
        <template slot-scope="{row}">
          <el-button size="mini" type="primary" @click="editUserButton(row)">
            权限编辑
          </el-button>
          <el-button size="mini" type="warning" @click="resetDiglog(row)">
            重置
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- reset user dialog -->
    <el-dialog
      width="30%"
      title="重置密码"
      :visible.sync="diglogresetUserPassword"
      append-to-body
    >
      <span style="color:red">确定要重置用户密码为 passwd@123 吗?</span>
      <span slot="footer" class="dialog-footer">
        <el-form ref="form" :model="tempUserData">
          <el-form-item>
            <el-button size="mini" @click="diglogresetUserPassword = false">取 消</el-button>
            <el-button type="primary" size="mini" @click="resetSubmit()">确 定</el-button>
          </el-form-item>
        </el-form>
      </span>
    </el-dialog>
    <el-dialog title="增加用户" :visible.sync="dialogAddUser" width="35%">
      <el-form ref="dataForm" :rules="rules" :model="addUserForm" label-position="left" label-width="70px" style="width: 450px; margin-left:50px;" size="medium">
        <el-form-item label="用户名" prop="username">
          <el-input v-model="addUserForm.username" />
        </el-form-item>
        <el-form-item label="中文名" prop="cnname">
          <el-input v-model="addUserForm.cnname" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="addUserForm.email" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button size="mini" @click="dialogAddUser = false">取消</el-button>
        <el-button type="primary" size="mini" @click="createUserSubmit()">确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import moment from 'moment'
import { getUserList, changeSysUserStatus, createUser, getExistUser, resetUserPassword } from '@/api/sys-user'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    },
    dateToStr(value) {
      if (typeof value === 'string') {
        return moment(value).format(' YYYY-MM-DD HH:mm:ss')
      } else {
        return ' 暂无'
      }
    }
  },
  data() {
    var checkUserName = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('用户名不能为空'))
      }
      getExistUser({ 'username': value }).then(response => {
        if (response.msg !== '1') {
          return callback(new Error('用户名已存在'))
        } else {
          callback()
        }
      })
    }
    return {
      tempUserData: {},
      userTableData: null,
      listLoading: true,
      formloading: true,
      dialogEditSubmit: false,
      diglogresetUserPassword: false,
      dialogAddUser: false,
      userId: '',
      roleNameData: [],
      existRoleData: [],
      roleValue: [],
      tempRoleData: {},
      addUserForm: {
        cnname: '',
        username: '',
        email: ''
      },
      value: [],
      listQuery: {
        Page: 1,
        PageSize: 100
      },
      rules: {
        cnname: [{ required: true, message: 'cnname is required', trigger: 'blur' }],
        username: [{ required: true, validator: checkUserName, trigger: 'blur' }],
        email: [{ required: true, message: 'email is required', trigger: 'blur' }]
      },
      search: '',
      timer: ''
    }
  },
  computed: {
    searchUserData: function() {
      var search = this.search
      if (search) {
        return this.userTableData.filter(function(dataNews) {
          return Object.keys(dataNews).some(function(key) {
            return String(dataNews[key]).toLowerCase().indexOf(search) > -1
          })
        })
      }
      return this.userTableData
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    reload() {
      this.$router.go(0)
    },
    fetchData() {
      this.listLoading = false
      getUserList(this.listQuery).then(response => {
        this.userTableData = response.data.list
        this.listLoading = false
      })
    },
    resetTemp() {
      this.roleNameData = []
      this.existRoleData = []
      this.roleValue = []
      this.addUserForm = {
        cnname: '',
        username: '',
        email: ''
      }
    },
    changeUserClick(row, status) {
      changeSysUserStatus({ id: row.ID }).then(() => {
        this.$message({
          message: '提交成功',
          type: 'success',
          showClose: true,
          duration: 1000
        })
      }).catch(error => {
        this.$message({
          message: '提交失败',
          type: 'error',
          showClose: true,
          duration: 1000,
          error: error
        })
      })
      row.status = status
    },
    resetDiglog(row) {
      this.diglogresetUserPassword = true
      this.tempUserData = Object.assign({}, row)
    },
    resetSubmit() {
      resetUserPassword({ username: this.tempUserData.username }).then(() => {
        this.$message({
          message: '提交成功',
          type: 'success',
          showClose: true,
          duration: 1000
        })
      }).catch(error => {
        this.$message({
          message: '提交失败',
          type: 'error',
          showClose: true,
          duration: 1000,
          error: error
        })
      })
      this.diglogresetUserPassword = false
    },
    userCreateDialog() {
      this.resetTemp()
      this.dialogAddUser = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    createUserSubmit() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          createUser(this.addUserForm).then(() => {
            this.$message({
              message: '提交成功',
              type: 'success',
              showClose: true,
              duration: 1000
            })
          })
          this.dialogAddUser = false
          this.timer = setTimeout(this.reload, 2000)
        }
      })
    }
  }
}
</script>

<style lang="scss">
  //  .edit-user-content .el-transfer-panel {
  //   width: 210px;
  //   margin-left: 40px;
  // }

  // .edit-user-content .el-transfer-panel__body {
  //   height: 450px;
  // }

  // .edit-user-content .el-transfer-panel__list.is-filterable {
  //   height: 400px;
  // }

  // .edit-user-content .el-transfer__buttons {
  //   width: 80px;
  //   margin-left: -10px;
  // }

  // .edit-user-content .el-transfer__button {
  //   margin-left: 10px;
  // }

  // .filter-item .el-input__inner {
  //   height: 36px;
  //   padding-top: 8px;
  //   resize: none;
  // }
</style>
