package features

import "fmt"

func PanicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		} else {
			fmt.Println("Recovered. without error:\n", r)
		}
	}()
	r := getReciprocal(17)
	fmt.Println(r)
}

func getReciprocal(n float64) float64 {
	if n == 0 {
		panic("devide by 0")
	}
	return 1 / n
}
