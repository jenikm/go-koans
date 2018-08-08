package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

const x = 0

func aboutPanics() {
  defer func(){
    recover()
  }()

	n := divideFourBy(x)
	assert(n == 2) // panics are exceptional errors at runtime
}
