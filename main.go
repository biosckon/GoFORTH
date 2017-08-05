package main

import (
	"fmt"
)

var stk = make(map[interface{}]interface{})
var lst = []string{
	"a", "av",
	"b", "bv",
	"c", "cv",
	"d", "dv",
}

func f1() {
	fmt.Printf("func f1 type %T\n", f1)
}

func main() {

	// populate the map
	//
	i := 0
	l := len(lst) - 1
	for {
		if i > l {
			break
		}
		stk[lst[i]] = lst[i+1]
		i += 2
	}
	stk["f"] = f1

	// pull the values from the
	//
	for k := range stk {
		fmt.Printf("%v = %v\n", k, stk[k])
	}

	switch f := stk["f"].(type) {
	case func():
		f()
	}
}
