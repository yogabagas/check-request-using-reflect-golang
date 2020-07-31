package main

import (
	"fmt"
	"reflect"
)

type Request struct {
	Long, Wide, High float64
}

func (r *Request) Cube() (resp float64, err error) {

	allow := r.validatorRequest()
	if !allow {
		return 0.0, err
	}
	return r.Long * r.Wide * r.High, nil
}

func (r *Request) validatorRequest() bool {
	var allow = make(map[int]bool)
	var reflectValue = reflect.ValueOf(r)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	// var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		if reflectValue.Field(i).Interface() == 0.0 {
			allow[i] = false
		} else {
			allow[i] = true
		}
	}

	for _, v := range allow {
		if v == true {
			return true
		}
	}
	return false
}

func main() {
	req := Request{
		Long: 0,
		Wide: 15.0,
		High: 5.00,
	}
	resp, err := req.Cube()
	if err != nil {
		return
	}
	fmt.Println("HASIL", resp)

}
