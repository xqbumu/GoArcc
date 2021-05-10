package user_profile

import (
	"go.uber.org/fx"
)

// Module Providing client to the Global fx pipe line, so other
//module can use our client
var Module = fx.Options(
	fx.Provide(
		NewUserProfileServer,
	),
)
