package main

import "fmt"

func dataProducer() <- chan int {
	result := make(chan int, 5)

	go func() {
		defer close(result)
		for i := 0; i <= 5; i++ {
			result <- i
		}
	}()
	return result
}

func dataConsumer(data <- chan int) {
	for item := range data {
		fmt.Printf("Got %d\n", item)
	}
	fmt.Println("Got all the data")
}

func main() {
	data := dataProducer()
	dataConsumer(data)
}
