package config

// HTTP defines the available http configuration.
type HTTP struct {
	Addr      string `ocisConfig:"addr" env:"ACCOUNTS_HTTP_ADDR"`
	Namespace string
	Root      string `ocisConfig:"root" env:"ACCOUNTS_HTTP_ROOT"`
	CacheTTL  int    `ocisConfig:"cache_ttl" env:"ACCOUNTS_CACHE_TTL"`
	CORS      CORS   `ocisConfig:"cors"`
}

// CORS defines the available cors configuration.
type CORS struct {
	AllowedOrigins   []string `ocisConfig:"allowed_origins"`
	AllowedMethods   []string `ocisConfig:"allowed_methods"`
	AllowedHeaders   []string `ocisConfig:"allowed_headers"`
	AllowCredentials bool     `ocisConfig:"allowed_credentials"`
}