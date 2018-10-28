package core

import "log"

type Drone struct {
	 modelNumber string
	 id string
}

func (d Drone) Start() {
	log.Print("Starting...")


}