package base

import (
	"sync/atomic"
)

var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

//func main() {
//	for i := 0; i < 100; i++ {
//		go add()
//		go sub()
//	}
//
//	time.Sleep(time.Second * 2)
//	fmt.Printf("end is: %v\n", i)
//}
