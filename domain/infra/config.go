package infra

type Config interface {
	JwtAccessSecret() string
	JwtRefreshSecret() string
	JwtAccessExpireMinutes() string
	JwtRefreshExpireHours() string
	JwtPrivatePemPath() string
	JwtPublicPemPath() string
	FullHttpHost() string
	HttpPort() string
	DbHost() string
	DbPort() string
	DbUser() string
	DbPassword() string
	DbName() string
	DbSslMode() string
	AdminLogin() string
	AdminPassword() string
}
