// department/department.go
package department

// Дефолтные температурные ограничения согласно условию задачи
const (
	DefaultMinTemp = 15
	DefaultMaxTemp = 30
)

// Department хранит текущие температурные ограничения отдела
type Department struct {
	MinTemp int // Актуальная нижняя граница
	MaxTemp int // Актуальная верхняя граница
}

// New создает отдел с начальными температурными настройками
func New() *Department {
	return &Department{
		MinTemp: DefaultMinTemp,
		MaxTemp: DefaultMaxTemp,
	}
}

// Update обновляет температурные границы согласно условию сотрудника
func (d *Department) Update(op string, temp int) {
	switch op {
	case ">=":
		// Обновляем нижнюю границу только если новое значение больше текущей
		if temp > d.MinTemp {
			d.MinTemp = temp
		}
	case "<=":
		// Обновляем верхнюю границу только если новое значение меньше текущей
		if temp < d.MaxTemp {
			d.MaxTemp = temp
		}
	}
}

// OptimalTemperature вычисляет оптимальную температуру
func (d *Department) OptimalTemperature() int {
	// Если границы пересеклись - решения нет
	if d.MinTemp > d.MaxTemp {
		return -1
	}
	// Возвращаем нижнюю границу
	return d.MinTemp
}
