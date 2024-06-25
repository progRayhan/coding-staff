package structs

import "fmt"

type A struct {
	FieldA int
}

type B struct {
	A
	FieldB int
}

func CallB() {
	aa := B{
		A: A{
			FieldA: 10,
		},
		FieldB: 20,
	}
	fmt.Println(aa.A)
	fmt.Println(aa.A.FieldA)
	fmt.Println(aa.FieldB)
}
