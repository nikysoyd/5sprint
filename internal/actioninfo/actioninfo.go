package actioninfo

import (
	"fmt"
	"log"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {

	for i, data := range dataset {
		fmt.Printf("\nОбработка записи %d: %s\n", i+1, data)

		//training := data{}

		err := dp.Parse(data)
		if err != nil {
			log.Printf("parsing error: %v\n", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("formating error %v\n", err)
			continue
		}

		fmt.Println(info)
	}

}
