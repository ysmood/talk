package types

type Profile struct {
	Name string
}

type Manager interface {
	Get() Profile
}
