package main

import (
	"fmt"
	"reflect"
	"strings"
)

func sqlinject(sql string, in ...interface{}) (sqlOut string, params []interface{}) {
	if in == nil {
		return "", nil
	}

	var str string = ""
	var nstr = strings.Split(sql, "?")
	idx := 0
	var result []interface{}
	for i := 0; i < len(in); i++ {
		p := reflect.ValueOf(in[i])
		str = str + nstr[idx]
		idx++
		if p.Kind() == reflect.Slice {
			for j := 0; j < p.Len(); j++ {
				result = append(result, p.Index(j).Int())
				str = str + "?"
				if j < (p.Len() - 1) {
					str = str + ","
					fmt.Println(j, p.Len())
				}
			}
		} else {
			result = append(result, p.Int())
			str = str + "?"
		}
	}
	fmt.Println(len(nstr), idx)

	if len(nstr) == idx+1 {
		str = str + nstr[idx]
	}
	return str, result
}

func main() {
	s, r := sqlinject("SELECT * FROM ? WHERE id=?, num=? ORDER BY id", 2, []int{3, 4, 5}, 5)
	fmt.Println("Result: ", s) //Result:  SELECT * FROM ? WHERE id=?,?,?, num=? ORDER BY id
	fmt.Println("Result: ", r) //Result:  [2 3 4 5 5]

	//Test easyJson
	p := P{
		count: 1,
		name:  "Test json",
		value: 0.50,
	}

	fmt.Println(p.MarshalJSON())

	return
}
