package config

//MongoConfig contains mongo db server info
type MongoConfig struct {
	//
	IP     string `json:"ip"`
	Port   string `json:"port"`
	DbName string `json:"dbName"`
}

//ServerConfig contains server port
type ServerConfig struct {
	Port string `json:"port"`
}
