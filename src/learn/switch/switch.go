package main

//基础switch
//func main() {
//	fmt.Print("Go runs on ")
//	switch os := runtime.GOOS; os {
//	case "darwin":
//		fmt.Println("OS X.")
//	case "linux":
//		fmt.Println("Linux.")
//	default:
//		// freebsd, openbsd,
//		// plan9, windows...
//		fmt.Printf("%s.", os)
//	}
//}

//switch执行顺序
//func main() {
//	fmt.Println("When's Saturday?")
//	today := time.Now().Weekday()
//	switch time.Saturday {
//	case today + 0:
//		fmt.Println("Today.")
//	case today + 1:
//		fmt.Println("Tomorrow.")
//	case today + 2:
//		fmt.Println("In two days.")
//	default:
//		fmt.Println("Too far away.")
//	}
//}

//没有条件的 switch 同 `switch true` 一样。
//func main() {
//	t := time.Now()
//	switch {
//	case t.Hour() < 12:
//		fmt.Println("Good morning!")
//	case t.Hour() < 17:
//		fmt.Println("Good afternoon.")
//	default:
//		fmt.Println("Good evening.")
//	}
//}
