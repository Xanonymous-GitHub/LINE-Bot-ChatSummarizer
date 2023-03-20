package env

import (
	"github.com/spf13/viper"
	"strings"
)

var (
	IsDebugMode        bool
	Port               string
	RedisHosts         []string
	RedisPassword      string
	DatabaseUrl        string
	EnableRedeem       bool
	ChannelSecret      string
	ChannelAccessToken string
	OpenaiApiToken     string
)

func init() {
	v := viper.GetViper()
	v.AutomaticEnv()
	IsDebugMode = v.GetBool("DEBUG")
	Port = v.GetString("PORT")
	RedisHosts = strings.Split(v.GetString("REDIS_HOSTS"), ",")
	RedisPassword = v.GetString("REDIS_PASSWORD")
	DatabaseUrl = v.GetString("DATABASE_URL")
	EnableRedeem = v.GetBool("REDEEM_ENABLE")
	ChannelSecret = v.GetString("CHANNEL_SECRET")
	ChannelAccessToken = v.GetString("CHANNEL_ACCESS_TOKEN")
	OpenaiApiToken = v.GetString("OPENAI_API_TOKEN")
}
