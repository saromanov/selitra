package v1

// Config defines configuration for the app
type Config struct {
	Address    string `yaml:"address"`
	DBName     string `yaml:"dbname`
	DBPassword string `yaml:dbpassword"`
}
