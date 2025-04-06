package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("Неверное количество данных")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %v", err)
	}
	t.Steps = steps

	TrainingType := parts[1]
	if TrainingType != "Бег" && TrainingType != "Ходьба" {
		return fmt.Errorf("неизвестный тип активности: %s", TrainingType)
	}
	t.TrainingType = TrainingType

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга продолжительности: %v", err)
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()

func (t Training) ActionInfo() (string, error) {

	if t.Duration < 0 {
		return "", fmt.Errorf("Продолжительность меньше нуля и равна %v", t.Duration)
	}

	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}

	distance := spentenergy.Distance(t.Steps)

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	durationInHours := t.Duration.Hours()

	var calories float64

	if t.TrainingType == "Бег" {

		calories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)

	} else {
		calories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}

	toBePrinted := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f", t.TrainingType, durationInHours, distance, meanSpeed, calories)

	return toBePrinted, nil

}
