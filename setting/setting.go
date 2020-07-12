package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)
type Setting struct {
	vp *viper.Viper
}

type ServerSettingS struct {
	RunMode string
	HttpPort string
	ReadTimeOut time.Duration
	WriteTimeOut time.Duration
}

type AppSettingS struct {
	LogSavePath string
	LogFileName string
	LogFileExt string

}

type DatabaseSettingS struct {

	DBType string
	UserName string
	Password string
	Host string
	DBName string
	TablePrefix string
	Charset string
	ParseTime bool
	MaxIdleConns int
	MaxOpenConns int
}

//
func NewSetting()(*Setting,error)  {
	vp:=viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err:=vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp},nil

}
//
var sections =make(map[string]interface{})

//
func (s*Setting) ReadSection(k string,v interface{}) error {
	err:=s.vp.UnmarshalKey(k,v)
	if err != nil {
		return err
	}
	if _,ok:=sections[k];!ok {
		sections[k]=v
	}
	return nil
}
//
func (s*Setting)ReloadAllSection()error  {
	for k, v := range sections {
		err:=s.ReadSection(k,v)
		if err != nil {
			return err
		}
	}
	return nil
}
//
var (
	SeverSetting *ServerSettingS
	AppSetting *AppSettingS
	DatabaseSetting *DatabaseSettingS
)

func SetUpSetting()error  {

	setting,err:=NewSetting()
	if err != nil {
		return err
	}
	err=setting.ReadSection("server",&SeverSetting)
	if err != nil {
		return err
	}
	err=setting.ReadSection("app",&AppSetting)
	if err != nil {
		return err
	}
	err=setting.ReadSection("database",&DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func (s*Setting)WatchSettingChange()  {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_=s.ReloadAllSection()
		})
	}()
}


