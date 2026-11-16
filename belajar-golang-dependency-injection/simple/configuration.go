package simple

type Configuration struct {
	Name string
}

type Aplication struct{
	*Configuration
}

func NewAplication() *Aplication {
	return &Aplication{
		Configuration: &Configuration{
			Name: "Belajar Golang Dependency Injection",
		},
	}
}