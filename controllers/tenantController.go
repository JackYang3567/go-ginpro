package controllers

import (
	_"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ginpro/global"
	_"ginpro/middleware"
	"ginpro/models"
)


type TenantController struct{}


// @Summary 创建承租户
// @Description 输入承租户名及描述
// @Produce json
// @Param tenant_name path string true "承租户名"
// @Param tenant_des path string false "承租户描述"
// @Param cctasp path string false "云计算技术和服务提供商"
// @Param start_at path string false "承租起始日期"
// @Param due_at path string false "承租到期日期"
// @Success 200 {object} swag.ResultContCreatedTenant 成功后返回值
// @Router /tenant [post]
// @Security ApiKeyAuth
func CreatedTenant(c *gin.Context) {
	
	result := global.NewResult(c)
	data := models.Tenant{}
	err := c.ShouldBind(&data)

	tenant_name := c.PostForm("tenant_name")
	if  tenant_name == "" {
		//444客户端没做验证
		result.Error(http.StatusNoContent,"Validation error: "+ "tenant_name field required")
		return
	}
	 
 
	if err == nil {
	 
		data.TenantId   = global.CreateUUID()
		data.TenantName = tenant_name
		data.TenantDes  = c.PostForm("tenant_des")
		data.Cctasp     = c.PostForm("cctasp")
		data.StartAt    = c.PostForm("start_at")
		data.DueAt      = c.PostForm("due_at") 
		data.Status     = c.PostForm("status")
		data.CreatedAt  = time.Now().Unix() 
		data.UpdatedAt  = time.Now().Unix() 		
	 
		res := global.DB().Create(&data) 
		if res.Error != nil {
			//1000操作数据库错误
			result.Error(1000, "Insert error: "+ res.Error.Error())
		 }else{
			result.Success(201, gin.H{"total":&res.RowsAffected})//  返回插入记录的条数
		 }
		 return
	}
	
}


// @Summary 修改承租户
// @Description 输入承租户名及描述
// @Produce json
// @Param tenant_id path int true "承租户tenant_id"
// @Param tenant_name path string true "承租户名"
// @Param tenant_des path string false "承租户描述"
// @Param cctasp path string false "云计算技术和服务提供商"
// @Param start_at path string false "承租起始日期"
// @Param due_at path string false "承租到期日期"
// @Param disabled_is path string false "1/true禁用当前记录"
// @Success 200 {object} swag.ResultContUpdateTenant 成功后返回值
// @Router /tenant [put]
// @Security ApiKeyAuth
func UpdateTenant(c *gin.Context) {
	result := global.NewResult(c)
	tenant := new(models.Tenant)
	err := c.ShouldBind(&tenant)
	//tenant := new(models.Tenant)
	
	tenant_id   := c.PostForm("tenant_id")	
	tenant_name := c.PostForm("tenant_name")
	
	
	if  tenant_id == "" {
		//444客户端没做验证
		result.Error(444,"Validation error: "+ "tenant_id field required")
		return
	}
	if  tenant_name == "" {
		//444客户端没做验证
		result.Error(444,"Validation error: "+ "tenant_name field required")
		return
	}
	if err == nil {
	
			res := global.DB().Where("tenant_id = ?", tenant_id).First(tenant)
			if res.Error != nil {
				result.Error(444,"Select error: "+ res.Error.Error())
				return
			}else{
				/*
				res = global.DB().Model(&tenant).Updates(models.Tenant{
					TenantName: tenant_name,
					TenantDes:  c.PostForm("tenant_des"),			 
					UpdatedAt: time.Now().Unix(),
				})
				*/
				tenant.TenantName = tenant_name
				tenant.TenantDes  = c.PostForm("tenant_des")
				tenant.Cctasp     = c.PostForm("cctasp")
				tenant.Status     = c.PostForm("status")
				DisabledIs, err := strconv.Atoi(c.PostForm("disabled_is") )
				if err != nil {
					fmt.Println(err)
				}
				tenant.DisabledIs  =  DisabledIs
				
			//	StartAt, err := strconv.ParseInt(c.PostForm("start_at") , 10, 64)
			//	if err != nil {
			//		fmt.Println(err)
			//	}
				tenant.StartAt     = c.PostForm("start_at") //StartAt

				//DueAt, err := strconv.ParseInt(c.PostForm("due_at"), 10, 64 )
				//if err != nil {
				//	fmt.Println(err)
				//}
				tenant.DueAt       = c.PostForm("due_at")//DueAt
				
				tenant.UpdatedAt  = time.Now().Unix()
				res = global.DB().Save(&tenant)
				if res.Error != nil {
					result.Error(444,"Update error: "+ res.Error.Error())
					return
				}		 
				//result.Success(204,gin.H{"total":&res.RowsAffected, "data":&res.Value})
				result.Success(204,gin.H{"total":&res.RowsAffected})
				return
			}	
	} 
}

