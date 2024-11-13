package provider

// Conf 配置集
type Conf struct {
	Feishu  *Feishu  `json:"feishu" yaml:"feishu"`
	Account *Account `json:"account" yaml:"account"`
}

type Feishu struct {
	AppID     string `json:"appID" yaml:"appID"`
	AppSecret string `json:"appSecret" yaml:"appSecret"`
	BaseUrl   string `json:"baseUrl" yaml:"baseUrl"`
}

type Account struct{}
