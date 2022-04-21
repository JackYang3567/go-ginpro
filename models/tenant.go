package models 

import (
	_"time"
)

type Tenant struct {	
	 Id         int    `gorm:"column:id;primaryKey" json:"id" xml:"id" form:"id"` // 承租户id
	TenantId   string  `gorm:"column:tenant_id" json:"tenant_id" xml:"tenant_id" form:"tenant_id"` // 承租户uuid
    TenantName string  `gorm:"column:tenant_name;unique" json:"tenant_name" xml:"tenant_name" form:"tenant_name" binding:"required"` // 承租户名
	TenantDes  string  `gorm:"column:tenant_des" json:"tenant_des" xml:"tenant_des" form:"tenant_des"` // 承租户描述	
	DisabledIs int     `gorm:"column:disabled_is;" json:"disabled_is" xml:"disabled_is"` // 1:禁用当前记录
    Cctasp     string  `gorm:"column:cctasp" json:"cctasp" xml:"cctasp" form:"cctasp"` // 云计算技术和服务提供商
	Status     string  `gorm:"column:status" json:"status" xml:"status" form:"status"` // 状态
    StartAt    string  `gorm:"column:start_at" json:"start_at" xml:"start_at" form:"start_at"` // 承租起始日期
    DueAt      string  `gorm:"column:due_at" json:"due_at" xml:"due_at" form:"due_at"` // 承租到期日期
	CreatedAt  int64   `gorm:"column:created_at;autoCreateTime:milli" json:"created_at" xml:"created_at"` // 创建时间戳
	UpdatedAt  int64   `gorm:"column:updated_at;autoUpdateTime:milli" json:"updated_at" xml:"updated_at"`  // 创建时间戳
}

func (Tenant) TableName() string {
	return "tenant"
}