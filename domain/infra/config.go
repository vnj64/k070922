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
}
