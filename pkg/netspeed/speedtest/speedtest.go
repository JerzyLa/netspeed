package speedtest

import (
	"context"
	"github.com/showwin/speedtest-go/speedtest"
	"netspeed/pkg/netspeed"
)

// RunSpeedTest runs the speed test using `speedtest.net` and returns download and upload speeds
func RunSpeedTest(ctx context.Context) (netspeed.Result, error) {
	user, err := speedtest.FetchUserInfoContext(ctx)
	if err != nil {
		return netspeed.Result{}, err
	}

	serverList, err := speedtest.FetchServerListContext(ctx, user)
	if err != nil {
		return netspeed.Result{}, err
	}

	targets, err := serverList.FindServer(nil)
	if err != nil {
		return netspeed.Result{}, err
	}

	server := targets[0]
	err = server.PingTestContext(ctx)
	if err != nil {
		return netspeed.Result{}, err
	}

	err = server.DownloadTestContext(ctx, false)
	if err != nil {
		return netspeed.Result{}, err
	}

	err = server.UploadTestContext(ctx, false)
	if err != nil {
		return netspeed.Result{}, err
	}

	return netspeed.Result{
		Download: server.DLSpeed,
		Upload:   server.ULSpeed,
	}, nil
}
