package pkg

type TakeCommand struct {
	rent Rent
}

func NewTakeCommand(rent Rent) *TakeCommand {
	return &TakeCommand{rent}
}

func (pc *TakeCommand) execute() {
	pc.rent.take()
}