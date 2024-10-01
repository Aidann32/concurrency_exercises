package files

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
)

func createFolder(folderName string) (err error) {
	err = os.Mkdir(folderName, 0777)
	return err
}

func createFiles(fileNumber int) (err error) {
	if err = createFolder("files"); err != nil {
		return err
	}
	if err = os.Chdir("files"); err != nil {
		return err
	}
	actions := []string{"*", "/", "+", "-"}
	for i := 0; i < fileNumber; i++ {
		firstNum, secondNum := rand.Float64(), rand.Float64()
		file, err := os.Create(fmt.Sprintf("in_%d.dat", i))
		if err != nil {
			return err
		}
		if _, err = file.WriteString(fmt.Sprintf("%s\n%.2f %.2f", actions[rand.Intn(len(actions))], firstNum, secondNum)); err != nil {
			file.Close()
			return err
		}
		file.Close()
	}
	return nil
}

func getResultFromFile(fileName string) (result float64) {
	data, _ := os.ReadFile(fileName)
	input := strings.Split(string(data), "\n")
	fmt.Println(input)
	operation := input[0]
	operands := strings.Split(input[1], " ")
	// fmt.Println(operands)
	floatOperand1, _ := strconv.ParseFloat(operands[0], 64)
	floatOperand2, _ := strconv.ParseFloat(operands[1], 64)
	switch operation {
	case "+":
		return floatOperand1 + floatOperand2
	case "-":
		return floatOperand1 - floatOperand2
	case "*":
		return floatOperand1 * floatOperand2
	case "/":
		return floatOperand1 / floatOperand2
	default:
		fmt.Println("Default used")
		return 0
	}
}

func fileProducers(fileNumber int, sharedChannel chan float64, wg *sync.WaitGroup) {
	for i := 0; i < fileNumber; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			sharedChannel <- getResultFromFile(fmt.Sprintf("in_%d.dat", i))
			fmt.Printf("Goroutine producer %d is done\n", i)
		}()
	}
}

func fileConsumer(sharedChannel chan float64, counter *int, result *float64, fileNumber int) {
	for i := 0; i < fileNumber; {
		select {
		case operationResult := <-sharedChannel:
			*result += operationResult
			*counter++
			i++
		}
	}
	fmt.Printf("File consumer is done. Result is %.2f\n", *result)
}

func Run(fileNumber int) {
	if _, err := os.Stat("files/"); os.IsNotExist(err) {
		if err := createFiles(fileNumber); err != nil {
			fmt.Printf("Create directory and files error: %s\n", err)
			return
		}
	}
	_ = os.Chdir("files")
	counter := 0
	result := 0.0

	sharedChannel := make(chan float64, fileNumber)
	var wg sync.WaitGroup

	// Running producers
	fileProducers(fileNumber, sharedChannel, &wg)

	// Running consumer
	fileConsumer(sharedChannel, &counter, &result, fileNumber)
	wg.Wait()

	fmt.Printf("%.2f %d\n", result, counter)
}
