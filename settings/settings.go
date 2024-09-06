package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 使用viper来管理配置

//  使用内部定义的结构体来存储配置

// Conf 全局变量，来保存程序的所有配置信息;  适合团队开发
//var Conf = new(AppConfig)
//
//type AppConfig struct {
//	Name      string `mapstructure:"name"`
//	Mode      string `mapstructure:"mode"`
//	Version   string `mapstructure:"version"`
//	StartTime string `mapstructure:"start_time"`
//	MachineID int64  `mapstructure:"machine_id"`
//	Port      int    `mapstructure:"port"`
//
//	*LogConfig   `mapstructure:"log"`
//	*MySQLConfig `mapstructure:"mysql"`
//	*RedisConfig `mapstructure:"redis"`
//}
//
//type MySQLConfig struct {
//	Host         string `mapstructure:"host"`
//	User         string `mapstructure:"user"`
//	Password     string `mapstructure:"password"`
//	DB           string `mapstructure:"dbname"`
//	Port         int    `mapstructure:"port"`
//	MaxOpenConns int    `mapstructure:"max_open_conns"`
//	MaxIdleConns int    `mapstructure:"max_idle_conns"`
//}
//
//type RedisConfig struct {
//	Host         string `mapstructure:"host"`
//	Password     string `mapstructure:"password"`
//	Port         int    `mapstructure:"port"`
//	DB           int    `mapstructure:"db"`
//	PoolSize     int    `mapstructure:"pool_size"`
//	MinIdleConns int    `mapstructure:"min_idle_conns"`
//}
//
//type LogConfig struct {
//	Level      string `mapstructure:"level"`
//	Filename   string `mapstructure:"filename"`
//	MaxSize    int    `mapstructure:"max_size"`
//	MaxAge     int    `mapstructure:"max_age"`
//	MaxBackups int    `mapstructure:"max_backups"`
//}

// Init 这是一个人的项目
func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	if err = viper.ReadInConfig(); err != nil {
		// 读取配置信息失败
		fmt.Println("viper.ReadInConfig() failed,err", err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	//if err = viper.Unmarshal(&Conf); err != nil {
	//	fmt.Println("viper.Unmarshal() failed,err", err)
	//}

	// 热加载
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		// 热加载反序列化
		//if err := viper.Unmarshal(&Conf); err != nil {
		//	fmt.Println("viper.Unmarshal() failed,err", err)
		//}
	})
	return
}
