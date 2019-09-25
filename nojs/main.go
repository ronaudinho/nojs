package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ronaudinho/nojs"
)

// idc float for the moment lol
type StructParent struct {
	Str         string                 `nojs:"str,reversed"`
	Strs        []string               `nojs:"strs,reversed"` // does not reverse order yet
	Int         int                    `nojs:"int,reversed"`
	Ints        []int                  `nojs:"ints,reversed"` // does not reverse order yet
	Bool        bool                   `nojs:"bool,reversed"`
	Bools       []bool                 `nojs:"bools,reversed"` // does not reverse order yet
	Map         map[string]interface{} `nojs:"map,reversed"`   // does not reverse order yet
	StructChild StructChild            `nojs:"structchild,reversed"`
}

type StructChild struct {
	Str  string `nojs:"str,reversed"`
	Int  int    `nojs:"int,reversed"`
	Bool bool   `nojs:"bool,reversed"`
}

func main() {
	/*
		sp := StructParent{
			Str: "json",
			Strs: []string{
				"json",
				"json",
			},
			Int: 1234,
			Ints: []int{
				1234,
				5678,
			},
			Bool:  true,
			Bools: []bool{true, false},
			Map: map[string]interface{}{
				"json":     "json",
				"jsonint":  1234,
				"jsonbool": true,
				"jsonmap": map[string]interface{}{
					"json":     "json",
					"jsonint":  1234,
					"jsonbool": true,
				},
			},
			StructChild: StructChild{
				Str:  "json",
				Int:  1234,
				Bool: true,
			},
		}

			m, _ := nojs.Marshal(sp)
			fmt.Println("marshalled: ", string(m))
			mi, _ := nojs.MarshalIndent(sp, "", "  ")
			fmt.Println("marshalled: ", string(mi))

			var newsp StructParent
			_ = nojs.Unmarshal(m, &newsp)
			fmt.Println("unmarshalled: ", newsp)
	*/

	f, _ := ioutil.ReadFile("nojs.json")
	var fsp StructParent
	_ = nojs.Unmarshal(f, &fsp)
	fmt.Println("from json: ", fsp)
	fi, _ := nojs.MarshalIndent(fsp, "", "  ")
	fmt.Println("marshalled from json: ", string(fi))
}
