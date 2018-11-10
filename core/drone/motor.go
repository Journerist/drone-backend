package drone

type Motor struct {
	running bool
	speed *MotorSpeed
}

func (m *Motor) Init() {
	m.running = false
	m.speed = new(MotorSpeed)
}

func (m *Motor) start() {
	m.running = true
}

func (m *Motor) stop() {
	m.running = false
}

func (m *Motor) faster() {
	m.speed.increase()
}

func (m *Motor) slower() {
	m.speed.decrease()
}

func (m *Motor) calibrate() {

}