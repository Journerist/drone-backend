package command

import "log"

type StartMotorCommand struct {

}

func (c *StartMotorCommand) Execute () {
	log.Println("Start motor...")
}

func (c *StartMotorCommand) GetId () string {
	return "START_MOTOR"
}