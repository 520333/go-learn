package goConcurrency

import (
	"math/rand"
	"time"
)

func TimeA() {
	t := time.NewTimer(time.Second)
	println("Set the timer,\ttime is ", time.Now().String())
	now := <-t.C
	println("The time is up, time is ", now.String())
}
func TimeB() {
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Intn(10)
			time.Sleep(400 * time.Millisecond)
		}
	}()
	t := time.NewTimer(time.Second * 3)
	counter, hint, miss, result := 5, 0, 0, 4
loopGuess:
	for i := 0; i < counter; i++ {
		for {
			select {
			case v := <-ch:
				println("Guess Num is", v)
				if result == v {
					println("Bingo! Hint the number.", result)
					hint++
					t.Reset(time.Second * 3)
					break loopGuess
				}
			case <-t.C:
				println("The time is up, don't hint.")
				miss++
				t = time.NewTimer(time.Second * 3)
				break loopGuess
			}
		}
	}
	println("Game Completed!", "Hint", hint, "Miss", miss, "result", result)
}
func TimeC() {
	ch := time.After(time.Second * 2)
	println("Set the timer,\ttime is ", time.Now().String())
	now := <-ch
	println("The time is up, time is ", now.String())
}

func TickerA() {
	ticker := time.NewTicker(time.Second) //断续器
	timer := time.After(time.Second * 5)  //定时器
loopFor:
	for t := range ticker.C { //周期性执行
		println("HeartBeat, http.Get(https://domain/ping), time", t.String())
		select {
		case <-timer:
			println("HeartBeat Completed. ")
			ticker.Stop()
			break loopFor
		default:
		}
	}

}
