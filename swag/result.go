package swag
 
import "ginpro/models"



//测试用 路由、控制器、中间件
type ResultContTest struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.User `json:"data" `  //返回的数据
}

// 获取验证码
type ResultContCaptcha struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.CaptchaResponse `json:"data" `  //返回的数据
}
// 获取验证码图片
type ResultContGetCaptchaImage struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.CaptchaResponse `json:"data" `  //返回的数据
}
//验证验证码
type ResultContVerifyCaptcha struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.CaptchaResponse `json:"data" `  //返回的数据
}

//创建承租户
type ResultContCreatedTenant struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.Tenant `json:"data" `  //返回的数据
}

//修改承租户
type ResultContUpdateTenant struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.Tenant `json:"data" `  //返回的数据
}

//用tenant_id获取承租户详情
type ResultContGetTenantById struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.Tenant `json:"data" `  //返回的数据
}

// 用tenant_name获取承租户详情
type ResultContGetTenantByName struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.Tenant `json:"data" `  //返回的数据
}
//获取承租户列表
type ResultContGetTenantList struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.Tenant `json:"data" `  //返回的数据
}


//删除承租户
type ResultContDeleteTenantById struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.Tenant `json:"data" `  //返回的数据
}

//注册信息
type ResultContSingUp struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.User `json:"data" `  //返回的数据
}

// 登录信息
type ResultContSingIn struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.User `json:"data" `  //返回的数据
}

// 修改密码信息
type ResultContChangePass struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.User `json:"data" `  //返回的数据
}

// Token信息
type ResultContToken struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.Token `json:"data" `  //返回的数据
}

// User信息
type ResultContUserById struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.User `json:"data" `  //返回的数据
}
// User by Name 信息
type ResultContUserByName struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.User `json:"data" `  //返回的数据
}

// User by Email 信息
type ResultContUserByEmail struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.User `json:"data" `  //返回的数据
}

// User by Tenant 信息
type ResultContUserByTenantId struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.User `json:"data" `  //返回的数据
}



//商品详情
type ResultContGoods struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data *models.Goods `json:"data" `  //返回的数据
}
 
//商品列表
type ResultContGoodsList struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,0:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data []*models.Goods `json:"data" `  //商品列表

}