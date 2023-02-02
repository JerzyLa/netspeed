package speedtester

import (
	"context"
	"github.com/stretchr/testify/assert"
	"netspeed/pkg/netspeed"
	"testing"
)

func TestRunSpeedTestWithMocks(t *testing.T) {
	srv := New()
	srv.mapProviders = map[Provider]speedTest{
		Fast: func(_ context.Context) (netspeed.Result, error) {
			return netspeed.Result{Download: 1, Upload: 1}, nil
		},
		Speedtest: func(_ context.Context) (netspeed.Result, error) {
			return netspeed.Result{Download: 2, Upload: 2}, nil
		},
	}

	res, err := srv.RunSpeedTest(context.Background(), Fast)
	assert.NoError(t, err)
	assert.Equal(t, netspeed.Result{Download: 1, Upload: 1}, res)

	res, err = srv.RunSpeedTest(context.Background(), Speedtest)
	assert.NoError(t, err)
	assert.Equal(t, netspeed.Result{Download: 2, Upload: 2}, res)
}

func TestRunSpeedTest(t *testing.T) {
	srv := New()
	res, err := srv.RunSpeedTest(context.Background(), Fast)
	t.Logf("fast: download speed: %f, upload speed: %f", res.Download, res.Upload)
	assert.NoError(t, err)

	res, err = srv.RunSpeedTest(context.Background(), Speedtest)
	t.Logf("speedtest: download speed: %f, upload speed: %f", res.Download, res.Upload)
	assert.NoError(t, err)
}

func BenchmarkFastProvider(b *testing.B) {
	srv := New()
	for i := 0; i < b.N; i++ {
		_, err := srv.RunSpeedTest(context.Background(), Fast)
		assert.NoError(b, err)
	}
}

func BenchmarkSpeedtestProvider(b *testing.B) {
	srv := New()
	for i := 0; i < b.N; i++ {
		_, err := srv.RunSpeedTest(context.Background(), Speedtest)
		assert.NoError(b, err)
	}
}
