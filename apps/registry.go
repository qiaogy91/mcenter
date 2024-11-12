package apps

import (
	_ "github.com/qiaogy91/mcenter/apps/endpoint/impl"
	_ "github.com/qiaogy91/mcenter/apps/role/impl"
	_ "github.com/qiaogy91/mcenter/apps/token/api"
	_ "github.com/qiaogy91/mcenter/apps/token/impl"
	_ "github.com/qiaogy91/mcenter/apps/token/provider"
	_ "github.com/qiaogy91/mcenter/apps/user/impl"

	_ "github.com/qiaogy91/mcenter/apps/token/provider/account/impl"
	_ "github.com/qiaogy91/mcenter/apps/token/provider/feishu/impl"
)
