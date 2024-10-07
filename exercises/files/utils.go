package files

import (
	"fmt"
	"math/rand"
	"os"
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
