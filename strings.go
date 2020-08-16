package redis

import (
	"encoding/json"
	"reflect"
	"time"
)

func (con *Connector) Get(key string, rtype reflect.Kind) (data interface{}, err error) {
	res := con.client.Get(key)
	// Success
	return responseMapping(*res, rtype)
}

func (con *Connector) GetStruct(key string, pointer interface{}) error {
	bytes, err := con.Get(key, reflect.Invalid)
	if err != nil {
		return err
	}
	// Success
	return json.Unmarshal(bytes.([]byte), pointer)
}

func (con *Connector) Set(key string, value interface{}, expires time.Duration) error {
	// Success
	return con.client.Set(key, value, expires).Err()
}

func (con *Connector) SetStruct(key string, value interface{}, expires time.Duration) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// Success
	return con.Set(key, bytes, expires)
}
