package config

//MongoConfig contains mongo db server info
type MongoConfig struct {
	//
	IP     string `json:"ip"`
	Port   string `json:"port"`
	DbName string `json:"dbName"`
}

//ServerConfig contains server port
type HttpServerConfig struct {
	Port string `json:"port"`
}

//GetConfig get http server config
func GetConfig() *HttpServerConfig {
	return &HttpServerConfig{
		Port:":8888"}
}