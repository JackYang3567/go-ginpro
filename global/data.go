package global

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	_"database/sql"	
	"fmt"	 
	"log"
	"time"
	"os"
	 
	
	_ "github.com/go-sql-driver/mysql"
   "github.com/jinzhu/gorm"

   "github.com/go-redis/redis/v8"


	"ginpro/conf"
 
)
// redis
var Ctx = context.Background()
var RDB = redis.NewClient(&redis.Options{
	Addr:  	  conf.RedisHttpaddr,
	Password: conf.RedisPass, // no password set
	DB:       conf.RedisDb,  // use default DB
})

var logger *log.Logger
// MySql
var _DB *gorm.DB

func DB() *gorm.DB {
   return _DB
}

func init() {
   _DB = initDB()

   file, err := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
  
}
func initDB() *gorm.DB {
	var mysqldb = conf.MySqlDb
	var mysqluser = conf.MySqlUser
	var mysqlpass = conf.MySqlPass
	var httpaddr = conf.HttpAddr
   
   db, err := gorm.Open("mysql", ""+mysqluser+":"+mysqlpass+"@tcp("+httpaddr+":3306)/"+mysqldb+"?charset=utf8&parseTime=True&loc=Local")
   if err != nil {
      panic(err)
   }
   db.DB().SetMaxOpenConns(100)
   db.DB().SetMaxIdleConns(10)
   db.DB().SetConnMaxLifetime(time.Second * 300)
   if err = db.DB().Ping(); err != nil {
      panic(err)
   }
   return db


}

func GetExpiresIn(keyName string,value string) (ExpiresIn int64) {
	newKey := "tei_"+keyName
	fmt.Println("=value====",value)
	result, err := RDB.Exists(Ctx, newKey).Result()
	if err != nil {
		fmt.Println("=err=result====",err.Error())
		Danger(err, "Cannot parse form")
		panic(err.Error() + "==333333")
	}else{
		fmt.Println("11111=result====",newKey,"=====",result)
		if result == 0 {
			err = RDB.Set(Ctx, newKey, value, 0).Err()
			 if err != nil {
				fmt.Println("11111=err====",newKey,"=====",err)	 
			 }else{
				fmt.Println("00000=err====",newKey,"=====",err)	 
			   errExpire := RDB.Expire(Ctx, newKey, 3600).Err()
			   if errExpire != nil {
				fmt.Println("00000=errExpire====",newKey,"=====",errExpire)	 
			   }else{
				fmt.Println("cccc=errExpire====",newKey,"=====",errExpire)	  
			   }
			  return
			 }	
			
		} 
		if result == 1 {
			fmt.Println("=result====",result)
				val ,errTtl := RDB.TTL(Ctx, newKey).Result()
				fmt.Println("=val====",val)
				if errTtl != nil {
					fmt.Println("=errTtl====",errTtl.Error())
				}else{
				   if val == -2 {
					  errExpire := RDB.Expire(Ctx, newKey, 3600*1000).Err()
					  if errExpire != nil {
						panic(errExpire.Error() + "==vvvvvvv")
					 }
				   }
				   if val == -1 {
						 errExpire := RDB.Expire(Ctx, newKey, 3600*1000).Err()
							if errExpire != nil {
								panic(errExpire.Error() + "==wwwwwww")
						}
				   }
				   if val >=0 {
                     return  
				   }
				}
				
		}
	}
	return
}
/*
func init() {
	
	var mysqldb = conf.MySqlDb
	var mysqluser = conf.MySqlUser
	var mysqlpass = conf.MySqlPass
	var httpaddr = conf.HttpAddr
	
	//连接数据库
    constr := ""+mysqluser+":"+mysqlpass+"@tcp("+httpaddr+":3306)/"+mysqldb+""
    //打开连接
    Db ,err := sql.Open("mysql",constr) //返回mysql实例db
    if err != nil {
       // log.Panic(err.Error())
		log.Fatalln("Cannot Open ", err)
        return
    }

	//插入数据
    _,err = Db.Exec("insert into user(name,password,uuid,email,created_at) values(?,?,?,?,?);","张三88","123456","3453252","ewrw@163.com",time.Now().Unix())
    if err != nil {
        //log.Panic(err.Error())
		log.Fatalln("Cannot insert", err)
        return
    }else{
        fmt.Println("插入成功")
	}
}

*/
// create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

// for logging
func Info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func Danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func Warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

