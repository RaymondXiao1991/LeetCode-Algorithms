package common

import (
	"reflect"
	"time"
)

type Interface interface {
	DeepCopyIface() interface{}
}

func Iface(iface interface{}) interface{} {
	return DeepCopy(iface)
}

func DeepCopy(src interface{}) interface{} {
	if src == nil {
		return nil
	}

	origin := reflect.ValueOf(src)

	cpy := reflect.New(origin.Type()).Elem()

	copyRecursive(origin, cpy)

	return cpy.Interface()
}

func copyRecursive(origin, cpy reflect.Value) {
	if origin.CanInterface() {
		if copier, ok := origin.Interface().(Interface); ok {
			cpy.Set(reflect.ValueOf(copier.DeepCopyIface()))
		}
	}

	switch origin.Kind() {
	case reflect.Ptr:
		originVal := origin.Elem()
		if !originVal.IsValid() {
			return
		}
		cpy.Set(reflect.New(originVal.Type()))
		copyRecursive(originVal, cpy.Elem())

	case reflect.Interface:
		if origin.IsNil() {
			return
		}
		originVal := origin.Elem()
		copyVal := reflect.New(originVal.Type()).Elem()
		copyRecursive(originVal, copyVal)
		cpy.Set(copyVal)

	case reflect.Struct:
		if t, ok := origin.Interface().(time.Time); ok {
			cpy.Set(reflect.ValueOf(t))
			return
		}
		for i := 0; i < origin.NumField(); i++ {
			if len(origin.Type().Field(i).PkgPath) != 0 {
				continue
			}
			copyRecursive(origin.Field(i), cpy.Field(i))
		}

	case reflect.Slice:
		if origin.IsNil() {
			return
		}
		cpy.Set(reflect.MakeSlice(origin.Type(), origin.Len(), origin.Cap()))
		for i := 0; i < origin.Len(); i++ {
			copyRecursive(origin.Index(i), cpy.Index(i))
		}

	case reflect.Map:
		if origin.IsNil() {
			return
		}
		cpy.Set(reflect.MakeMap(origin.Type()))
		for _, key := range origin.MapKeys() {
			originVal := origin.MapIndex(key)
			copyVal := reflect.New(originVal.Type()).Elem()
			copyRecursive(originVal, copyVal)
			copyKey := DeepCopy(key.Interface())
			cpy.SetMapIndex(reflect.ValueOf(copyKey), copyVal)
		}

	default:
		cpy.Set(origin)

	}
}
