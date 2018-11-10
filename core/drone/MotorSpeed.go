package drone

type MotorSpeed struct {
	value 	uint	// between [0,100]
}

const increment = 10
const initValue = 10
const maxRange = 100
const minRange = 0


func (ms *MotorSpeed) Init() {
	ms.value = initValue
}
func (ms *MotorSpeed) increase() {
	if ms.value +increment <= maxRange {
		ms.value += increment
	}
}
func (ms *MotorSpeed) decrease() {
	if ms.value -increment <= minRange {
		ms.value -= increment
	}
}
