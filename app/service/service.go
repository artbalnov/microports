package service

type Service interface {
	Attach(address string) error
}
