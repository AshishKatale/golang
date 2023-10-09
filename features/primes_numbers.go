package features

func IsPrime(n uint) (primeFlag bool) {
	primeFlag = true

	switch {
	case n < 2:
		primeFlag = false
		return
	case n == 2 || n == 3:
		return
	case (n-1)%6 != 0 && (n+1)%6 != 0:
		primeFlag = false
		return
	}

	for i := uint(3); i < n/3; i += 2 {
		if n%i == 0 {
			primeFlag = false
			break
		}
	}

	return
}

func checkIfPrime(n int) (primeFlag bool) {
	primeFlag = true
	for i := 3; i < n/2; i += 2 {
		if n%i == 0 {
			primeFlag = false
			break
		}
	}
	return
}

func GetPrimesInRangeParallel(low, high, jobs int) int {
	c := make(chan int)
	block := (high - low) / jobs
	total := 0

	for i := 0; i < jobs; i++ {
		go getPrimesInRangeParallel(i*block+1, (i+1)*block, c)
	}

	for i := 0; i < jobs; i++ {
		n := <-c
		total += n
	}

	return total
}

func getPrimesInRangeParallel(low, high int, receiver chan<- int) {
	primes := 0

	if low <= 2 {
		primes += 2
	} else if low <= 3 {
		primes++
	}

	var i int
	if low < 6 {
		i = 6
	} else {
		i = low / 6 * 6
	}

	for i <= high {
		j := i - 1
		k := i + 1

		if j >= low && j <= high && checkIfPrime(j) {
			primes++
		}
		if k >= low && k <= high && checkIfPrime(k) {
			primes++
		}

		i += 6
	}

	receiver <- primes
}

func GetPrimesInRange(low, high int) int {
	primes := 0

	if low <= 2 {
		primes += 2
	} else if low <= 3 {
		primes++
	}

	var i int
	if low < 6 {
		i = 6
	} else {
		i = low / 6 * 6
	}

	for i <= high {
		j := i - 1
		k := i + 1

		if j >= low && j <= high && checkIfPrime(j) {
			primes++
		}
		if k >= low && k <= high && checkIfPrime(k) {
			primes++
		}

		i += 6
	}

	return primes
}
