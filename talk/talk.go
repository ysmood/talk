package talk

type Talk interface {
	Register(t interface{})
	Connect(t interface{}) interface{}
}

type Service struct {
}
