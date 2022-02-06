package storage

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config ... 接続情報
type Config struct {
	Address  string
	User     string
	Password string
	DB       string
	IsTrace  bool
}

func (conf *Config) DSN() string {
	return fmt.Sprintf("%s:%s@%s/%s?parseTime=true&loc=Asia%%2FTokyo&charset=utf8mb4",
		conf.User, conf.Password, conf.Address, conf.DB)
}

func ConfigDB(db string) *Config {
	hKey := fmt.Sprintf("MYSQL.%s.HOST", db)
	h := viper.GetString(hKey)
	if h == "" {
		panic(fmt.Errorf("no config key %s", hKey))
	}

	portKey := fmt.Sprintf("MYSQL.%s.PORT", db)
	port := viper.GetString(portKey)
	if port == "" {
		port = "3306"
	}

	uKey := fmt.Sprintf("MYSQL.%s.USER", db)
	u := viper.GetString(uKey)
	if u == "" {
		panic(fmt.Errorf("no config key %s", uKey))
	}

	pKey := fmt.Sprintf("MYSQL.%s.PASSWORD", db)
	p := viper.GetString(pKey)
	if p == "" {
		panic(fmt.Errorf("no config key %s", pKey))
	}

	dKey := fmt.Sprintf("MYSQL.%s.DATABASE", db)
	d := viper.GetString(dKey)
	if d == "" {
		panic(fmt.Errorf("no config key %s", dKey))
	}

	t := viper.GetBool(fmt.Sprintf("MYSQL.%s.TRACE", db))

	return &Config{
		Address:  fmt.Sprintf("tcp(%s:%s)", h, port),
		User:     u,
		Password: p,
		DB:       d,
		IsTrace:  t,
	}
	//return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", u, p, h, port, d)
}
