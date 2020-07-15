package config

import "os"

// SecurityConfig structure
type SecurityConfig struct {
	AuthPrivateKey string `env:"AUTH_PRIVATE_KEY"`
	TokenLifeTime  int    `env:"TOKEN_LIFE_TIME" envDefault:"15"`
}

// ProvideSecurityConfig provide the security config
func ProvideSecurityConfig(config Config) SecurityConfig {
	securityConfig := config.SecurityConfig

	if os.Getenv("AUTH_PRIVATE_KEY") == "" {
		// Find out cleaner to set this kind of default value this way
		securityConfig.AuthPrivateKey = `----BEGIN RSA PRIVATE KEY-----
		MIIEpAIBAAKCAQEAx8RG2RjLntWQLyamnWWGsqrfTYsmXMkccIj56iUHIpH47VAe
		fJ/D/u9yb9IFypSKM4WLqfUcWS1qNBjvrPMynXP7bLCFSizIL4UNWskFMdOdTAhk
		Hcrpl5eSlOEVm370rH4wVM8mJMbt5ttRsYFIH2ORD8wduJBJg+VPNGGcKkZc5EBS
		dP+zB+eHBtS/mdXOm4yh3HH294j0/C78p2KiJBmlSP11TzDHtutQv4Of6zcNGZkj
		zNMB9/EzSy9xAcExBRUtLwSo9HVM9vSH5PBX9ZfqfEMqPQ7GFdL61QcICejttd9o
		ZemqcvTYfHR2+5kRF64O/gJ4+b2dceNLryrRIwIDAQABAoIBAGwp5WaVG2NkJU6y
		cdKVb0p8ZK5a4qmopwJz2EWkSOm5+fucvvUhmIEH64GKSWJ37SCadZXPxsquYWvV
		+Kg/gi9o5kDijzkraDT6qNF4OJzTBL9vZ64N+U8AXfjIqudn5VkkdkzlYLTXFWdD
		Z0diRcfhJGFWVt1hGr0Uwz8Mfv8P50EV7b2xh+QKtCgz2XOPIiTPCITd/hq0Tm2C
		CP8mYHdeDvCXF5pQz2eWXhl4Ym025tHQhILpGsQZfs2fQ0uaJQB4vvmfJY+dhnAx
		qhORUz3xQ9lCUy86SHLDsL7S8/t2114V3PiOs9CEQuA8DSi3JND3mM40Rhor9K4e
		9zRo+SECgYEA5Nb09r4szhWkmUvwbxxXYbxJf2LUJthpAbS5Zi6SrS9xKkN9EAy9
		/V9iVyZ10aGn1Osx7HtjvPfCQ40itVSxXM//VYVME6hEale0fWR1xivr4/Rn0sCx
		ROTsOTLK0Nj1Yy/AMbJt+yMqvu8jDQEDkwXJw830iJ8K+U1BA2rq0lECgYEA33n4
		MWBFjfADdKVI8TNZHW+KQhxNhJF4AAVdmo+kmCjoWph+NU2VlkmruK1qnIwqt3Gf
		riCpU1HVRT0HR5lBX1wfsafCAtAKB08TpGVfKA2GszzVeSpbVAlJXXz0LMGdM45N
		IeRZGpn1k+GwogV2AV6godyoCNJwzBS8cj52ezMCgYEA2J1xfC9AkkOCb2mmykZ1
		oJpuqBqQ38RygZ7k8IidZzpFUNBZjdvB7BliKUTKugjfqXiI5WVXYATW6sKqXCMz
		ULvC4t0uyly5hkBHJXr20QO6YylzLdgg9SoNQYeG3HLCtYT51RElclgzmlEbSK/x
		BjawapyGIN0LsjkilW2bokECgYAcDpAegtu4qj2E04uI3M1/aJWPH0aMMkdCmWnb
		EdW0VM53MVYHEyqIp88rbHpKFT73sQg/MFKR52BDRRo43wDaHBLOlyVAfSYFLORm
		a2O6/5e9wRF8GNlGf7P6TMnLJLO72U3Nx4F4WYP/gQJQFSdvnJXGPJHMKb+Hn8LZ
		BQaTCQKBgQCOSpjGM882KHjC/Mr3e8BKii1BAR1aEI0lDyAHD90ozi9sA6xAh7tM
		wccmqrtok3rcgaxkEa44lwIg/QfODmOiZ/sQbVj3srgvUEs85or1cX82cwbBXBxK
		Jg31gSVi79LcB/+W7BezGtex7Ru6rTja6DJalM0HzD5s3OsLWV2h+Q==
		-----END RSA PRIVATE KEY-----`
	}

	return securityConfig
}
