package interfaces

import (
	"github.com/journerist/drone-backend/core/command"
	"github.com/kataras/iris"
)

type CommandController struct {
	commandRepository command.CommandRepository
	app   *iris.Application
}

func (c *CommandController) Init(app *iris.Application, repository command.CommandRepository) {
	c.app = app
	c.commandRepository = repository
}

func (c *CommandController) Start() {
	c.app.Get("/command", c.readAll)
	c.app.Get("/command/{id:string}/execute", c.executeCommand)

	c.app.Run(iris.Addr(":8080"))
}

func (c *CommandController) readAll(ctx iris.Context) {
	commands := make([]CommandRO, 1)
	commands[0] = CommandRO{
		"START_MOTOR",
		"Start motor",
		"Start the motor with low speed.",
	}

	ctx.JSON(commands)
}

func (c *CommandController) executeCommand(ctx iris.Context) {
	commandId := ctx.Params().Get("id")

	commandToExecute,_ := c.commandRepository.Find(commandId)
	commandToExecute.Execute()

	ctx.JSON(iris.Map{
		"ok": "1",
	})
}

type CommandRO struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}