package controllers

import (
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strings"
	"time"

	"ginpro/global"
	"ginpro/models"
	_"ginpro/middleware"
)

 

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	fmt.Println("file : " + file)
	fmt.Println("ext : " + ext)
	fmt.Println("id : " + id)
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("reload : " + r.FormValue("reload"))
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}


// @Summary 获取验证码
// @Description 获取验证码 GetCaptcha
// @Produce json
// @Success 200 {object} swag.ResultContCaptcha 成功后返回值
// @Router /captcha [get]
func GetCaptcha(c *gin.Context) {

	
	    result := global.NewResult(c)
		length := 4 //captcha.DefaultLen
		captchaId := captcha.NewLen(length)
		var captcha models.CaptchaResponse
		captcha.CaptchaId = captchaId
		captcha.ImageUrl =  captchaId + ".png"
		fmt.Println("GetCaptchaPng : " + captchaId)
		

       // header := c.Request
		//fmt.Println("Header==== : " + header)
	//	c.JSON(http.StatusOK, header)
	

		//c.JSON(http.StatusOK, captcha)
		//c.Redirect(http.StatusFound, captcha.ImageUrl)
		
		result.Success(200,gin.H{"Authorization":c.Request.Header.Get("Authorization"),
								 "MB-Property-ID":c.Request.Header.Get("MB-Property-ID"),
							
								 "captcha":captcha});
}

// @Summary 获取验证码图片
// @Description 获取验证码 GetCaptcha
// @Produce json
// @Param captchaId path int true "验证码Id"
// @Success 200 {object} swag.ResultContGetCaptchaImage 成功后返回值
// @Router /captcha/{captchaId} [get]
func GetCaptchaImage(c *gin.Context) {
	//result := global.NewResult(c)
	//captchaId := c.Param("captchaId")
	//c.JSON(http.StatusOK, captchaId)
	//result.Success(200,captchaId);
    ServeHTTP(c.Writer, c.Request)
}


// @Summary 验证 验证码码图
// @Description 获取验证码 GetCaptcha
// @Produce json
// @Param captchaId path int true "验证码Id"
// @Param value path int true "验证码value"
// @Success 200 {object} swag.ResultContVerifyCaptcha 成功后返回值
// @Router /captcha/verify/{captchaId}/{value} [get]
func VerifyCaptcha(c *gin.Context) {
	result := global.NewResult(c)

	    captchaId := c.Param("captchaId")
		value := c.Param("value")
		if captchaId == "" || value == "" {
			//c.String(http.StatusBadRequest, "参数错误")
			result.Error(http.StatusBadRequest,"参数错误");
		}
		if captcha.VerifyString(captchaId, value) {
			//c.JSON(http.StatusOK, "验证成功")
			result.Success(http.StatusOK,gin.H{
				 "data":true});
		} else {
			//c.JSON(http.StatusOK, "验证失败")
			result.Error(http.StatusOK,"验证失败");
		}

	
}