package logcopter

// Config describes logcopter runtime configuration after all application,
// profile, and CLI sources have been merged by the caller.
type Config struct {
	Output      string            `yaml:"output" json:"output"`
	Format      string            `yaml:"format" json:"format"`
	Level       string            `yaml:"level" json:"level"`
	Caller      bool              `yaml:"caller" json:"caller"`
	Timestamp   bool              `yaml:"timestamp" json:"timestamp"`
	Areas       map[string]string `yaml:"areas" json:"areas"`
	StrictAreas bool              `yaml:"strict_areas" json:"strict_areas"`
}

// AreaWarning reports a configured area that does not currently match a known
// generated/registered logging area.
type AreaWarning struct {
	Area    string
	Message string
}
