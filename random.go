package vault_mock

func Hello() string {
	return "it works!"
}

type Please struct {
	name string
	success bool
}

func (p Please) Print() string {
	return "hello?"
}

var S = "greetings"