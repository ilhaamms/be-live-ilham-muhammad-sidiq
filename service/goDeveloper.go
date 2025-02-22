package service

type GoDeveloperService interface {
	GoDeveloper() (string, error)
}

type GoDeveloperServices struct {
	// UserRepository repository.UserRepository
}

func NewGoDeveloperService() GoDeveloperService {
	return &GoDeveloperServices{}
}

func (service *GoDeveloperServices) GoDeveloper() (string, error) {
	return "Hello Go developers", nil
}
