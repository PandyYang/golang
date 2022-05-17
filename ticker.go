package main

//func main() {
//
//	//counter := 1
//
//	ticker := time.NewTicker(time.Second)
//
//	//for _ = range ticker.C {
//	//	fmt.Println("ticker...")
//	//	counter++
//	//	if counter > 5 {
//	//		ticker.Stop()
//	//		break
//	//	}
//	//}
//
//	chanInt := make(chan int)
//
//	go func() {
//		for _ = range ticker.C {
//			select {
//			case chanInt <- 1:
//				fmt.Println("send 1")
//			case chanInt <- 2:
//				fmt.Println("send 2")
//			case chanInt <- 3:
//				fmt.Println("send 3")
//			}
//		}
//	}()
//
//	sum := 0
//	for v := range chanInt {
//		fmt.Printf("read v: %v\n", v)
//		sum += v
//		if sum >= 10 {
//			break
//		}
//	}
//
//}
