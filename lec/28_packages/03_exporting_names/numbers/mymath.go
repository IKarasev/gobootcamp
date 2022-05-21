package numbers

// note: package name stays is same
// only files with same package name allowed in one directory

func IsPrime(n int) bool {
	for i := 2; i <= int(n/2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
