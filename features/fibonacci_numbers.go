package features

func FibonacciClosure() func() uint {
	a, b := uint(0), uint(1)
	return func() uint {
		defer func() { a, b = b, a+b }()
		return a
	}
}

var fibMap = map[uint]uint{
	0: 0,
	1: 1,
}

func Fibonacci(n uint) uint {
	fib, ok := fibMap[n]
	if ok {
		return fib
	}
	fibMap[n] = Fibonacci(n-1) + Fibonacci(n-2)
	return fibMap[n]
}
