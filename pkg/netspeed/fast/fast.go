package fast

import (
	"context"
	"github.com/gesquive/fast-cli/fast"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"net/url"
	"netspeed/pkg/netspeed"
	"strings"
	"time"
)

var client = http.Client{}

const (
	workload      = 8
	payloadSizeMB = 25.0 // download and upload payload is by default 25MB
)

// RunSpeedTest runs the speed test using `fast.com` and returns download and upload speeds
func RunSpeedTest(_ context.Context) (netspeed.Result, error) {
	urls := fast.GetDlUrls(1)
	if len(urls) == 0 {
		urls = append(urls, fast.GetDefaultURL())
	}
	fastUrl := urls[0]

	downloadSpeed, err := measureNetworkSpeed(download, fastUrl)
	if err != nil {
		return netspeed.Result{}, err
	}

	uploadSpeed, err := measureNetworkSpeed(upload, fastUrl)
	if err != nil {
		return netspeed.Result{}, err
	}

	return netspeed.Result{
		Download: downloadSpeed,
		Upload:   uploadSpeed,
	}, nil
}

// measureNetworkSpeed sends multiple requests to the target website, calculates the time it took and returns network speed
func measureNetworkSpeed(operation func(url string) error, url string) (float64, error) {
	eg := errgroup.Group{}
	startTime := time.Now()

	for i := 0; i < workload; i++ {
		eg.Go(func() error {
			return operation(url)
		})
	}
	err := eg.Wait()
	if err != nil {
		return 0, err
	}

	endTime := time.Now()

	return calculateSpeed(startTime, endTime), nil
}

// calculateSpeed calculates speed based on payload size and how long it took to send/receive it
func calculateSpeed(sTime time.Time, fTime time.Time) float64 {
	return payloadSizeMB * 8 * float64(workload) / fTime.Sub(sTime).Seconds()
}

func download(url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

func upload(uri string) error {
	v := url.Values{}

	v.Add("content", createUploadPayload())

	resp, err := client.PostForm(uri, v)
	if err != nil {
		return err
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return resp.Body.Close()
}

func createUploadPayload() string {
	return strings.Repeat("0123456789", payloadSizeMB*1024*1024/10)
}
