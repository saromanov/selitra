package v1

// Config defines configuration for the app
type Config struct {
	Address    string `yaml:"Address" json:"address"`
	DBName     string `yaml:"Dbname" json:"db_name"`
	DBPassword string `yaml:"Dbpassword" json:"-"`
	DBUser     string `yaml:"Dbuser" json:"-"`
}
