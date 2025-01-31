package service

type PostService struct {
	base *BaseService
}

func newPostService(base *BaseService) *PostService {
	return &PostService{
		base: base,
	}
}

func (s *PostService) Create() error {

	return nil
}
