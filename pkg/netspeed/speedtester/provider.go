package speedtester

// Provider defines providers used by speed tester to process speed tests
type Provider string

const (
	// Fast defines `fast.com` speed test provider
	Fast Provider = "fast"
	// Speedtest defines `speedtest.net` speed test provider
	Speedtest Provider = "speedtest"
)
