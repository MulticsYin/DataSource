package redis

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var redisPool *redis.Pool

func InitRedis(redisHost, redisPassword string, maxIdle, maxActive, idleTimeout int) {
	redisPool = &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		MaxActive:   maxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			if redisPassword != "" {
				if _, err := c.Do("AUTH", redisPassword); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func CloseRedisPool() {
	redisPool.Close()
}

// 通过缓存Key从Redis获取数据
func GetString(key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

// 设置缓存
func SetEX(key, value string, timeout int) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()

	return redis.String(conn.Do("SETEX", key, timeout, value))
}

// 删除缓存
func Delete(key string) (interface{}, error) {
	conn := redisPool.Get()
	defer conn.Close()

	//return redis.String(conn.Do("DEL", key))
	return conn.Do("DEL", key)
}

// 将值插入到列表头部
func PushListHead(key, value string) (interface{}, error) {
	conn := redisPool.Get()
	defer conn.Close()

	return conn.Do("LPUSH", key, value)
}

// 移出并获取列表的第一个元素
func PopListHead(key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()

	return redis.String(conn.Do("LPOP", key))
}

func PushListEnd(key, value string) (interface{}, error) {
	conn := redisPool.Get()
	defer conn.Close()

	return conn.Do("RPUSH", key, value)
}

// 移除列表的最后一个元素，返回值为移除的元素
func PopListEnd(key string) (string, error) {
	conn := redisPool.Get()
	defer conn.Close()

	return redis.String(conn.Do("RPOP", key))
}

// 接口方法
// Do sends a command to the server and returns the received reply.
func Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := redisPool.Get()
	defer conn.Close()

	return conn.Do(commandName, args)
}

// 接口方法
// Send writes the command to the client's output buffer.
func Send(commandName string, args ...interface{}) error {
	conn := redisPool.Get()
	defer conn.Close()

	return conn.Send(commandName, args)
}
