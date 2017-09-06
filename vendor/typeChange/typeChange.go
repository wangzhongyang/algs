package typeChange

import (
	"reflect"
	"strconv"
)

type TypeChange struct{}

// Create a new typeChange
func New() *TypeChange {
	return &TypeChange{}
}

// ObjToFloat type interface to float64
func (this *TypeChange) ObjToFloat(obj interface{}) float64 {
	var f float64
	switch reflect.TypeOf(obj).Name() {
	case "string":
		f, _ = strconv.ParseFloat(obj.(string), 64)
	}
	return f
}
