package service

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/JanitSri/equity-pulse/mocks"
	"github.com/JanitSri/equity-pulse/model"
	"github.com/JanitSri/equity-pulse/net"
	rd "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/testcontainers/testcontainers-go/wait"
)

type EquityIntegrationTestSuite struct {
	suite.Suite
	ctx                   context.Context
	rContainer            *redis.RedisContainer
	rClient               *rd.Client
	stockDataProviderMock *mocks.StockDataProvider
	equityService         *EquityService
	equityReq             *model.EquityRequest
}

func TestEquityIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(EquityIntegrationTestSuite))
}

func (s *EquityIntegrationTestSuite) SetupSuite() {
	s.ctx = context.Background()

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "redis:latest",
			ExposedPorts: []string{"6379/tcp"},
			WaitingFor:   wait.ForLog("Ready to accept connections"),
		},
	}

	redisContainer, err := redis.RunContainer(s.ctx, testcontainers.CustomizeRequest(req))
	s.Require().NoError(err)

	rdConnStr, err := redisContainer.ConnectionString(s.ctx)
	s.Require().NoError(err)

	rdConnOptions, err := rd.ParseURL(rdConnStr)
	s.Require().NoError(err)

	s.rContainer = redisContainer
	s.rClient = rd.NewClient(rdConnOptions)

	pong, err := s.rClient.Ping(ctx).Result()
	s.Require().NoError(err)
	s.Equal("PONG", pong)
}

func (s *EquityIntegrationTestSuite) TearDownSuite() {
	err := s.rContainer.Terminate(s.ctx)
	s.Require().NoError(err)
}

func (s *EquityIntegrationTestSuite) BeforeTest(suiteName, testName string) {
	s.stockDataProviderMock = &mocks.StockDataProvider{}
	cs := NewCacheService(net.NewRedisCache(s.rClient))

	s.equityService = NewEquityService(s.stockDataProviderMock, cs)
	s.equityReq = model.NewEquityRequestBuilder().Ticker("AAPL").Build()
}

func (s *EquityIntegrationTestSuite) AfterTest(suiteName, testName string) {
	s.rClient.FlushAll(s.ctx)
}

func (s *EquityIntegrationTestSuite) TestStockTickerInfo() {
	mockData := &model.TickerInfoBuilder{
		Symbol:    "AAPL",
		QuoteType: "stock",
		Exchange:  "nasdaq",
		ShortName: "Apple Inc.",
		LongName:  "Apple Inc.",
	}
	t := s.equityReq.Ticker()

	s.stockDataProviderMock.On("RetrieveStockTickerInfo", t).Return(mockData, nil)
	d, err := s.equityService.StockTickerInfo(s.equityReq)

	s.NoError(err)
	s.Equal(mockData.Build(), d)

	key := fmt.Sprintf(stockTickerInfoKey, t)
	val, err := s.rClient.Get(ctx, key).Result()
	s.NoError(err)
	s.NotEmpty(val)
}

func (s *EquityIntegrationTestSuite) TestStockTickerInfoRedis() {
	t := s.equityReq.Ticker()
	key := fmt.Sprintf(stockTickerInfoKey, t)
	mockData := &model.TickerInfoBuilder{
		Symbol:    "AAPL",
		QuoteType: "stock",
		Exchange:  "nasdaq",
		ShortName: "Apple Inc.",
		LongName:  "Apple Inc.",
	}
	jsonBytes, err := json.Marshal(mockData)
	s.Require().NoError(err)

	err = s.rClient.Set(s.ctx, key, string(jsonBytes), time.Minute).Err()
	s.Require().NoError(err)

	d, err := s.equityService.StockTickerInfo(s.equityReq)

	s.NoError(err)
	s.Equal(mockData.Build(), d)
	s.stockDataProviderMock.AssertNotCalled(s.T(), "RetrieveStockTickerInfo")
}

func (s *EquityIntegrationTestSuite) TestStockTickerInfoError() {
	t := s.equityReq.Ticker()

	e := fmt.Errorf("error")
	s.stockDataProviderMock.On("RetrieveStockTickerInfo", t).Return(nil, e)
	d, err := s.equityService.StockTickerInfo(s.equityReq)

	s.Equal(e, err)
	s.Empty(nil, d)

	key := fmt.Sprintf(stockTickerInfoKey, t)
	val, err := s.rClient.Get(ctx, key).Result()
	s.Equal(rd.Nil, err)
	s.Empty(val)
}
