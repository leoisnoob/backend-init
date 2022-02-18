package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg Config

// yaml file 的优先级高于环境变量
func InitConfig(cfgPath string) {
	log.Println("init config")
	f, err := os.Open(cfgPath)
	if err != nil {
		log.Fatalf("open config file %s failed\n", cfgPath)
	}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		log.Fatalf("decoder yaml config failed, error %s\n", err)
	}

}

type LDAP struct {
	Addr     string `yaml:"addr"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	Ldap     LDAP   `yaml:"ldap"`
	MysqlDNS string `yaml:"mysql_dns"`
	OsgUrl   string `yaml:"osg_url"`
	Minio    Minio  `yaml:"minio"`
}

type Mysql struct {
	Dns  string `yaml:"dns" validate:"required`
	pass string `yaml:"pass"`
}

type Minio struct {
	URL        string `yaml:"url"`
	AccessKey  string `yaml:"access_key"`
	SecretKey  string `yaml:"secret_key"`
	BucketName string `yaml:"bucket_name"`
}
