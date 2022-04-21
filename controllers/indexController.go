//@Summary Test
//@title XXXX项目名称-Swagger Example API
package controllers
 
import (
	"fmt"		
	"time"
    "net/http"

	"github.com/gin-gonic/gin"	

	"github.com/gin-gonic/gin/binding"

	"ginpro/global"
	"ginpro/models"
)

// Controller
type IndexController struct{}
func NewIndexController() IndexController {
	return IndexController{}
}
// @Summary Test
// @Description Test===
// @Produce json
// @Param check_in path string true "check_in"
// @Param check_out path string true "check_out"
// @Success 200 {object} swag.ResultContTest 成功后返回值
// @Router /test [get]
func Test(c *gin.Context) {
	fmt.Println("controller:Test: "+time.Now().String())
	result := global.NewResult(c)
	//
   
	

	var b models.Booking
	
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
	
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	
	//
	
	
	result.Success(200,time.Now())
}


// @Summary 商品详情
// @Description 显示商品的详情
// @Produce json
// @Param goodsid path int true "商品id"
// @Success 200 {object} swag.ResultContGoods 成功后返回值
// @Router /index/goodsone/{goodsid} [get]
func (g *IndexController) GoodsOne(c *gin.Context) {
	fmt.Println("controller:index: "+time.Now().String())
	result := global.NewResult(c)
	result.Success(200,"success");
	return
}
// @Summary 商品列表
// @Description 显示商品列表
// @Produce json
// @Param categoryId path int true "分类id"
// @Success 200 {object} swag.ResultContGoodsList 成功后返回值
// @Router /index/goodslist [get]
func (g *IndexController) GoodsList(c *gin.Context) {
	fmt.Println("controller:index: "+time.Now().String())
	result := global.NewResult(c)
	result.Success(200,"success");
	return

}