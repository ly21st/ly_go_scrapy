package api

type DBConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Passwd   string `json:"passwd"`
	Database string `json:"database"`
}

var (
	DbConfig DBConfig
)

func ParserConfig() {
	DbConfig.Host = "localhost"
	DbConfig.Port = 0
	DbConfig.User = ""
	DbConfig.Passwd = ""
	DbConfig.Database = "./data/login_user"
}

