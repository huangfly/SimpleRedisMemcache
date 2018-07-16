package conf

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type RedisConf struct {
	Mysql MysqlInfo
	Redis RedisInfo
}

type MysqlInfo struct {
	Ip       string
	User     string
	Password string
}

type RedisInfo struct {
	Ip          string
	Password    string
	Maxactive   int
	MaxIdle     int
	IdleTimeout int
}

func NewRedisConf(path string) (*RedisConf, error) {
	redisconf := new(RedisConf)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Read config file <%s> failure. err:%+v", "test", err)
		return nil, err
	}
	err = xml.Unmarshal(content, redisconf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return redisconf, nil
}
