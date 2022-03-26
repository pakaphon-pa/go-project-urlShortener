package service

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"urlShortener/storage"

	"github.com/itchyny/base58-go"
)

type GeneratorService struct {
	storage storage.RedisInterface
}

type GeneratorServiceInterface interface {
	GenerateShortLink(initialLink string) string
	Get(shortUrl string) string
}

func (g *GeneratorService) GenerateShortLink(initialLink string) string {
	urlHashBytes := sha2560f(initialLink)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	g.storage.Save(finalString[:8], initialLink)

	return finalString[:8]
}

func (g *GeneratorService) Get(shortUrl string) string {
	return g.storage.Read(shortUrl)
}

func sha2560f(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

func NewGeneratorService(storage storage.RedisInterface) GeneratorServiceInterface {
	return &GeneratorService{
		storage: storage,
	}
}
