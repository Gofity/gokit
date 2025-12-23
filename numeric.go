package gokit

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

type Num[T Numeric] struct {
	Data T
}

func (x Num[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Data)
}

func (x *Num[T]) UnmarshalJSON(data []byte) error {
	return x.Scan(data)
}

func (x *Num[T]) Scan(v any) (err error) {
	switch data := v.(type) {
	case T:
		x.Data = data

	case []byte:
		data = bytes.Trim(data, "'\" ")
		return json.Unmarshal(data, &x.Data)

	case string:
		data = strings.Trim(data, "'\" ")
		return json.Unmarshal([]byte(data), &x.Data)

	case int:
		x.Data = T(data)
		return

	case int8:
		x.Data = T(data)
		return

	case int16:
		x.Data = T(data)
		return

	case int32:
		x.Data = T(data)
		return

	case int64:
		x.Data = T(data)
		return

	case uint:
		x.Data = T(data)
		return

	case uint8:
		x.Data = T(data)
		return

	case uint16:
		x.Data = T(data)
		return

	case uint32:
		x.Data = T(data)
		return

	case uint64:
		x.Data = T(data)
		return

	case float32:
		x.Data = T(data)
		return

	case float64:
		x.Data = T(data)
		return
	}

	err = errors.New("Unsupported numeric type")

	return
}

func (x Num[T]) Value() (driver.Value, error) {
	return x.Data, nil
}

func (x *Num[T]) Get() T {
	return x.Data
}

func (x *Num[T]) String() string {
	return fmt.Sprint(x.Data)
}
