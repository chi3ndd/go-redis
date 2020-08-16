package redis

import (
	"reflect"
	"time"

	r "github.com/go-redis/redis"
)

func responseMapping(res r.StringCmd, rtype reflect.Kind) (data interface{}, err error) {
	switch rtype {
	case reflect.String:
		data, err = res.Result()
	case reflect.Int:
		data, err = res.Int()
	case reflect.Int64:
		data, err = res.Int64()
	case reflect.Uint64:
		data, err = res.Uint64()
	case reflect.Float32:
		data, err = res.Float32()
	case reflect.Float64:
		data, err = res.Float64()
	default:
		data, err = res.Bytes()
	}
	// Success
	return data, err
}

func (con *Connector) Expire(key string, expires time.Duration) (bool, error) {
	/* Return
	1: Success
	0: Key not found
	*/
	return con.client.Expire(key, expires).Result()
}

func (con *Connector) ExpireAt(key string, expires time.Time) (bool, error) {
	/* Return
	1: Success
	0: Key not found
	*/
	return con.client.ExpireAt(key, expires).Result()
}

func (con *Connector) Exists(key ...string) (int64, error) {
	/* Return
	number: Number of existed keys
	*/
	return con.client.Exists(key...).Result()
}

func (con *Connector) Del(key ...string) (int64, error) {
	/* Return
	number: Number of deleted keys
	*/
	return con.client.Del(key...).Result()
}
