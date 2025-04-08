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
		return fmt.Errorf("wrong data amount")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("steps parsing error: %v", err)
	}
	t.Steps = steps

	TrainingType := parts[1]
	if TrainingType != "Бег" && TrainingType != "Ходьба" {
		return fmt.Errorf("unknown training type %s", TrainingType)
	}
	t.TrainingType = TrainingType

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("duration parsing error: %v", err)
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()

func (t Training) ActionInfo() (string, error) {

	if t.Duration < 0 {
		return "", fmt.Errorf("duration less 0 %v", t.Duration)
	}

	distance := spentenergy.Distance(t.Steps)

	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	durationInHours := t.Duration.Hours()

	var calories float64

	switch t.TrainingType {
	case "Бег":

		calories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)

	case "Ходьба":

		calories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	default:

		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)

	}

	toBePrinted := fmt.Sprintf(
		"\nТип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f", t.TrainingType, durationInHours, distance, meanSpeed, calories)

	return toBePrinted, nil

}
