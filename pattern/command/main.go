package command

import "patterns/command/pkg"

func Command() {
	scooter := &pkg.Scooter{}

	takeCommand := pkg.NewTakeCommand(scooter)
	putCommand := pkg.NewPutCommand(scooter)

	takeScooterBtn := pkg.NewButton(takeCommand)
	takeScooterBtn.Press()

	putScooterBtn := pkg.NewButton(putCommand)
	putScooterBtn.Press()
}