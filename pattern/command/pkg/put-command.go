package pkg

type PutCommand struct {
	rent Rent
}

func NewPutCommand(rent Rent) *PutCommand {
	return &PutCommand{rent}
}

func (pc *PutCommand) execute() {
	pc.rent.put()
}