package fibonacci

// fibonacci 函数会返回一个返回 int 的函数。
func Fibonacci() func() int {
	fn1, fn2 := 0, 1
	return func() int {
		fn2 = fn1 + fn2
		fn1 = fn2 - fn1
		return fn1
	}
}

// func main() {
// 	f := Fibonacci()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }
