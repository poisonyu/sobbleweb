package global

import "fmt"

type Config struct {
	Mysql
	Jwt
	Local
}

type Local struct {
	Path string `yaml:"path"`
}
type Mysql struct {
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
	UserName string `json:"userName" yaml:"username"`
	PassWord string `json:"passWord" yaml:"password"`
	DbName   string `yaml:"dbName"`
}

// func (mysql Mysql) GormDsn() string {
// 	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysql.UserName, mysql.PassWord, mysql.Host, mysql.Port, mysql.DbName)

// }

func (mysql Mysql) SqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", mysql.UserName, mysql.PassWord, mysql.Host, mysql.Port)
}

type Jwt struct {
	SignKey    string `yaml:"signKey"`
	ExpireTime string `yaml:"expiretime"`
	Issuer     string `yaml:"issuer"`
}
