package speedtest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunSpeedTest(t *testing.T) {
	res, err := RunSpeedTest(context.Background())
	t.Logf("speedtest: download speed: %f, upload speed: %f", res.Download, res.Upload)
	assert.NoError(t, err)
}
