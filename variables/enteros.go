package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("int:", intcomun)
	fmt.Println("int32:", intde32)
	fmt.Println("intde64:", intde64)
}