// @Summary 用tenant_id获取承租户详情
// @Description 显示承租户详情
// @Produce json
// @Param tenant_id path string true "承租id"
// @Success 200 {object} swag.ResultContGetTenantById 成功后返回值
// @Router /tenant/{tenant_id} [get]
// @Security ApiKeyAuth
func GetTenantById(c *gin.Context) {
	
	result := global.NewResult(c)
	//claims :=   c.MustGet("claims").(*middleware.CustomClaims)
	
	//	result.Success(http.StatusOK, claims)
	
	
        // 继续交由下一个路由处理,并将解析出的信息传递下去
        

	tenant_id := c.Param("tenant_id")
	
	tenant := new(models.Tenant)

    res := global.DB().Where("tenant_id = ?", tenant_id).First(&tenant) 
	//res := global.DB().First(&tenant,"tenant_id = ?", tenant_id) 
    if res.Error != nil {
		result.Error(444,"Select error: "+ res.Error.Error())
        return
    }else{		 
		//result.Success(200,gin.H{"total":&res.RowsAffected, "data":&res.Value})
		result.Success(200, gin.H{"total":&res.RowsAffected, "data":&res.Value})
		return
	}	
}

// @Summary 用tenant_name获取承租户详情
// @Description 显示承租户详情
// @Produce json
// @Param tenant_name path string true "承租tenant_name"
// @Success 200 {object} swag.ResultContGetTenantByName成功后返回值
// @Router /tenant/name/{tenant_name} [get]
// @Security ApiKeyAuth
func GetTenantByName(c *gin.Context) {
	result := global.NewResult(c)
	tenant_name := c.Param("tenant_name")
//	result.Success(200,tenant_name)
	
	tenant := new(models.Tenant)
     
    res := global.DB().Where("tenant_name = ?", tenant_name).First(tenant) 
    if res.Error != nil {
		result.Error(4494,"Select error: "+ res.Error.Error())
        return
    }else{		 
		result.Success(200, gin.H{"total":&res.RowsAffected, "data":&res.Value})
		return
	}	
}

// @Summary 获取所有承租户详情
// @Description 显示承租户详情
// @Produce json  
// @Param pageNo path string true "承租pageNo"
// @Param pageSize path string true "承租pageSize"
// @Success 200 {object} swag.ResultContGetTenantList成功后返回值
// @Router /tenant [get]
// @Security ApiKeyAuth
func GetTenantList(c *gin.Context) {
	result := global.NewResult(c)	
	tenants := new([] models.Tenant) 
    
	var pageNo int  = 1
	var pageSize int = 10

	pageNo, err := strconv.Atoi(c.Query("pageNo") )
	if err != nil {
		fmt.Println(err)
	}

	pageSize, err1 :=  strconv.Atoi(c.Query("pageSize") )
	if err1 != nil {
	   fmt.Println(err1)
   }
    var total int
	global.DB().Model(&models.Tenant{}).Count(&total)
	res := global.DB().Offset((pageNo-1)*pageSize).Limit(pageSize).Order("created_at desc, tenant_name").Find(&tenants) 
    if res.Error != nil {
		result.Error(444,"Select error: "+ res.Error.Error())
        return
    }else{
		result.Success(200, gin.H{"pageSize":&res.RowsAffected, "data":&res.Value,"total":total})
		return
	}	
}
 
// @Summary 删除承租户
// @Description 删除承租户
// @Produce json
// @Param tenant_id path string true "承租id"
// @Success 200 {object} swag.ResultContDeleteTenantById 成功后返回值
// @Router /tenant/{tenant_id} [delete]
// @Security ApiKeyAuth
func DeleteTenantById(c *gin.Context) {
	result := global.NewResult(c)
	tenant_id := c.Param("tenant_id")
	
	tenant := new(models.Tenant) 
     
	res := global.DB().Where("tenant_id = ?", tenant_id).Delete(tenant) 
	if res.Error != nil {
		result.Error(1000,"Delete error: "+ res.Error.Error())
		return
	}
	result.Success(http.StatusNoContent,gin.H{"total":&res.RowsAffected})
	 
}
