package appv2

var ENV struct {
	Server struct {
		Port string `env:"PORT" envDefault:"8080"`
	} `envPrefix:"SERVER_"`

	Database struct {
		Host     string `env:"HOST" envDefault:"localhost"`
		Port     string `env:"PORT" envDefault:"5432"`
		User     string `env:"USER" envDefault:"postgres"`
		Password string `env:"PASSWORD" envDefault:"postgres"`
		DBName   string `env:"DB_NAME" envDefault:"postgres"`
	} `envPrefix:"DATABASE_"`
}
