package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/* комментарий по решению:
скорее всего на практике такое решение в рамках указанного ТЗ не дает никаких преимуществ
в сравнении с выполнением на одной горутине, так как критическая секция находится под одним мьютексом и 4 горутины
превращаются в одну очередь. Так еще и оверхед за счет свича межуд горутинами.
*/

func main() {

	// keynumber - 2026 (очередная предполагаемая дата выхода Half-life 3)
	keynumber := 2026
	gorutinesNum := 4
	customMap := NewMyCustomMap()
	wg := sync.WaitGroup{}

	jobs := genJobs(keynumber)
	fmt.Printf("jobs buffered: %d\n", len(jobs))

	for i := 0; i < gorutinesNum; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for job := range jobs {
				customMap.Lock()
				old, ok := customMap.m[job]
				customMap.CallKeysCnt++
				if !ok {
					customMap.AddKeysCnt++
					old = 0
				}
				customMap.m[job] = old + 1
				customMap.Unlock()
			}
		}()
	}
	wg.Wait()

	checkMap(customMap)
}

func genJobs(keyNumber int) chan int {
	jobs := make([]int, 0, keyNumber*3)

	for i := 1; i <= keyNumber; i++ {
		for j := 0; j < 3; j++ {
			jobs = append(jobs, i)
		}
	}

	rand.Shuffle(len(jobs), func(i, j int) {
		jobs[i], jobs[j] = jobs[j], jobs[i]
	})

	chanJobs := make(chan int, len(jobs))
	for _, v := range jobs {
		chanJobs <- v
	}
	close(chanJobs)
	return chanJobs
}

func checkMap(cm *MyCystomMap) {
	fmt.Printf("CallKeysCnt: %d\n", cm.CallKeysCnt)
	fmt.Printf("AddKeysCnt: %d\n", cm.AddKeysCnt)

	for k, v := range cm.m {
		if v != 3 {
			fmt.Printf("err: key %d has value %d; expected 3\n", k, v)
			return
		}
	}

	fmt.Println("test ok: all values are equal to 3")
}
