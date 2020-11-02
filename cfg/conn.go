package cfg

import (
	"fmt"
	"github.com/spf13/viper"
)

type Conn struct {
	host string
	port int
	db   string
	user string
	pwd  string
}

func (conn *Conn) ToString(driver string) string {
	return fmt.Sprintf("%s://%s:%s@%s:%d", driver, conn.user, conn.pwd, conn.host, conn.port)
}

func NewConn() (conn *Conn, err error) {
	viper.SetConfigName("conn")
	viper.AddConfigPath("/data/pagestore/config")
	viper.SetConfigType("yaml")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("read config error: %v", err)
		return
	}

	viper.SetConfigName("conn-local")
	viper.AddConfigPath("/data/pagestore/config")

	if err = viper.MergeInConfig(); err != nil {
		fmt.Printf("merge config error: %v", err)
		return
	}

	conn = &Conn{
		host: viper.GetString("mongo.host"),
		port: viper.GetInt("mongo.port"),
		db:   viper.GetString("mongo.db"),
		user: viper.GetString("mongo.user"),
		pwd:  viper.GetString("mongo.pwd"),
	}

	return
}
