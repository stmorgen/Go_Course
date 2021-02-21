package main

import "fmt"
import "sync"

var wg sync.WaitGroup

type ChopS struct{ sync.Mutex }

type Philo struct {
	id int
	leftCS, rightCS *ChopS
	tableCleared, tablePrepared chan int
}

type Host struct {
	tableCleared, tablePrepared, abort chan int
}

func (h Host) manage() {
   
   h.tablePrepared <- 1
   h.tablePrepared <- 2
   
   for {
	   select {     
		 case t := <- h.tableCleared:
			//fmt.Println("table finished", t)
			h.tablePrepared <- t
		 case <- h.abort:
			//fmt.Println("eating finished")
		    return
	   }
   }
}

func (p Philo) eat() {
   eatOnce(p.id, p.leftCS, p.rightCS, p.tableCleared, p.tablePrepared)
   eatOnce(p.id, p.leftCS, p.rightCS, p.tableCleared, p.tablePrepared)
   eatOnce(p.id, p.leftCS, p.rightCS, p.tableCleared, p.tablePrepared)
   wg.Done()
}

func eatOnce(id int, leftCS, rightCS *ChopS, tableCleared, tablePrepared chan int) {
	
	x := <- tablePrepared
	
	leftCS.Lock()
    rightCS.Lock()

    fmt.Println("starting to eat", id)
	//fmt.Println("at table", x)
	fmt.Println("finishing eating", id)

    rightCS.Unlock()
    leftCS.Unlock()
	
	tableCleared <- x
}

func main() {

	tableCleared := make(chan int, 2)
	tablePrepared := make(chan int, 2)
	abort := make(chan int)
	host := &Host{tableCleared, tablePrepared, abort}

	wg.Add(5)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
	   CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
	   philos[i] = &Philo{i+1, CSticks[i], CSticks[(i+1)%5], tableCleared, tablePrepared}
	}
 
	go host.manage()
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	
	wg.Wait()
	abort <- 1
}