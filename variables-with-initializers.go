package main

import (
	"fmt"
	"reflect"
)

var i, j int = 1, 2

func usedVariablesInitializers() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	showTypesAll(c, python, java)
}

func showTypesAll(c interface{}, python interface{}, java interface{}) {
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(j), reflect.TypeOf(c),
		reflect.TypeOf(python), reflect.TypeOf(java))
}
