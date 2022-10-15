package base

var c = make(chan int)

//func main() {
//
//	go func() {
//		for i := 0; i < 2; i++ {
//			c <- i
//		}
//
//		// 如果没有关闭通道 会造成死锁
//		close(c)
//	}()
//
//	// for range的形式遍历通道
//	//for i2 := range c {
//	//	fmt.Printf("r: %v\n", i2)
//	//}
//
//	for {
//		v, ok := <-c
//		if ok {
//			fmt.Printf("r: %v\n", v)
//		} else {
//			break
//		}
//	}
//
//	//r := <-c
//	//fmt.Printf("r: %v\n", r)
//	//
//	//r = <-c
//	//fmt.Printf("r: %v\n", r)
//	//
//	//r = <-c
//	//fmt.Printf("r: %v\n", r)
//}
