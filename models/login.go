package models


type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //验证码Id
	ImageUrl  string `json:"imageUrl"`  //验证码图片url
}

type Login struct {
	Id        		int     `gorm:"column:id;primaryKey" json:"id" xml:"id"`// 用户id
	TenantId  		string  `gorm:"column:tenant_id" json:"tenant_id" xml:" tenant_id" form:"tenant_id" binding:"required"` // tenant_id
    UserName   		string  `gorm:"column:user_name" json:"user_name" xml:"user_name" form:"user_name" binding:"required"` // 用户名
	Uuid      		string  `gorm:"column:uuid" json:"uuid" xml:"uuid"`  // 用户uuid
	MobilePhone 	string  `gorm:"column:mobile_phone" json:"mobile_phone"  xml:"mobile_phone" form:"mobile_phone" binding:"required"` // 移动电话
	Email     		string  `gorm:"column:email" json:"email" xml:"email" form:"email" validate:"required,email" binding:"required"` // 用户邮箱
	FirstName 		string  `gorm:"column:first_name" json:"first_name" xml:"first_name"` // 用户名
	LastName  		string  `gorm:"column:last_name" json:"last_name" xml:"last_name" ` // 用户名
	Gender			string  `gorm:"column:gender" json:"gender" xml:"gender" ` // 性别
	Password  		string  `gorm:"column:password" json:"password" xml:"password" form:"password" validate:"required,email" binding:"required"` // 用户密码
	CreatedAt 		int64   `gorm:"column:created_at" json:"created_at" xml:"created_at"`  // 创建时间戳
	UpdatedAt 		int64   `gorm:"column:updated_at" json:"updated_at" xml:"updated_at"`  // 创建时间戳
}  

func (Login) TableName() string {
	return "user"
}