package common

const (
	KeyringService            = "kubex"
	DefaultGoBEConfigPath     = "$HOME/.kubex/gobe/config/config.json"
	DefaultGoBEKeyPath        = "$HOME/.kubex/gobe/gobe-key.pem"
	DefaultGoBECertPath       = "$HOME/.kubex/gobe/gobe-cert.pem"
	DefaultGodoBaseConfigPath = "$HOME/.kubex/gdbase/config/config.json"
)

const (
	DefaultRateLimitLimit  = 100
	DefaultRateLimitBurst  = 100
	DefaultRequestWindow   = 1 * 60 * 1000 // 1 minute
	DefaultRateLimitJitter = 0.1
)

const (
	DefaultMaxRetries = 3
	DefaultRetryDelay = 1 * 1000 // 1 second
)

const (
	DefaultMaxIdleConns          = 100
	DefaultMaxIdleConnsPerHost   = 100
	DefaultIdleConnTimeout       = 90 * 1000 // 90 seconds
	DefaultTLSHandshakeTimeout   = 10 * 1000 // 10 seconds
	DefaultExpectContinueTimeout = 1 * 1000  // 1 second
	DefaultResponseHeaderTimeout = 5 * 1000  // 5 seconds
	DefaultTimeout               = 30 * 1000 // 30 seconds
	DefaultKeepAlive             = 30 * 1000 // 30 seconds
	DefaultMaxConnsPerHost       = 100
)
