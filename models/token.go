package models

type Token struct{	
	Id       		 int     `gorm:"column:id;primaryKey" json:"id" xml:"id"` // token id
    AccessToken      string  `gorm:"column:access_token" json:"access_token" xml:"access_token" form:"access_token"` // token uuid
	ExpiresIn     	 int64   `gorm:"column:expires_in" json:"expires_in" xml:"expires_in" form:"expires_in"` // 存活时间3600秒 1时间
	RefreshToken     string  `gorm:"column:refresh_token" json:"refresh_token" xml:"refresh_token" form:"refresh_token"` // 刷新token uuid
	Scope  			 string  `gorm:"column:scope" json:"scope" xml:"scope" form:"scope"` // 范围
	TokenType 		 string  `gorm:"column:token_type" json:"token_type" xml:"token_type" form:"token_type"` // token类型
}

type TokenResult struct {
	AccessToken      string  `json:"access_token"` // token uuid
	ExpiresIn     	 int64   `json:"expires_in"` // 存活时间3600秒 1时间
	RefreshToken     string  `json:"refresh_token"` // 刷新token uuid
	Scope  			 string  `json:"scope"` // 范围
	TokenType 		 string  `json:"token_type"` // token类型
   // SysAdminstrator
}

func (Token) TableName() string {
	return "token"
}