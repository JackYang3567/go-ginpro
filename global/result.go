package global
 
import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	
)
 
type Result struct {
	Ctx *gin.Context
}
 
type ResultCont struct {
	StatusCode int  `json:"statusCode" example:"0"` //代码,200:成功，非0:失败
	Msg string  `json:"msg" example:"错误"` //报错信息
	Data interface{} `json:"result" `  //返回的数据
}
 
func NewResult(ctx *gin.Context) *Result {
	
	return &Result{Ctx: ctx}
}
 
//返回成功
func (r *Result) Success(code int,data interface{}) {
	if (data == nil) {
		data = gin.H{}
	}
	res := ResultCont{}
	res.StatusCode = code
	res.Msg = ""
	res.Data = data
	r.Ctx.JSON(http.StatusOK,res)
}
 
//返回失败
func (r *Result)Error(code int,msg string) {
	res := ResultCont{}
	res.StatusCode = code
	res.Msg = msg
	res.Data = gin.H{}
	r.Ctx.JSON(http.StatusOK,res)
	r.Ctx.Abort()

}