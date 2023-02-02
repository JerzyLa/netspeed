package speedtester

import (
	"context"
	"fmt"
	"netspeed/pkg/netspeed"
	"netspeed/pkg/netspeed/fast"
	"netspeed/pkg/netspeed/speedtest"
)

type speedTest func(ctx context.Context) (netspeed.Result, error)

// Service used to run speed tests on different providers
type Service struct {
	mapProviders map[Provider]speedTest
}

// New creates a new service
func New() Service {
	m := map[Provider]speedTest{
		Speedtest: speedtest.RunSpeedTest,
		Fast:      fast.RunSpeedTest,
	}
	return Service{mapProviders: m}
}

// RunSpeedTest executes speed test on specified provider
func (s Service) RunSpeedTest(ctx context.Context, provider Provider) (netspeed.Result, error) {
	f, ok := s.mapProviders[provider]
	if !ok {
		return netspeed.Result{}, fmt.Errorf("unknown provider: %s", provider)
	}
	return f(ctx)
}
