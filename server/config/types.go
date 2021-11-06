package config

// Config main configuration struct
type Config struct {
	Production            bool       `json:"production"`
	SystemID              int8       `json:"system_id"`
	Secrets               []string   `json:"secrets"`
	SecretJWT             string     `json:"secret_jwt"`
	Databases             []DataBase `json:"databases"`
	InternalAddress       string     `json:"internal_address"`
	ExternalAddress       string     `json:"public_address"`
	ExternalPublicAddress string     `json:"external_public_address"`
	AccessLogDirectory    string     `json:"access_log_directory"`
	ErrorLogDirectory     string     `json:"error_log_directory"`
	UploadDirectory       string     `json:"upload_directory"`
	MaximumUploadSize     int64      `json:"maximum_upload_size"`
	QueryTimeout          float32    `json:"query_timeout"`
}

// DataBase contém dados necessários para manter uma conexão com o banco de dados
type DataBase struct {
	Nick               string   `json:"nick"`
	Name               string   `json:"name"`
	Username           string   `json:"username"`
	Password           string   `json:"password"`
	Host               string   `json:"hostname"`
	Port               string   `json:"port"`
	MaxConn            int      `json:"max_conn"`
	MaxIdle            int      `json:"max_idle"`
	ReadOnly           bool     `json:"read_only"`
	Main               bool     `json:"main"`
	Addresses          []string `json:"addresses"`
	TransactionTimeout int      `json:"transaction_timeout"`
}
