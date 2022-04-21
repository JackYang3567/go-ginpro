package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"
    _"fmt"
    "strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

    "ginpro/global"
	"ginpro/models"
)

// 一些常量
var (
    TokenExpired     error  = errors.New("Token is expired")
    TokenNotValidYet error  = errors.New("Token not active yet")
    TokenMalformed   error  = errors.New("That's not even a token")
    TokenInvalid     error  = errors.New("Couldn't handle this token:")
    SignKey          string = "shiji"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
    Id                  int `json:"Id"`
    UserName            string `json:"user_name"`
    Email     		    string `json:"email"`
    MobilePhone         string `json:"mobile_phone"`
    Role                string `json:"role"`
    jwt.StandardClaims
}

// JWT 签名结构
type JWT struct {
    SigningKey []byte
}



// 新建一个jwt实例
func NewJWT() *JWT {
    return &JWT{
        []byte(GetSignKey()),
    }
}
// 获取signKey
func GetSignKey() string {
    return SignKey
}
// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        result := global.NewResult(c)
        authstr :=  c.Request.Header.Get("Authorization")       
        space := strings.Index(authstr, " ") +1

         token := authstr[space:]//c.Request.Header.Get("token")
        
        if token == "" {
            result.Error(http.StatusForbidden, "请求未携带token，无权限访问")
            
            c.Abort()
            return
        }

      

        j := NewJWT()
        // parseToken 解析token包含的信息
        claims, err := j.ParseToken(token)
        //claims.ExpiresAt = int64(time.Now().Unix() - 7200)
        if err != nil {
          
            switch err {
                case TokenExpired: result.Error(http.StatusForbidden,err.Error()+" 授权已过期") 

                case TokenInvalid: result.Error(http.StatusForbidden,err.Error())

                default: result.Error(http.StatusForbidden,err.Error())
            }
             
          //  result.Error(http.StatusForbidden,err.Error())
            c.Abort()
            return
        }
        //log.Print("get claims: ", claims)
        //result.Success(http.StatusOK, claims)
        // 继续交由下一个路由处理,并将解析出的信息传递下去
        c.Set("claims", claims)
    }
}


// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return j.SigningKey, nil
    })
    if err != nil {
        if ve, ok := err.(*jwt.ValidationError); ok {
            if ve.Errors&jwt.ValidationErrorMalformed != 0 {
                return nil, TokenMalformed
            } else if ve.Errors&jwt.ValidationErrorExpired != 0 {
                // Token is expired
                return nil, TokenExpired
            } else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
                return nil, TokenNotValidYet
            } else {
                return nil, TokenInvalid
            }
        }
    }
    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
    jwt.TimeFunc = func() time.Time {
        return time.Unix(0, 0)
    }
    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return j.SigningKey, nil
    })
    if err != nil {
        return "", err
    }
    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        jwt.TimeFunc = time.Now
        claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
        return j.CreateToken(*claims)
    }
    return "", TokenInvalid
}

// 生成令牌  sysAdmin.Genre="sys$admin",
func GenerateTokenForSys(c *gin.Context, sysAdmin models.SysAdminstrator) {
    result := global.NewResult(c)
    j :=JWT{  
          []byte(SignKey),
    }
    expat := int64(time.Now().Unix() + 3600)
    claims := CustomClaims{
        sysAdmin.Id,
        sysAdmin.UserName,
        sysAdmin.Email,
        sysAdmin.MobilePhone,
        sysAdmin.Role,
        jwt.StandardClaims{
            NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
            ExpiresAt: expat,   // 过期时间 一小时
            Issuer:    SignKey, //签名的发行者
        },
    }
    
    token, err := j.CreateToken(claims)
   
    if err != nil {
        result.Error(http.StatusOK, err.Error())
        /*
        c.JSON(http.StatusOK, gin.H{
            "status": -1,
            "msg":    err.Error(),
        })
        */
        return
    }

   // log.Println(token)

    data := models.TokenResult{
        AccessToken: token,
		ExpiresIn: global.GetExpiresIn(sysAdmin.Email,token)  ,//int64(expat - time.Now().Unix() ),
		RefreshToken: "token",
		Scope: "app",
		TokenType: "bearer",
       // SysAdminstrator:sysAdmin,
    }
    result.Success(http.StatusOK, data)
    /*c.JSON(http.StatusOK, gin.H{
        "status": 0,
        "msg":    "登录成功！===",
        "data":   data,
    })
    */
    return
}

// 生成令牌
func GenerateToken(c *gin.Context, sysAdmin models.SysAdminstrator) {
    j :=JWT{  
          []byte(SignKey),
    }
    expat := int64(time.Now().Unix() + 3600)
    claims := CustomClaims{
        sysAdmin.Id,
        sysAdmin.UserName,
        sysAdmin.Email,
        sysAdmin.MobilePhone,
        sysAdmin.Role,
        jwt.StandardClaims{
            NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
            ExpiresAt: expat,   // 过期时间 一小时
            Issuer:    SignKey, //签名的发行者
        },
    }

    token, err := j.CreateToken(claims)

    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "status": -1,
            "msg":    err.Error(),
        })
        return
    }

    log.Println(token)
   
    data := models.TokenResult{
        AccessToken: token,
		ExpiresIn: int64(expat - time.Now().Unix() ),
		RefreshToken: "token",
		Scope: "app",
		TokenType: "bearer",
       // SysAdminstrator:sysAdmin,
    }
    c.JSON(http.StatusOK, gin.H{
        "status": 0,
        "msg":    "登录成功！",
        "data":   data,
    })
    return
}