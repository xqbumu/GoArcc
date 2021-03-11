package config

import "go.uber.org/fx"

//ConfigProviderFx: config globally in the pipeline stream
var ConfigProviderFx = fx.Options(
	fx.Provide(
		GetConfig,
	),
)
