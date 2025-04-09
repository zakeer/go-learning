package main

import (
	"math"
)

const str string = "Syed Zakeer Hussain"

func main() {
	println(str)

	const n = 500000000

	const d = 3e20 / n
	println(d)

	println(int64(d))

	println(math.Sin(n))
}
