package producers_consumer

import (
	"fmt"
	"github.com/Aidann32/concurrency_exercises/exercises/files/fetcher"
	"sync"
)

type ProducersConsumer struct {
	wg            sync.WaitGroup
	sharedChannel chan float64
	result        float64
	itemQuantity  int
	fetcher       fetcher.Fetcher
}

func (pr *ProducersConsumer) getResult(input fetcher.FetchOutput) (result float64) {
	switch input.Operation {
	case "+":
		return input.Operand1 + input.Operand2
	case "-":
		return input.Operand1 - input.Operand2
	case "*":
		return input.Operand1 * input.Operand2
	case "/":
		return input.Operand1 / input.Operand2
	default:
		return 0
	}
}

func (pr *ProducersConsumer) fileProducers() {
	for i := 0; i < pr.itemQuantity; i++ {
		pr.wg.Add(1)
		i := i
		go func() {
			defer pr.wg.Done()
			pr.sharedChannel <- pr.getResult(pr.fetcher.Fetch(fmt.Sprintf("in_%d.dat", i)))
			fmt.Printf("Goroutine producer %d is done\n", i)
		}()
	}
}

func (pr *ProducersConsumer) fileConsumer() {
	for i := 0; i < pr.itemQuantity; {
		select {
		case operationResult := <-pr.sharedChannel:
			pr.result += operationResult
			i++
		}
	}
	fmt.Printf("File consumer is done. Result is %.2f\n", pr.result)
}

func (pr *ProducersConsumer) Run() {
	pr.fileProducers()
	pr.fileConsumer()
	pr.wg.Wait()
}

func NewProducersConsumer(quantity int, fetcher fetcher.FileFetcher) *ProducersConsumer {
	return &ProducersConsumer{
		itemQuantity:  quantity,
		sharedChannel: make(chan float64, quantity),
		fetcher:       &fetcher,
		result:        0.0,
	}
}
