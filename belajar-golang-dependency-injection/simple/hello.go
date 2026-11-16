package simple

type SayHello interface {
	SayHello(string) string
}

type HelloService struct {
	SayHello
}

type SayHelloImpl struct {

}

func (s SayHelloImpl) SayHello(name string) string {
	return "Hello, " + name
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewHelloService(SayHello SayHello) *HelloService {
	return &HelloService{
		SayHello: SayHello,
	}
}