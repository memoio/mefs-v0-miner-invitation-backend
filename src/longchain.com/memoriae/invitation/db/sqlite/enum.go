package sqlite

type SubscribeEnum struct {
	// 待确认
	Unconfirmed string
	// 已订阅
	Subscribed string
	// 已取消
	Canceled string
}
type RoleEnum struct {
	User string
	Keeper string
	Provider string
}
type EnumStruct struct {
	Subscribe SubscribeEnum
	Role RoleEnum
}

var Enum = EnumStruct{
	Subscribe:SubscribeEnum{
		"unconfirmed",
		"subscribed",
		"canceled",
	},
	Role:RoleEnum{
		"user",
		"keeper",
		"provider",
	},
}