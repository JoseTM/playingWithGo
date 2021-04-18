package main

type englishBot struct{}
type spanishBot struct{}

// si no hacemos uso de la variable del receiver en la función podemos omitirla
func (englishBot) getGreeting() string {
	return "hello there"
}

func (spanishBot) getGreeting() string {

	return "hola a todos"
}
