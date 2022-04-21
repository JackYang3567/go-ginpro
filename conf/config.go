package conf

import (
	     "log"
	     "time"	 
         
	     "github.com/go-ini/ini"
	 )
	 
    var (
        Cfg *ini.File
    
        RunMode string
    
        HTTPPort     string
        ReadTimeout  time.Duration
        WriteTimeout time.Duration

		HttpAddr     string
		MySqlPort    int 
		MySqlUser    string
		MySqlPass    string
		MySqlDb      string

        RedisHttpaddr string
        RedisPass string
        RedisDb int

        PageSize  int
        JwtSecret string
    )
    
    func init() {
        var err error
        Cfg, err = ini.Load("conf/app.ini")
        if err != nil {
            log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
        }
    
        LoadBase()
        LoadServer()
        LoadApp()
		LoadDataBaseServer()
    }
    
    func LoadBase() {
        RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
    }
	    
    func LoadServer() {
        sec, err := Cfg.GetSection("server")
        if err != nil {
            log.Fatalf("Fail to get section 'server': %v", err)
        }
    
		RunMode = sec.Key("runmode").MustString("debug") 
        HTTPPort =sec.Key("httpport").MustString("8000")
        ReadTimeout = time.Duration(sec.Key("ReadTimeout").MustInt(60)) * time.Second
        WriteTimeout = time.Duration(sec.Key("WriteTimeout").MustInt(60)) * time.Second
    }
    
    func LoadApp() {
        sec, err := Cfg.GetSection("app")
        if err != nil {
            log.Fatalf("Fail to get section 'app': %v", err)
        }
    
        JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
        PageSize = sec.Key("PAGE_SIZE").MustInt(10)
    }

	func LoadDataBaseServer() {
        sec, err := Cfg.GetSection("dbserver")
        if err != nil {
            log.Fatalf("Fail to get section 'server': %v", err)
        }
    
		HttpAddr     = sec.Key("httpaddr").MustString("0.0.0.0") 
		MySqlPort    = sec.Key("mysqlport").MustInt(3306)  
		MySqlUser    = sec.Key("mysqluser").MustString("root") 
		MySqlPass    = sec.Key("mysqlpass").MustString("") 
		MySqlDb      = sec.Key("mysqldb").MustString("repository") 
		 
    }

    func LoadRedisBaseServer() {
        sec, err := Cfg.GetSection("redisserver")
        if err != nil {
            log.Fatalf("Fail to get section 'server': %v", err)
        }
    
		RedisHttpaddr     = sec.Key("redishttpaddr").MustString("localhost:6379") 
		RedisPass    = sec.Key("redispass").MustString("")  
		RedisDb    = sec.Key("redisdb").MustInt(0) 
		 
		 
    }

	 
		    
		     
		     
		  