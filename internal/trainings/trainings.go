package trainings

import (
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
//...

// создайте метод ActionInfo()
