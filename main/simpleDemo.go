package main

import (
	"flag"
	"golang.org/x/net/netutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/huangfly/SimpleRedisMemcache/cache"
	"github.com/huangfly/SimpleRedisMemcache/conf"
	"github.com/huangfly/SimpleRedisMemcache/encrypt"
	"github.com/huangfly/SimpleRedisMemcache/handle"
	"github.com/huangfly/SimpleRedisMemcache/mysql"
)

type Req struct {
	Key   string
	Value string
}

var (
	handler   *handle.Handle
	dbsql     mysql.SqlSvrInterface
	redispool cache.CacheInterface
	file      = flag.String("cfg", "", "config file")
)

func main() {
	//解析配置文件
	flag.Parse()
	config, err := conf.NewRedisConf(*file)
	if err != nil {
		log.Println(err)
		return
	}

	//解密mysql的password 并连接数据库
	sqlPass, _ := encrypt.Decode(config.Mysql.Password)
	dbsql = mysql.NewDataBase("mysql", config.Mysql.User+":"+sqlPass+"@tcp("+config.Mysql.Ip+")/test?charset=utf8")
	defer dbsql.Close()

	//解密rediss的password并连接redis
	redisPass, _ := encrypt.Decode(config.Redis.Password)
	redispool = cache.NewCachePool(config.Redis.Ip,
		redisPass,
		config.Redis.Maxactive,
		config.Redis.MaxIdle,
		config.Redis.IdleTimeout)
	defer redispool.Close()

	//创建数据库和缓存交互的handler
	handler = handle.NewHandle(dbsql, redispool)

	//创建httpserver
	conn, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Println("Listen: ", err)
		return
	}
	defer conn.Close()

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGHUP)
	go func() {
		<-c
		log.Println("Catch a Ctrl-c or Interrupt Signal Soft Will Exit")
		conn.Close()
	}()

	conn = netutil.LimitListener(conn, 10000) //限制并发上线
	http.HandleFunc("/GetRes", GetFuncHandler)
	http.HandleFunc("/StoreRes", StoreFuncHandler)
	http.HandleFunc("/DelRes", DelFuncHandler)
	printSoftInfo()
	http.Serve(conn, nil)
}
