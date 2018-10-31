package command

import "github.com/kataras/iris/core/errors"

type CommandRepository struct {

}

func (repository *CommandRepository) Find(id string) (Command, error) {

	startMotorCommand := new(StartMotorCommand)

	if(startMotorCommand.GetId() == id) {
		return startMotorCommand, nil
	}

	return nil, errors.New("No command found with id '" + id + "'")
}