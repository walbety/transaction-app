package service

type Service struct{}

func New() Service {
	return Service{}
}

func (svc *Service) welcome() {
	return
}
