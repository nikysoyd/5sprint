package actioninfo

import (
	"fmt"
	//"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/daysteps"
	//"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/trainings"
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
			fmt.Errorf("ошибка парсинга %v", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Errorf("ошибка формирования %v", err)
			continue
		}

		fmt.Println(info)
	}

}
