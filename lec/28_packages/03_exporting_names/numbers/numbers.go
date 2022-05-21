package numbers

import "fmt"

// note: func name starts with capital letter
// only funcs started with cap 1 letter can be accessed from other
// packages
func Even(n uint) bool {
	return n%2 == 0
}

func notGlobalFunc() {
	fmt.Println("non global func")
}

// funcs started with lower case are private to the package
func odd(n int) bool {
	return n%2 != 0
}

func isPrime(n int) bool {
	for i := 2; i <= int(n/2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// private funcs are used only inside package
func OddAndPrime(n int) bool {
	return odd(n) && isPrime(n)
}
