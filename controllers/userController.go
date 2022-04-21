package controllers
 
import (
	"fmt"	
	"time"
	_"net/http"

	"github.com/gin-gonic/gin"
	
	"ginpro/global"
	_"ginpro/models"
)


type UserController struct{}

//@title User API
//@version 1.0
//@description XXXX项目描述
func init(){

}

// @Summary 用户详情
// @Description 显示用户的详情
// @Produce json
// @Param id path int true "用户id"
// @Success 200 {object} swag.ResultContUserById 成功后返回值
// @Router /user/{id} [get]
func GetUserById(c *gin.Context) {
	fmt.Println("controller:/user/{id}: "+time.Now().String())
	result := global.NewResult(c)
	result.Success(200,"success");
	return
}

// @Summary 用户详情
// @Description 显示用户的详情
// @Produce json
// @Param name path int true "用户name"
// @Success 200 {object} swag.ResultContUserByName 成功后返回值
// @Router /user/name/{name} [get]
func GetUserByName(c *gin.Context) {
	fmt.Println("controller:/user/{name}: "+time.Now().String())
	result := global.NewResult(c)
	result.Success(200,"success");
	return
}

// @Summary 用户详情
// @Description 显示用户的详情
// @Produce json
// @Param email path int true "用户email"
// @Success 200 {object} swag.ResultContUserByEmail 成功后返回值
// @Router /user/email/{email} [get]
func GetUserByEmail(c *gin.Context) {
	fmt.Println("controller:/user/{email}: "+time.Now().String())
	result := global.NewResult(c)
	result.Success(200,"success");
	return
}

// @Summary 用户列表
// @Description 显示用户列表
// @Produce json
// @Param tenant_id path int true "分类id"
// @Success 200 {object} swag.ResultContUserByTenantId 成功后返回值
// @Router /user/tenant/{tenant_id} [get]
func GetUserByTenantId(c *gin.Context) {
	fmt.Println("controller:index: "+time.Now().String())
	result := global.NewResult(c)
	result.Success(200,"success");
	return

}