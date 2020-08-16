package redis

import (
	"encoding/json"
	"reflect"
)

func (con *Connector) HGetAll(key string) (map[string]string, error) {
	// Success
	return con.client.HGetAll(key).Result()
}

func (con *Connector) HGetAllStruct(key string, pointer interface{}) error {
	res, err := con.HGetAll(key)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	// Success
	return json.Unmarshal(bytes, pointer)
}

func (con *Connector) HGet(key, field string, rtype reflect.Kind) (interface{}, error) {
	res := con.client.HGet(key, field)
	// Success
	return responseMapping(*res, rtype)
}

func (con *Connector) HGetStruct(key, field string, pointer interface{}) error {
	bytes, err := con.HGet(key, field, reflect.Invalid)
	if err != nil {
		return err
	}
	// Success
	return json.Unmarshal(bytes.([]byte), pointer)
}

func (con *Connector) HSet(key, field string, value interface{}) error {
	// Success
	return con.client.HSet(key, field, value).Err()
}

func (con *Connector) HSetStruct(key, field string, value interface{}) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	// Success
	return con.HSet(key, field, bytes)
}

func (con *Connector) HDel(key string, field ...string) error {
	// Success
	return con.client.HDel(key, field...).Err()
}

func (con *Connector) HKeys(key string) ([]string, error) {
	// Success
	return con.client.HKeys(key).Result()
}

func (con *Connector) HExists(key, field string) (bool, error) {
	// Success
	return con.client.HExists(key, field).Result()
}
