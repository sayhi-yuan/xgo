package models

import "time"

// Admin  人员表
type Admin struct {
	ID            int64     `gorm:"column:id"`
	Username      string    `gorm:"column:username"`      //  姓名
	DepartmentId  int64     `gorm:"column:department_id"` //  部门id 关联部门表
	RoleId        int64     `gorm:"column:role_id"`       //  角色id 关联角色表
	Phone         string    `gorm:"column:phone"`         //  联系电话
	Email         string    `gorm:"column:email"`         //  邮箱
	Remark        string    `gorm:"column:remark"`        //  备注
	Status        int64     `gorm:"column:status"`        //  账号状态 1启用 0停用
	Password      string    `gorm:"column:password"`      //  账号密码
	UpdatedTime   time.Time `gorm:"column:updated_time"`  //  修改时间
	IsDelete      int64     `gorm:"column:is_delete"`     //  是否删除 0否1是
	TokenData     string    `gorm:"column:token_data"`
	LastLoginTime time.Time `gorm:"column:last_login_time"`
	AuthRoleIds   string    `gorm:"column:auth_role_ids"` //  可见角色id
	CreatedUser   int64     `gorm:"column:created_user"`  //  创建者
	UpdatedUser   int64     `gorm:"column:updated_user"`  //  最后一次更新人
	CreatedTime   time.Time `gorm:"column:created_time"`  //  创建时间
}

func (Admin) TableName() string {
	return "admin"
}
