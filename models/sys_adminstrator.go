package models 

import (
	_"fmt"
	_"net/http"


	_"github.com/gin-gonic/gin"

	"ginpro/global"
	
	
)

// sys_adminstrators 用户类
type SysAdminstrator struct {
	Id        		int     `gorm:"column:id;primaryKey" json:"id" xml:"id"` // 用户id
    UserName      	string  `gorm:"column:user_name" json:"user_name" xml:"user_name" form:"user_name" binding:"required"` // 用户名
	MobilePhone     string  `gorm:"column:mobile_phone" json:"mobile_phone"  xml:"mobile_phone" form:"mobile_phone" binding:"required"` // 移动电话
	Phone     		string  `gorm:"column:phone" json:"phone"  xml:"phone" form:"phone" binding:"required"` // 坐机电话
	Email     		string  `gorm:"column:email" json:"email"  xml:"email" form:"email" binding:"required"` // 用户邮箱
	Role			string  `gorm:"column:role" json:"role" xml:"role" ` // 角色
	Password  		string  `gorm:"column:password" json:"-" xml:"-" form:"password" binding:"required"` // 用户密码
	Status          int     `gorm:"column:status" json:"status" xml:"status"` // 状态 0：正常，1：禁用 
	CreatedAt 		int64   `gorm:"column:created_at" json:"created_at" xml:"created_at"` // 创建时间戳
	UpdatedAt 		int64   `gorm:"column:updated_at" json:"updated_at" xml:"updated_at"`  // 创建时间戳
}




// SysLoginReq 登录请求参数类
type SysLoginReq struct {
	UserName   string `form:"user_name" json:"user_name" xml:"user_name" binding:"required" validate:"required"`
	Password   string `form:"password" json:"password" xml:"password" binding:"required" validate:"required"`
	Role      string `json:"role"`
}

func (SysAdminstrator) TableName() string {
	return "sys_adminstrator"
}

func  LoginCheck(sysLoginReq *SysLoginReq) (resultBool bool,  loginResult SysAdminstrator,  err error) {
		resultBool = true
       // res := map[string]interface{}{}
	   err = global.DB().Model(&SysAdminstrator{}).Where("user_name = ? AND password = ?", sysLoginReq.UserName , sysLoginReq.Password ).First(&loginResult).Error
	   if err != nil {
	    	resultBool = false
	   }
	   return 
}

