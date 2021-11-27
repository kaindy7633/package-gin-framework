package config

type Jwt struct {
	Secret                  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl                  int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"`                                                          // token有效期(秒)
	JwtBlacklistGracePeriod int64  `mapstructure:"jwt_blacklist_grace_period" json:"jwt_blacklist_grace_period" yaml:"jwt_blacklist_grace_period"` // 黑名单宽限时间(秒)
	RefreshGracePeriod      int64  `mapstructure:"refresh_race_period" json:"refresh_race_period" yarm:"refresh_race_period"`                      // token 自动刷新宽限时间(秒)
}
