package apps

// Controller
import (
	_ "github.com/qiaogy91/mcenter/apps/endpoint/impl"
	_ "github.com/qiaogy91/mcenter/apps/role/impl"
	_ "github.com/qiaogy91/mcenter/apps/token/impl"
	_ "github.com/qiaogy91/mcenter/apps/user/impl"

	// handler
	_ "github.com/qiaogy91/mcenter/apps/endpoint/api"
	_ "github.com/qiaogy91/mcenter/apps/role/api"
	_ "github.com/qiaogy91/mcenter/apps/token/api"
	_ "github.com/qiaogy91/mcenter/apps/user/api"

	// provider
	_ "github.com/qiaogy91/mcenter/apps/token/provider"
	_ "github.com/qiaogy91/mcenter/apps/token/provider/account"
	_ "github.com/qiaogy91/mcenter/apps/token/provider/feishu"
)

// Handler

// Provider
