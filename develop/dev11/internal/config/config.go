package config

import (
	"encoding/json"
	"os"
	"time"
)

const (
	defaultHTTPPort               = "5011"
	defaultHTTPRWTimeout          = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1

	defaultLimiterRPS   = 10
	defaultLimiterBurst = 20
	defaultLimiterTTL   = 10 * time.Minute
)

type (
	Config struct {
		HTTP    HTTPConfig
		Cache   CacheConfig
		Limiter LimiterConfig
	}

	HTTPConfig struct {
		Port               string         `mapstructure:"port"`
		Timeouts           TimeoutsConfig `mapstructure:"timeouts"`
		MaxHeaderMegabytes int            `mapstructure:"maxHeaderMegabytes"`
	}

	TimeoutsConfig struct {
		Read  time.Duration `mapstructure:"read"`
		Write time.Duration `mapstructure:"write"`
	}

	CacheConfig struct {
		TTL   time.Duration `mapstructure:"ttl"`
		Clean time.Duration `mapstructure:"clean"`
	}

	LimiterConfig struct {
		RPS   int
		Burst int
		TTL   time.Duration
	}
)

func (cfg *Config) Init(path string) error {
	*cfg = setupDefault()

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(jsonFile, &cfg); err != nil {
		return err
	}

	cfg.Limiter.TTL *= time.Minute
	cfg.HTTP.Timeouts.Read *= time.Second
	cfg.HTTP.Timeouts.Write *= time.Second
	cfg.Cache.TTL *= time.Minute
	cfg.Cache.Clean *= time.Minute

	return nil
}

func setupDefault() Config {
	timeoutsConf := TimeoutsConfig{
		Read:  defaultHTTPRWTimeout,
		Write: defaultHTTPRWTimeout,
	}

	httpConf := HTTPConfig{
		Port:               defaultHTTPPort,
		Timeouts:           timeoutsConf,
		MaxHeaderMegabytes: defaultHTTPMaxHeaderMegabytes,
	}

	limiterConf := LimiterConfig{
		RPS:   defaultLimiterRPS,
		Burst: defaultLimiterBurst,
		TTL:   defaultLimiterTTL,
	}

	return Config{
		HTTP:    httpConf,
		Limiter: limiterConf,
	}
}
