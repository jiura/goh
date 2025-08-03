package goh

import (
	"errors"
)

/* TYPES START */

type Json map[string]any
type JsonArray []any

/* TYPES END */

func (j *Json) GetJson(key string) *Json {
	if j == nil {
		return nil
	}

	val, ok := (*j)[key]
	if !ok {
		return nil
	}

	obj, ok := val.(Json)
	if !ok {
		return nil
	}

	return &obj
}

func (j *Json) GetArray(key string) JsonArray {
	if j == nil {
		return nil
	}

	val, ok := (*j)[key]
	if !ok {
		return nil
	}

	arr, ok := val.(JsonArray)
	if !ok {
		return nil
	}

	return arr
}

func (j *Json) GetString(key string) (string, error) {
	if j == nil {
		return "", errors.New("Invalid JSON object")
	}

	val, ok := (*j)[key]
	if !ok {
		return "", errors.New("Value not found in JSON object")
	}

	str, ok := val.(string)
	if !ok {
		return "", errors.New("Value not a string")
	}

	return str, nil
}

func (j *Json) GetInt(key string) (int, error) {
	if j == nil {
		return 0, errors.New("Invalid JSON object")
	}

	val, ok := (*j)[key]
	if !ok {
		return 0, errors.New("Value not found in JSON object")
	}

	num, ok := val.(int)
	if !ok {
		return 0, errors.New("Value not an int")
	}

	return num, nil
}

func (j *Json) GetFloat(key string) (float64, error) {
	if j == nil {
		return 0.0, errors.New("Invalid JSON object")
	}

	val, ok := (*j)[key]
	if !ok {
		return 0.0, errors.New("Value not found in JSON object")
	}

	num, ok := val.(float64)
	if !ok {
		return 0.0, errors.New("Value not a float64")
	}

	return num, nil
}

func (ja *JsonArray) JsonAt(index int) *Json {
	if index < 0 || ja == nil || index >= len(*ja) {
		return nil
	}

	obj, ok := (*ja)[index].(Json)
	if !ok {
		return nil
	}

	return &obj
}

func (ja *JsonArray) ArrayAt(index int) JsonArray {
	if index < 0 || ja == nil || index >= len(*ja) {
		return nil
	}

	arr, ok := (*ja)[index].(JsonArray)
	if !ok {
		return nil
	}

	return arr
}

func (ja *JsonArray) StringAt(index int) (string, error) {
	if index < 0 || ja == nil || index >= len(*ja) {
		return "", errors.New("Invalid JsonArray or index")
	}

	str, ok := (*ja)[index].(string)
	if !ok {
		return "", errors.New("Not a string")
	}

	return str, nil
}

func (ja *JsonArray) IntAt(index int) (int, error) {
	if index < 0 || ja == nil || index >= len(*ja) {
		return 0, errors.New("Invalid JsonArray or index")
	}

	num, ok := (*ja)[index].(int)
	if !ok {
		return 0, errors.New("Not an int")
	}

	return num, nil
}

func (ja *JsonArray) FloatAt(index int) (float64, error) {
	if index < 0 || ja == nil || index >= len(*ja) {
		return 0, errors.New("Invalid JsonArray or index")
	}

	num, ok := (*ja)[index].(float64)
	if !ok {
		return 0, errors.New("Not a float64")
	}

	return num, nil
}
