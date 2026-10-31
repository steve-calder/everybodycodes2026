package notes

// Config holds the settings from config.yaml.
type Config struct {
	Event   int    `yaml:"EVENT"`
	Seed    int    `yaml:"SEED"`
	Session string `yaml:"SESSION"`
}
