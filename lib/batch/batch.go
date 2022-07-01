package batch

import (
	"time"
	"sync"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	var wg sync.WaitGroup
	res = make([]user, n)
	sem := make(chan struct{}, pool)
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int64) {	
			user := getOne(i)
			res[i] = user
			<-sem
			wg.Done()
		}(i)
	}

	wg.Wait()	

	return
}
