package middleware
import (
	"net/http"
	_"reflect"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	
	"ginpro/models"
	"ginpro/global"
)

var Validate *validator.Validate
var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func init(){
	Validate = validator.New()

	fmt.Printf("Context: %s;",   "===init==="  )	
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
}

func GetBookable(c *gin.Context) {
	result := global.NewResult(c)
	

	var b models.Booking
	result.Success(2001,&b)
	result.Success(2002,binding.Query)
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		result.Success(200,"success-GetBookable-1111")
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		result.Success(200,"success-GetBookable-222")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	
}

func ValidateEmail(inputEmail string ) {
	//myEmail := "joeybloggs.gmail.com"
	errs := Validate.Var(inputEmail, "required,email")

	if errs != nil {
	  // 	fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return 
	}else{
		return 
	}

	// email ok, move on
}