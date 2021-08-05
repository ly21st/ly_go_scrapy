package utils

import (
	"errors"
	"fmt"
	"strings"

	logp "yannscrapy/logger"
)

/**
读取yml文件配置
*/
func ReadYml(root interface{}, key string) (interface{}, error) {
	switch root.(type) {
	case map[interface{}]interface{}:
		keyArr := strings.Split(key, ".")
		if len(keyArr) == 0 {
			return root, nil
		}
		if (len(keyArr) == 1) && keyArr[0] == "" {
			return root, nil
		}
		m, ok := root.(map[interface{}]interface{})
		if !ok {
			msg := fmt.Sprintf("type:%T error", root)
			logp.Infof(msg)
			return nil, errors.New(msg)
		}
		value, ok := m[keyArr[0]]
		if !ok {
			return nil, nil
		}
		if len(keyArr) == 1 {
			return value, nil
		} else {
			key = strings.Join(keyArr[1:], ".")
			return ReadYml(value, key)
		}
	default:
		msg := fmt.Sprintf("unknown type:%T", root)
		logp.Infof(msg)
		return nil, errors.New(msg)
	}

}

/**
把yml文件转换成对象
*/
func BuildGoObject(root interface{}) (err error) {
	switch root.(type) {
	case []interface{}:
		arr, _ := root.([]interface{})
		for _, arrValue := range arr {
			BuildGoObject(arrValue)
		}
	case map[interface{}]interface{}:
		m, _ := root.(map[interface{}]interface{})
		for key, value := range m {
			key, ok := key.(string)
			if !ok {
				msg := fmt.Sprintf("key error,key:%v,key type:%T", key, key)
				logp.Infof(msg)
				return errors.New(msg)
			}
			keyArr := strings.Split(key, ".")
			if (len(keyArr) == 0) || (len(keyArr) == 1) {
				switch value.(type) {
				case []interface{}:
					err = BuildGoObject(value)
					if err != nil {
						return err
					}
				case map[interface{}]interface{}:
					err = BuildGoObject(value)
					if err != nil {
						return err
					}
				default:
				}
			} else {
				firstKey := keyArr[0]
				nextKey := strings.Join(keyArr[1:], ".")
				newValue := m[firstKey]
				if newValue == nil {
					newValue = map[interface{}]interface{}{
						nextKey: value,
					}
					m[firstKey] = newValue
				} else {
					newValueObj, _ := newValue.(map[interface{}]interface{})
					newValueObj[nextKey] = value
					m[firstKey] = newValueObj
				}
				delete(m, key)
				err = BuildGoObject(newValue)
				if err != nil {
					return err
				}

			}
		}
	default:

	}
	return nil
}
