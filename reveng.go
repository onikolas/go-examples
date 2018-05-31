package main

func main() {

}

// used to find out the number of cores on the machine
func FindCores(n int) {
	var a int64 = 7
	for i := 0; i < n; i++ {
		a = a * a
	}
}

// Used to estimate the size of the CPUs cache
func FindCache(n int) {

}
