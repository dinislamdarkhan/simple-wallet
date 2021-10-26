package grace

type Connections interface {
	Close()
}

type Service interface {
	Name() string
	Stop()
}

type service struct {
	name string
	stop func()
}

func NewService(name string, stop func()) Service {
	return service{
		name: name,
		stop: stop,
	}
}

func (s service) Stop() {
	s.stop()
}

func (s service) Name() string {
	return s.name
}
