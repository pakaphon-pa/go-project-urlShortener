package service

import "urlShortener/storage"

type GeneratorService struct {
	storage storage.RedisInterface
}

type GeneratorServiceInterface interface {
	GenerateShortLink(initialLink string) string
}

func (g *GeneratorService) GenerateShortLink(initialLink string) string {
	return "test"
}

func NewGeneratorService(storage storage.RedisInterface) GeneratorServiceInterface {
	return &GeneratorService{
		storage: storage,
	}
}
