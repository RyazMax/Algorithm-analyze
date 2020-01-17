package main

import (
	"log"
	"os"
	"sync"
	"time"
)

type transferObject struct {
	id  int
	obj interface{}
}

type job func(in, out chan transferObject)

// ExecutePipeline executes pipeline of jobs
func ExecutePipeline(chanSize int, args ...job) {
	wg := &sync.WaitGroup{}
	// Создаем входной канал размера chanSize
	in := make(chan transferObject, chanSize)
	for _, val := range args {
		// Создаем входной канал размера chanSize
		out := make(chan transferObject, chanSize)
		// Инкрементируем число запущенных исполнителей
		wg.Add(1)
		go func(in, out chan transferObject, f job) {
			defer wg.Done()  // Декрементируем число запущенных исполнителей
			defer close(out) // Закрываем выходной канал
			f(in, out)       // Запуск исполнителя
		}(in, out, val)
		in = out // Заменяем входной канал следующего, выходным текущего
	}
	// Ждем завершения всех горутин
	wg.Wait()
}

// CreateSleeperJob creates function of job type that sleeps for seconds seconds
func CreateSleeperJob(seconds, delta, id int) job {
	return func(in, out chan transferObject) {
		for val := range in {
			log.Printf("Worker %d: got object with id = %d\n",
				id, val.id)
			time.Sleep(time.Duration(seconds) * time.Second)
			log.Printf("Worker %d: sending object with id = %d\n",
				id, val.id)
			out <- val
		}
	}
}

// CreateWritingJob creates function of job type that writes obj into out chan
func CreateWritingJob(objs []transferObject) job {
	return func(in, out chan transferObject) {
		for _, val := range objs {
			log.Printf("Writer: sending object with id = %d\n", val.id)
			out <- val
		}
	}
}

// CreateReadingJob creates function of type job that reads in in objs slice
func CreateReadingJob(objs []transferObject) job {
	return func(in, out chan transferObject) {
		count := 0
		for val := range in {
			log.Printf("Reader: got object with id = %d\n",
				val.id)
			objs[count] = val
			count++
		}
	}
}

func main() {
	f, err := os.OpenFile("main.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Can't open log file with err ", err)
	}
	defer f.Close()

	log.SetOutput(f)

	// profile, err := os.OpenFile("profile.3", os.O_WRONLY|os.O_CREATE, 0664)

	size := 10
	a := make([]transferObject, size)
	for i := 0; i < size; i++ {
		a[i] = transferObject{i, nil}
	}
	b := make([]transferObject, size)

	ExecutePipeline(0,
		CreateWritingJob(a),
		CreateSleeperJob(1, 0, 1),
		CreateSleeperJob(2, 0, 2),
		CreateSleeperJob(1, 0, 3),
		CreateReadingJob(b),
	)
}
