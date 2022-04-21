package controllers

import (
	 _"encoding/json"
	"fmt"
	"net/http"
	_"strings"
	"time"

	"github.com/gin-gonic/gin"

	"ginpro/global"
	"ginpro/models"
	"ginpro/middleware"
)


type LoginController struct{}




// @Summary 注册用户
// @Description 输入用户名及密登录
// @Produce json
// @Param tenant_id path int true "承租户tenant_id"
// @Param email path string true "用户邮箱"
// @Param username path string true "登录名"
// @Param password path string true "登录密码"
// @Success 200 {object} swag.ResultContSingUp 成功后返回值
// @Router /singup [post]
func SingUp(c *gin.Context) {
	result := global.NewResult(c)
	data := models.Login{}
	err := c.ShouldBind(&data)
	if err == nil {
		data.Uuid      = global.CreateUUID()
		data.CreatedAt = time.Now().Unix() 
		data.UpdatedAt = time.Now().Unix() 
		data.Password  = global.Encrypt( c.PostForm("password") )

		err := global.DB().Create(&data).Error
		if err != nil {
			//1000操作数据库错误
			result.Error(1000, "Insert error: "+err.Error())
		 }else{
			result.Success(201,&data);
		 }
	}else{
		//444客户端没做验证
		result.Error(http.StatusNoContent,"Validation error: "+ err.Error())
	}
	return
}


// @Summary 系统超级管理登录
// @Description 输入用户名及密登录
// @Produce json
// @Param username path string true "登录名"
// @Param password path string true "登录密码"
// @Success 200 {object} swag.ResultContSingIn 成功后返回值
// @Router /login [post]
func LogIn(c *gin.Context) {
	result := global.NewResult(c)
	var sysLoginReq models.SysLoginReq
	if err :=  c.ShouldBind(&sysLoginReq); err == nil {
			//loginReq := models.SysAdminstrator{}
			sysLoginReq.UserName = c.PostForm("user_name")
			sysLoginReq.Password  = global.Encrypt( c.PostForm("password") )			
			sysLoginReq.Role = "sys$admin"
			isPass, sysAdmin, err :=  models.LoginCheck(&sysLoginReq)
		
			if isPass {	
				middleware.GenerateTokenForSys(c, sysAdmin)
			} else {
				result.Error(http.StatusOK, "验证失败," + err.Error())				 
			}
		}else {
			 
			result.Error(http.StatusBadRequest, err.Error())
		}
}

//



// @Summary 系统登录
// @Description 输入用户名及密登录
// @Produce json
// @Param tenant path string true "承租"
// @Param username path string true "登录名"
// @Param password path string true "登录密码"
// @Success 200 {object} swag.ResultContSingIn 成功后返回值
// @Router /singin [post]
func SingIn(c *gin.Context) {
	result := global.NewResult(c)

    username := c.PostForm("username")
	password := global.Encrypt( c.PostForm("password") )
	
    if username == "root" {
		result.Success(207,username+"="+password)
		loginReq := models.SysAdminstrator{}
		loginReq.UserName = username
		loginReq.Password  = password
		sysAdminTab := new(models.SysAdminstrator)

        res := global.DB().Where("user_name = ? AND password = ?", loginReq.UserName , loginReq.Password ).First(&sysAdminTab) 

		if res.Error != nil {
			result.Error(444,"Select error: "+ res.Error.Error())
			return
		}else{		 
			//result.Success(200,gin.H{"total":&res.RowsAffected, "data":&res.Value})
			result.Success(900, gin.H{"total":&res.RowsAffected, "data":&res.Value})
			return
		}
	}else{
		data := models.Login{}
		data.Password  = password
		result.Success(208,username)
	}
}

// @Summary 修改密码
// @Description 输入用户名及密登录
// @Produce json
// @Param email path string true "用户邮箱"
// @Param oldpassword path string true "旧密码"
// @Param newpassword path string true "新密码"
// @Param confirmpassword path string true "确认新密码"
// @Success 200 {object} swag.ResultContChangePass 成功后返回值
// @Router /changepass [put]
func ChangePass(c *gin.Context) {
	result := global.NewResult(c)
	/*
    data := map[string]interface{}{
		"username": c.PostForm("username") ,
		"uuid": global.CreateUUID(),
	    "password": global.Encrypt(c.PostForm("password")),
		"email":c.PostForm("email"),
		"createdAt":  time.Now().Unix(),
	}
    */
	fmt.Println("controller:ChangePass: "+time.Now().String())
    if c.PostForm("newpassword") != c.PostForm("confirmpassword") {
		errMsg  := "两次新密码不相同"
		code := 1000
		//fmt.Printf("Context: %s;",   errMsg  )	
		result.Error(code,errMsg)
		return
	}
	data := map[string]interface{}{
		"email": c.PostForm("email"),
	    "password": global.Encrypt(c.PostForm("newpassword")),
		"updateAt":  time.Now().Unix(),
	}
	fmt.Printf("Context: %s;",  data  )	
	result.Success(201,data);
	//result.Error(code int,msg string)	
	return
}
 