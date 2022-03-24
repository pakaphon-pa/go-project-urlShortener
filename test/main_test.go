package test

import (
	"context"
	"fmt"
	"log"
	"urlShortener/storage"

	"github.com/go-redis/redis"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestSuite struct {
	suite.Suite
	db *redis.Client

	storage storage.RedisInterface
}

func (s *TestSuite) SetupSuite() {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "redis:6",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("* Ready to accept connections"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer container.Terminate(ctx)

	if err != nil {
		log.Fatal("start:", err)
	}

	mappedPort, err := container.MappedPort(ctx, "6379")
	if err != nil {
		log.Fatal("mappedPort:", err)
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		log.Fatal("hostIp:", err)
	}

	uri := fmt.Sprintf("redis://%s:%s", hostIP, mappedPort.Port())
	options, err := redis.ParseURL(uri)
	if err != nil {
		log.Fatal("reis:", err)
	}
	client := redis.NewClient(options)

	s.db = client
}

func (s *TestSuite) SetupTest() {
	s.storage = storage.NewRedisService(s.db)

}
