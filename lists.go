package redis

import (
	"encoding/json"
	"reflect"
)

func (con *Connector) LPush(key string, values ...interface{}) error {
	total, err := con.client.LPush(key, values...).Result()
	if err != nil {
		return err
	}
	con.Logger.Infof("LPush - %d items", total)
	// Success
	return nil
}

func (con *Connector) LPop(key string, rtype reflect.Kind) (interface{}, error) {
	res := con.client.LPop(key)
	// Success
	return responseMapping(*res, rtype)
}

func (con *Connector) LPopStruct(key string, pointer interface{}) error {
	bytes, err := con.LPop(key, reflect.Invalid)
	if err != nil {
		return err
	}
	// Success
	return json.Unmarshal(bytes.([]byte), pointer)
}

func (con *Connector) RPush(key string, values ...interface{}) error {
	total, err := con.client.RPush(key, values...).Result()
	if err != nil {
		return err
	}
	con.Logger.Infof("RPush - %d items", total)
	// Success
	return nil
}

func (con *Connector) RPop(key string, rtype reflect.Kind) (interface{}, error) {
	res := con.client.RPop(key)
	// Success
	return responseMapping(*res, rtype)
}

func (con *Connector) RPopStruct(key string, pointer interface{}) error {
	bytes, err := con.RPop(key, reflect.Invalid)
	if err != nil {
		return err
	}
	// Success
	return json.Unmarshal(bytes.([]byte), pointer)
}
