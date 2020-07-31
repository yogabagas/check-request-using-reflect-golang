package main

import (
	"fmt"
	"reflect"
)

type Request struct {
	Long, Wide, High float64
}

func (r *Request) Cube() float64 {

	allow := r.validatorRequest()
	if !allow {
		fmt.Println("Not Allowed")
		return 0.0
	}
	return r.Long * r.Wide * r.High
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
	req1 := Request{
		Long: 0,
		Wide: 15.0,
		High: 5.00,
	}

	resp1 := req1.Cube()
	fmt.Println("Result of Cube, If Request has an empty value", resp1)

	req2 := Request{
		Long: 10.0,
		Wide: 20.0,
		High: 15.0,
	}

	resp2 := req2.Cube()
	fmt.Println("Result of Cube, If Request is fulfilled", resp2)

}
