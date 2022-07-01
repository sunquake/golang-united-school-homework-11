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
	if n==0 || pool==0 {return nil}
	var wg sync.WaitGroup

	sem := make(chan struct{}, pool)
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int64) {	
			user := getOne(i)
			res = append(res, user)
			<-sem
			wg.Done()
		}(i)
	}

	wg.Wait()	

	return
}
