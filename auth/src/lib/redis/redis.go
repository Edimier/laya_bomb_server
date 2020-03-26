package redis

import(
	"sync"
	"config"
	"github.com/go-redis/redis"
	"time"
	// "log"
)

var lock sync.Mutex
var instanceConn * redis.Client	

func instance() * redis.Client{
	if instanceConn == nil {
		lock.Lock()
		if instanceConn == nil {
			redisAddr, err := config.GetConfigString("redisAddr")
			if err != nil{ panic(err)}
			redisPwd, err := config.GetConfigString("redisPwd")
			if err != nil{ panic(err)}
			redisDb, err := config.GetConfigInt("redisDb")
			if err != nil{ panic(err)}


			conn := redis.NewClient(&redis.Options{
        		Addr:     redisAddr,
        		Password: redisPwd,
        		DB:       redisDb,
    		})
    		if conn != nil{
    			instanceConn = conn
    		}
		}
		lock.Unlock()
	}
	if instanceConn == nil{
        panic("redis init failed!")
    }
	return instanceConn
}


func IsNil (e error) bool {
	if e.Error() == "redis: nil"{
		return true
	}
	return false
}

func Set(key, value string) error {
	if err := instance().Set(key, value, 0).Err();err != nil{
		return err
	}
	return nil
}

func SetAndExpire(key, value string, expire time.Duration) error {
	if err := instance().Set(key, value, expire).Err();err != nil{
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	value, err := instance().Get(key).Result()
	// log.Println("Get", value, err)
	if err != nil{
		return "", err
	}
	return value, nil
}

func HGet(key, field string) (string, error) {
	value, err := instance().HGet(key, field).Result()
	// log.Println("Get", value, err)
	if err != nil{
		return "", err
	}
	return value, nil
}

func HSet(key, field, value string) error {
	if err := instance().HSet(key, field, value).Err(); err != nil{
		return err
	}
	return nil
}

func HIncrBy(key, field string, value int) error {
	if err:= instance().HIncrBy(key, field, int64(value)).Err(); err != nil{
		return err
	}
	return nil
}
