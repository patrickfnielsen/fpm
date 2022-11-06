package types

type Settings struct {
	FileDirectory string `yaml:"fileDirectory"`

	Matchers []Matcher `yaml:"matchers"`
}

type Matcher struct {
	Name     string
	Regex    Regex    `yaml:"regex"`
	External External `yaml:"external"`
}

type Regex struct {
	Pattern string `yaml:"pattern"`
	Message string `yaml:"message"`
}

type External struct {
	Binary     string `yaml:"binary"`
	Parameters string `yaml:"parameters"`
}
