package setting

import (
	"github.com/spf13/viper"
	"time"
)

type App struct {
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

// Setup initialize the configuration instance
func init() {
	viper.SetConfigFile("conf/app.yaml")
	//bytesContent, err := conf.Asset("conf/app.yaml")
	//if err != nil {
	//	panic("Asset() can not found setting file")
	//}
	//
	//viper.SetConfigType("yaml")
	//viper.ReadConfig(bytes.NewBuffer(bytesContent))
	viper.ReadInConfig()

	InitApp()
	InitServer()
}

func InitApp() {
	app := viper.Sub("app")

	AppSetting.RuntimeRootPath = app.GetString("RuntimeRootPath")
	AppSetting.LogFileExt = app.GetString("LogSavePath")
	AppSetting.LogSaveName = app.GetString("LogSaveName")
	AppSetting.LogFileExt = app.GetString("LogFileExt")
	AppSetting.TimeFormat = app.GetString("TimeFormat")
}

func InitServer() {
	server := viper.Sub("server")

	ServerSetting.RunMode = server.GetString("RunMode")
	ServerSetting.HTTPPort = server.GetInt("HTTPPort")
	ServerSetting.ReadTimeout = time.Duration(server.GetInt("ReadTimeout")) * time.Second
	ServerSetting.ReadTimeout = time.Duration(server.GetInt("ReadTimeout")) * time.Second
}

func InitDatabase() {
	database := viper.Sub("database")

	DatabaseSetting.Type = database.GetString("Type")
	DatabaseSetting.User = database.GetString("User")
	DatabaseSetting.Password = database.GetString("Password")
	DatabaseSetting.Host = database.GetString("Host")
	DatabaseSetting.Name = database.GetString("Name")
}

func Setup() {
	InitServer()
	InitApp()
	InitDatabase()
}
