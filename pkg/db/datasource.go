package db

type Datasource struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Port int    `yaml:"port"`
	Name string `yaml:"name"`
}
