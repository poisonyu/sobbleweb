package global

import "fmt"

type Config struct {
	Mysql
	Jwt
	Local
	Email
	Redis
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
	SignKey    string `yaml:"signkey"`
	ExpireTime string `yaml:"expiretime"`
	Issuer     string `yaml:"issuer"`
	Audience   string `yaml:"audience"`
}

type Email struct {
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	PassWord string `yaml:"password"`
	Db       string `yaml:"db"`
}
