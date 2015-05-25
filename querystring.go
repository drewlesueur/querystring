package querystring

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
)

// the go-querystring project isn't what I wanted.
// I started applying a patch, and might do that later.

func Valueify(v interface{}) (url.Values, error) {
	values := make(url.Values)

	val := reflect.ValueOf(v)
	err := valueify2(val, values, "")
	return values, err
}

func valueify2(val reflect.Value, values url.Values, scope string) error {

	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	fmt.Println("the kind: ", val.Kind())
	switch val.Kind() {
	case reflect.Bool:
		if val.Bool() {
			values.Add(scope, "true")
		} else {
			values.Add(scope, "false")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		values.Add(scope, strconv.Itoa(int(val.Int())))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		values.Add(scope, strconv.Itoa(int(val.Uint())))
	case reflect.Float32:
		values.Add(scope, strconv.FormatFloat(val.Float(), 'f', -1, 32))
	case reflect.Float64:
		values.Add(scope, strconv.FormatFloat(val.Float(), 'f', -1, 64))
	case reflect.String:
		values.Add(scope, val.String())
	case reflect.Interface:
		fmt.Println("got here!!!??")
	case reflect.Struct:
		// TODO: handle struct
		// see go-querystring for a nice loop example
	case reflect.Map:
		typ := val.Type()

		if typ.Key().Kind() != reflect.String {
			return errors.New("must have string keys")
		}

		myStringSlice := sort.StringSlice{}
		mapKeys := val.MapKeys()
		for _, mapKey := range mapKeys {
			myStringSlice = append(myStringSlice, mapKey.String())
		}
		myStringSlice.Sort()

		for _, mapKey := range myStringSlice {
			mapValue := val.MapIndex(reflect.ValueOf(mapKey))

			// possibly let the above case statement handle this
			if mapValue.Kind() == reflect.Interface { // or pointer?!
				mapValue = mapValue.Elem()
			}

			fmt.Println("yo!", mapValue)
			fmt.Println("yo!", mapValue.Kind() == reflect.String)
			var err error
			if scope == "" {
				err = valueify2(mapValue, values, mapKey)
			} else {
				err = valueify2(mapValue, values, scope+"["+mapKey+"]")
			}
			if err != nil {
				return err
			}
		}
	case reflect.Slice:
	case reflect.Array:
	case reflect.Ptr:
	default:
		fmt.Println("got a type we can't handle")
		return errors.New("can not handle this type")

	}
	return nil
}

func Stringify(v interface{}) (string, error) {
	values, err := Valueify(v)
	if err != nil {
		return "", err
	}
	encoded := values.Encode()
	unescaped, err := url.QueryUnescape(encoded)
	if err != nil {
		return "", err
	}
	return unescaped, nil
}
