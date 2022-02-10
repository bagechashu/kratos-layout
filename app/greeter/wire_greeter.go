package greeter

import (
	"github.com/bagechashu/kratos-layout/app/greeter/biz"
	"github.com/bagechashu/kratos-layout/app/greeter/data"
	"github.com/bagechashu/kratos-layout/app/greeter/service"
	"github.com/google/wire"
)

// ProviderSet is data providers.
// ProviderSet is biz providers.
// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	data.NewData,
	data.NewGreeterRepo,
	biz.NewGreeterUsecase,
	service.NewGreeterService,
)
