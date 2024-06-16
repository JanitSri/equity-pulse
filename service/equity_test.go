package service

import (
	"context"
	"testing"

	"github.com/JanitSri/equity-pulse/net"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
)

type EquityIntegrationTestSuite struct {
	suite.Suite
	ctx                   context.Context
	rClient               *redis.Client
	stockDataProviderMock net.StockDataProvider
}

func TestEquityIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(EquityIntegrationTestSuite))
}

func (s *EquityIntegrationTestSuite) SetupSuite() {
	s.T().Log("Running before all tests")
}

func (s *EquityIntegrationTestSuite) TearDownSuite() {
	s.T().Log("Running after all tests")
}

func (s *EquityIntegrationTestSuite) TestTrue() {
	s.T().Log("True Tests")
	s.True(true)
}

// TODO: add integration tests
