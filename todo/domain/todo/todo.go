package todo

// Service ...
type Service struct{}

// ServiceInput ...
type ServiceInput struct{}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	return &Service{}, nil
}

// List ...
func (s Service) List() [5]int {
	arr := [5]int{1, 2, 3, 4, 5}
	return arr
}
