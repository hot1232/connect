package conf

import (
	"errors"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var Setting Settings

type Config struct {

}

type ApiConfig struct {
	Address string `yaml:"address"`
	Port int `yaml:"port"`
}

type WebsocketConfig struct {
	Address string `yaml:"address"`
	Port int `yaml:"port"`
}

type AuthConfig struct {
	Provider string `yaml:"provider"`
	Url string `yaml:"url"`
	ProviderAuth bool `yaml:"auth"`
	ProviderAuthUser string `yaml:"username"`
	ProviderAuthPassword string `yaml:"password"`
}

type AuditConfig struct {
	DbType string `yaml:"dbtype"`
	DbHost string `yaml:"dbhost"`
	DbPort int `yaml:"dbport"`
	DbUser string `yaml:"dbuser"`
	DbPassword string `yaml:"dbpassword"`
	DbName string `yaml:"dbname"`
}

type SshdConfig struct {
	ListenPort int `yaml:"listenport"`
	ListenAddr string `yaml:"listenaddr"`
	TimeOut int `yaml:"timeout"`
	HostKey string `yaml:"hostkey"`
	AuthorisedKeys string `yaml:"authorisedkeys"`
	Auth AuthConfig `yaml:"auth"`
}

type CacheConfig struct {
	Path string `yaml:"path"`
	Format string `yaml:"format"`
}

type Settings struct {
	MaxConn int `yaml:"maxconn"`
	KeepAlive int `yaml:"keepalive"`
	Verbose bool `yaml:"verbose"`
	Debug bool `yaml:"debug"`
	LoginTitle string `yaml:"logintitle"`
	Ps1 string `yaml:"ps1"`

	Sshd SshdConfig `yaml:"sshd"`
	Cache CacheConfig `yaml:"cache"`
	Audit AuditConfig `yaml:"audit"`
	Api ApiConfig `yaml:"api"`
	Websocket WebsocketConfig `yaml:"websocket"`
}

func NewSettings() (Config){
	return Config{}
}

func (self *Config) FindConfigFile() (string,error) {
	err := errors.New("not found")
	return "",err
}


func (self *Config) Load(f string) (Settings,error) {
	content,err := ioutil.ReadFile(f)
	if err != nil {
		return Settings{},err
	}

	settings := &Settings{}
	err = yaml.Unmarshal([]byte(content),settings)
	if err != nil {
		return Settings{},err
	}

	return *settings,nil
}
