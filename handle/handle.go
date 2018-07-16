package handle

import (
	"github.com/huangfly/SimpleRedisMemcache/mysql"
	"github.com/huangfly/SimpleRedisMemcache/cache"
	"log"
)

type Handle struct{
	sqlsvr mysql.SqlSvrInterface
	memcache cache.CacheInterface
}

//创建一个
func NewHandle(sql mysql.SqlSvrInterface, memche cache.CacheInterface) *Handle{
	return &Handle{
		sqlsvr : sql,
		memcache : memche,
	}
}

//将key-value键值对储存到缓存和数据库，如果存在该键值对就更新
func (this *Handle)StoreValue(key, value string) error{
	//开始事物
	tx, err :=this.sqlsvr.DoCmd()
	if err != nil{
		return err
	}
	defer tx.Commit()
	
	//判断缓存是否有值
	_, err = this.memcache.Get(key)
	if err != nil{
		log.Println("StroreValue find value from database")
		_, err := tx.Exec("insert into hf values(?, ?)", key ,value)
		if err != nil{
			tx.Rollback()
			return err
		}
	
	}else{
			log.Println("StroreValue find value from memcache")
		_, err := tx.Exec("update hf set name = ? where serialnum = ?", value,key )
		if err != nil{
			tx.Rollback()
			return err
		}	
	}
	//修改缓存的值
	this.memcache.Set(key, value, 0)
	return nil
}

//删除指定的key-value键值对，删除缓存和数据库
func (this *Handle)Delete(key string)error{
	_, err := this.memcache.Get(key)
	if err != nil{
		return err
	}
	this.memcache.Delete(key)
	this.sqlsvr.Exec("delete from hf where serialnum = "+key)
	return err
}

//获取key对应的的value如果缓存没有去数据库查
func (this *Handle)GetValue(key string) (string,error){
	val, err := this.memcache.Get(key)
	if err != nil {
		log.Println("GetValue find value from database")
		row, err:= this.sqlsvr.Query("select name from hf where serialnum = "+key)
		if err != nil{
			return val,err
		}
		for row.Next(){
			row.Scan(&val)
		}
		 this.memcache.Set(key, val,0)
		 return val, nil
	}
	log.Println("GetValue find value from memcache")
	return val, nil
}