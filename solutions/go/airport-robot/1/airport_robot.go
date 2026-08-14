package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitor string) string
}

type Italian struct{}

func (robot Italian) LanguageName() string {
	return "Italian"
}

func (robot Italian) Greet(visitor string) string {
	return fmt.Sprintf("Ciao %s!", visitor)
}

type Portuguese struct{}

func (robot Portuguese) LanguageName() string {
	return "Portuguese"
}

func (robot Portuguese) Greet(visitor string) string {
	return fmt.Sprintf("Olá %s!", visitor)
}

func SayHello(visitor string, robot Greeter) string {
	return fmt.Sprintf(
		"I can speak %s: %s",
		robot.LanguageName(),
		robot.Greet(visitor),
	)
}
