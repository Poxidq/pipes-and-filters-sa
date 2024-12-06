package shared

type Service interface {
	Execute(*Process)
	SetNext(Service)
}

type Process struct {
	Message Message
}

type Message struct {
	User    string `json:"user"`
	Message string `json:"message"`
}
