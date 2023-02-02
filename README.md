# Netspeed

Contains the package for testing upload and download speeds using `https://www.speedtest.net` and `https://fast.com`

## Usage

```
ctx := context.Background()
srv := speedtester.New()

// Run speed test on fast.com
res, _ := srv.RunSpeedTest(ctx, speedtester.Fast)
fmt.Println(res.Download, res.Upload)

// Run speed test on speedtest.net
res, _ = srv.RunSpeedTest(ctx, speedtester.Speedtest)
fmt.Println(res.Download, res.Upload)
```
