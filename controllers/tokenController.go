package controllers
 
import (
	"fmt"	
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_"github.com/golang-jwt/jwt"
	
	"ginpro/global"
	"ginpro/models"
    _"ginpro/middleware"
)
 
type TokenController struct{}



func NewTokenController() TokenController {
	return TokenController{}
}

// @Summary 生成Token
// @Description 获取Token
// @Produce json
// @Param tenant path string true "承租"
// @Param name path string true "登录名"
// @Param password path string true "登录密码"
// @Success 200 {object} swag.ResultContToken 成功后返回值
// @Router /token [post]
func CreatedToken(c *gin.Context) {
	fmt.Println("controller:GetToken: "+time.Now().String())
    	 
	

	var t models.Token
	c.Bind(&t)
	c.JSON(http.StatusOK, gin.H{
		"access_token": global.CreateUUID(),
		"expires_in": 3600,
	    "refresh_token":global.CreateUUID(),
		"scope":"app",
		"token_type":"bearer",
	})
	 
    
	result := global.NewResult(c)
	result.Success(200,"success");
	return
}


// @Summary 取Token
// @Description 获取Token
// @Produce json
// @Param tenant path string true "承租"
// @Param user path string true "登录名"
// @Param password path string true "登录密码"
// @Success 200 {object} swag.ResultContToken 成功后返回值
// @Router /token [get]
func (g *TokenController) GetToken(c *gin.Context) {
	fmt.Println("controller:GetToken: "+time.Now().String())
	result := global.NewResult(c)
	result.Success(200,"success");
	return
} 

/*
// 生成令牌
func GenerateToken(c *gin.Context, sysAdmin models.SysAdminstrator) {
	result := global.NewResult(c)
	
    j := &middleware.JWT{
        []byte("newtrekWang"),
    }
    claims := middleware.CustomClaims{
        sysAdmin.Id,
        sysAdmin.UserName,
        sysAdmin.MobilePhone,
        jwt.StandardClaims{
            NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
            ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
            Issuer:    "newtrekWang",                   //签名的发行者
        },
    }

    token, err := j.CreateToken(claims)

    if err != nil {
		result.Error(6666, "1111111")
        c.JSON(http.StatusOK, gin.H{
            "status": -1,
            "msg":    err.Error(),
        })
        return
    }

   // log.Println(token)

    data := models.TokenResult{   
		AccessToken: token,
		ExpiresIn: int64(time.Now().Unix() + 3600),
		RefreshToken: "token",
		Scope: "app",
		TokenType: "bearer",
	 
    }

	result.Success(http.StatusOK,data)
}
*/