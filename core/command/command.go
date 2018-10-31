package command

type Command interface {
	Execute()
	GetId() string
}
