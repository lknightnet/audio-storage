package service

type AudioStorage interface {
	GetAudioByID(id string) ([]byte, error)
}

type Services struct {
	AS AudioStorage
}

type Dependencies struct {
	APIurl string
}

func NewServices(deps *Dependencies) *Services {
	return &Services{AS: newAudioService(deps.APIurl)}
}
