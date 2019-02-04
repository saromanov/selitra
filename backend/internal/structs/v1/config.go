package v1

// Config defines configuration for the app
type Config struct {
	Address    string `yaml:"Address"`
	DBName     string `yaml:"Dbname"`
	DBPassword string `yaml:"Dbpassword"`
	DBUser     string `yaml:"Dbuser"`
}
