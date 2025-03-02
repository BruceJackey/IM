package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var Red *redis.Client

// InitConfig 初始化配置文件
func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited................")
}

// InitMySQL 初始化mysql连接
func InitMySQL() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)
	//自定义日志，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //NySQL阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //是否彩色打印
		},
	)

	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	fmt.Println("mysql inited................")
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println(user)
}

// InitRedis 初始化redis连接
func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.database"),
		PoolSize:     viper.GetInt("redis.pool_size"),
		MinIdleConns: viper.GetInt("redis.min_idle_conns"),
	})
	fmt.Println("redis inited.")
}

const (
	PublishKey = "websoket"
)

// Publish 发布消息
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("publish msg:", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("subscribe msg:", msg)
	return msg.Payload, err
}
