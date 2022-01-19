package vault_mock

func hello() string {
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