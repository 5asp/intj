package configfx

import (
	"github.com/knadh/koanf/v2"
	"go.uber.org/fx"
)

// ProvideConfig to fx
func ProvideConfig() fx.Option {
	return fx.Provide(func() *koanf.Koanf {
		conf := koanf.New(".")
		// conf.Load(file.Provider("configs/"+configFilePath+"/app.yaml"), yaml.Parser())
		return conf
	})
}
