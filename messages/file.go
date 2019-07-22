package messages

type Hello struct{}

func NewHello() *Hello {
	return &Hello{}
}

func (Hello) Say() string {
	message := " Hello World !"
	return message
}
