package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"time"
)

type cfg struct {
	RSS           []string `json:"rss"`
	RequestPeriod int      `json:"request_period"`
	Timeout       time.Duration
	Postgres      struct {
		DSN      string
		Host     string `json:"host" env-default:"postgres"`
		Port     string `json:"port" env-default:"5432"`
		Dbname   string `json:"db_name" env-default:"postgres"`
		Username string `json:"user_name" env-default:"postgres"`
		Pwd      string `json:"pwd" env-default:"postgres"`
	} `json:"postgres"`
	NewsHost          string `json:"news_host" env-default:"localhost:8081"`
	CommentsHost      string `json:"comments_host" env-default:"localhost:8083"`
	CommentCensorHost string `json:"comment_censor_host" env-default:"localhost:8084"`
}

func config() cfg {
	configs := cfg{}
	if err := cleanenv.ReadConfig("./config/config.json", &configs); err != nil {
		log.Println("cannot read configs")
	}
	configs.Postgres.DSN = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", configs.Postgres.Dbname, configs.Postgres.Username,
		configs.Postgres.Pwd, configs.Postgres.Host, configs.Postgres.Port, configs.Postgres.Dbname)
	configs.Timeout = time.Duration(configs.RequestPeriod) * time.Minute
	return configs
}

var Configs = config()
