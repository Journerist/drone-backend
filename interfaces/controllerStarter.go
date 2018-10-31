package interfaces

type ControllerStarter struct {
	commandController *CommandController
}

func (c *ControllerStarter) Init(commandController *CommandController) {
	c.commandController = commandController
}

func (c *ControllerStarter) StartAllController() {
	// listen and serve on http://0.0.0.0:8080.
	c.commandController.Start()
}
