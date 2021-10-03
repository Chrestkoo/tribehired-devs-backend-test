package setting

import (
	"log"
	"strings"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	Domain       string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Email struct {
	MailHost     string
	MailPort     string
	MailUserName string
	MailName     string
	MailPassword string
}

var EmailSetting = &Email{}

type Cors struct {
	CorsDomain  string
	AllowOrigin []string
}

var CorsSetting = &Cors{}

type Explorer struct {
	ExplorerUrl string
}

var ExplorerSetting = &Explorer{}

var Cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("email", EmailSetting)
	mapTo("cors", CorsSetting)
	mapTo("explorer", ExplorerSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	CorsSetting.AllowOrigin = strings.Split(CorsSetting.CorsDomain, ",")
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := Cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
