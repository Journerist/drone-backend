package drone

import "log"

type Drone struct {
	motors []*Motor
}

func (d *Drone) Init() {
	d.motors = []*Motor{new(Motor), new(Motor), new(Motor), new(Motor)}
}

func (d *Drone) start() {
	log.Print("Calibrate motors")
	for _, motor := range d.motors {
		motor.calibrate()
	}

	log.Print("Start motors")
	for _, motor := range d.motors {
		motor.start()
	}
}

func (d *Drone) stop() {
	log.Print("Stop motors")
	for _, motor := range d.motors {
		motor.stop()
	}
}


