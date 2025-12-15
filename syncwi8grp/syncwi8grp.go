package syncwi8grp

import (
	"fmt"
	"sync"
)

//basically checklist
//after all go routiens then func starts

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() //signal to go routien
	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d ended\n", i)

}

func LearnsyncWi8() {
	fmt.Println("learn sync wi8")
	var wg sync.WaitGroup

	//start 3 workers

	for i := 0; i < 3; i++ {
		wg.Add(1) //incremnet of wi8grpcounter
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("worker task completed")
}
