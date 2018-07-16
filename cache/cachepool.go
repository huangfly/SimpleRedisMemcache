package cache

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

//redispool
type RedisCachePool struct {
	pool *redis.Pool
}

//从reidspool中获取一个连接
func (this *RedisCachePool) get() redis.Conn {
	return this.pool.Get()
}

//从redis数据中找出key对应的value
func (this *RedisCachePool) Get(key string) (string, error) {
	conn := this.get()
	if err := conn.Err(); err != nil {
		log.Println("Redis Get is failed : ", err)
		return "", err
	}
	defer conn.Close()
	contant, err := redis.String(conn.Do("Get", key))
	return contant, err
}

//设置key对应的value的值，并且设置该值的生命周期
func (this *RedisCachePool) Set(key string, val string, second int64) error {
	conn := this.get()
	if err := conn.Err(); err != nil {
		log.Println("Redis Get is failed : ", err)
		return err
	}
	defer conn.Close()
	conn.Do("MULTI")
	conn.Do("SET", key, val)
	if second > 0 {
		conn.Do("EXPIRE", key, second)
	}
	conn.Do("EXEC")
	return nil
}

//删除key-value键值对
func (this *RedisCachePool) Delete(key string) error {
	conn := this.get()
	if err := conn.Err(); err != nil {
		log.Println("Redis Delete is failed : ", err)
		return err
	}
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}

//关闭redispool
func (this *RedisCachePool) Close() {
	this.pool.Close()
}

//创建一个redispool
func NewCachePool(ipport, password string, maxactive, maxidle, idletime int)  CacheInterface{
	redispool := &redis.Pool{
		MaxActive:   maxactive,
		MaxIdle:     maxidle,
		IdleTimeout: time.Duration(idletime) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", ipport)
			if err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				panic(err)
			}
			return err

		},
	}
	return &RedisCachePool{pool: redispool}
}
