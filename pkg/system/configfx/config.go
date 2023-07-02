package configfx

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// ProvideConfig to fx
func ProvideConfig() *koanf.Koanf {
	conf := koanf.New(".")
	conf.Load(file.Provider("app.yaml"), yaml.Parser())
	return conf
}
