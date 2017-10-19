package main

// xiaorui.cc
import (
	"fmt"
	"time"
)

func testTimer1() {
	go func() {
		fmt.Println("test timer1")
	}()

}

func testTimer2() {
	go func() {
		fmt.Println("test timer2")
	}()
}

func timer1() {
	timer1 := time.NewTicker(2 * time.Second)
	select {
	case <-timer1.C:
		testTimer1()
	}

}

func timer2() {
	timer2 := time.NewTicker(3 * time.Second)

	select {
	case <-timer2.C:
		testTimer2()
	}

}

func main() {
	go timer1()
	go timer2()
	fmt.Println("will end")
	time.Sleep(5 * time.Second)

	//ticker := time.Tick(2 * time.Second)
	//i := 0
	//for range ticker {
	//	i++
	//	fmt.Printf("time %d\n", i)
	//}

	timer := time.NewTimer(1*time.Second)
	i := 0
	for{
		select {
			case tmp := <-timer.C:
				i++
				fmt.Printf("newtime %d ch %+v\n", i, tmp)
				//timer stop 后会重新 start
				timer.Reset(2*time.Second)
		}
	}
}
