package pkg

type Button struct {
	command Command
}

func NewButton(command Command) *Button {
	return &Button{command}
}

func (b *Button) Press() {
	b.command.execute()
}