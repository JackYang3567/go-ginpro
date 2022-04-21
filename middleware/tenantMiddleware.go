package middleware

import (

	
	"ginpro/models"
	"ginpro/global"

)


func CreateTenant(tenant *models.Tenant) (err error) {
    err = global.DB().Create(tenant).Error
    return
}