package config

//MongoConfig contains mongo db server info
type MongoConfig struct {
	//
	IP     string `json:"ip"`
	Port   string `json:"port"`
	DbName string `json:"dbName"`
}

//HTTPServerConfig contains server port
type HTTPServerConfig struct {
	Port string `json:"port"`
}

//GetHTTPServerConfig get http server config
func GetHTTPServerConfig() *HTTPServerConfig {
	return &HTTPServerConfig{
		Port: ":8888"}
}

//GetMongoConfig returns mongo db data
func GetMongoConfig() *MongoConfig {
	return &MongoConfig{
		IP:     "127.0.0.1:27017",
		Port:   ":27017",
		DbName: "WishDB"}
}
