package setting

import (
	"github.com/go-ini/ini"
	"github.com/gpmgo/gopm/modules/log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("Failed to parse 'conf/app.ini' :%v", err)
	}

	LoadBase()

	LoadServer()

	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get section 'server':v%", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal("Failed to  get section 'app' :v%", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)SD##@S")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
