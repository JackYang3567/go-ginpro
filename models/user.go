package models 

import (
	_"fmt"

	_"github.com/gin-gonic/gin"


	_"ginpro/global"
)
// User 用户类

type User struct {
	Id        		int     `gorm:"column:id;primaryKey" json:"id" xml:"id"` // 用户id
	TenantId  		string  `gorm:"column:tenant_id" json:"tenant_id" xml:" tenant_id" binding:"required"` // tenant_id
    UserName      	string  `gorm:"column:user_name" json:"user_name" xml:"user_name" form:"user_name" binding:"required"` // 用户名
	Uuid      		string  `gorm:"column:uuid" json:"uuid" xml:" uuid"` // 用户uuid
	MobilePhone     string  `gorm:"column:mobile_phone" json:"mobile_phone"  xml:"mobile_phone" form:"mobile_phone" binding:"required"` // 移动电话
	Email     		string  `gorm:"column:email" json:"email"  xml:"email" form:"email" binding:"required"` // 用户邮箱
	FirstName 		string  `gorm:"column:first_name" json:"first_name" xml:"first_name"` // 用户名
	LastName  		string  `gorm:"column:last_name" json:"last_name" xml:"last_name" ` // 用户名
	Gender			string  `gorm:"column:gender" json:"gender" xml:"gender" ` // 性别
	Password  		string  `gorm:"column:password" json:"password" xml:"password" form:"password" binding:"required"` // 用户密码
	CreatedAt 		int64   `gorm:"column:created_at" json:"created_at" xml:"created_at"` // 创建时间戳
	UpdatedAt 		int64   `gorm:"column:updated_at" json:"updated_at" xml:"updated_at"`  // 创建时间戳
}

/*
type User struct {
	Id         string `json:"userId"`
	Name       string `json:"userName"`
	Gender     string `json:"gender"`
	Phone      string `json:"userMobile"`
	Pwd        string `json:"pwd"`
	Permission string `json:"permission"`
}
*/

// EditUserReq 更新用户信息数据类
type EditUserReq struct {
	Id     string `json:"userId"`
	UserName   string `json:"user_name"`
	Gender string `json:"gender"`
}

// LoginReq 登录请求参数类
type LoginReq struct {
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
/*
// LoginCheck 登录验证 //(bool, SysAdminstrator, error)
func LoginCheck(c *gin.Context, loginReq LoginReq)  {
	result := global.NewResult(c)
//	resultBool := false
	
	if loginReq.UserName == "root" {
		sysAdminTab := new(SysAdminstrator)
        res := global.DB().Where("user_name = ? AND password = ?", loginReq.UserName, loginReq.Password).First(&sysAdminTab) 
        
		if res.Error != nil {
			result.Error(444,"Select error: "+ res.Error.Error())
			 
		}else{		 
			//result.Success(200,gin.H{"total":&res.RowsAffected, "data":&res.Value})
		//	resultBool = true
			result.Success(900, gin.H{"total":&res.RowsAffected, "data":&res.Value})
		//	return resultBool, &res.Value[0], res.Error
	    //	return  &res.Value, res.Error

		}
	}//else{
	//	data := Login{}
		
	//}
	//return resultBool, res.Value, res.Error
}
*/