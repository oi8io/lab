package configs

const (
	CheckEvictInterval          = 3000
	RenewInterval               = 3000
	SelfProtectThreshold        = 3000
	InstanceExpireDuration      = 3000
	ResetGuardNeedCountInterval = 3000
)

type GlobalConfig struct {
	HttpServer string
}

func NewGlobalConfig(httpServer string) *GlobalConfig {
	return &GlobalConfig{HttpServer: httpServer}
}

func LoadConfig(conf string) (*GlobalConfig, error) {
	return NewGlobalConfig("127.0.0.1:8866"), nil
}
