package models

// ArchiConfig is about the whole CLI config
type ArchiConfig struct {
	DefaultArchitecture string `json:"default_architecture"`
	DefaultLanguage     string `json:"default_language"`
	TemplatesPath       string `json:"templates_path"`
}
