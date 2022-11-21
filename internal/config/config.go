package config

type Config struct {
	Auth []struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"auth"`
	Host     string `json:"host"`
	AuthPort string `json:"auth_port"`
	AcctPort string `json:"acct_port"`
	Secret   string `json:"secret"`
}
