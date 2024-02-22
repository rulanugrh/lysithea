package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Server struct {
		Host   string
		Port   string
		Origin string
		Secret string
		URL    string
	}

	Database struct {
		Name string
		Port string
		Host string
		User string
		Pass string
	}

	Elasticsearch struct {
		URL      string
		Username string
		Password string
	}

	ElasticAPM struct {
		Name    string
		APIKey  string
		Version string
		Env     string
	}
}

var app *App

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *App {
	conf := App{}
	if err := godotenv.Load(); err != nil {
		conf.Database.Host = ""
		conf.Database.Name = ""
		conf.Database.Pass = ""
		conf.Database.User = ""
		conf.Database.Port = ""

		conf.Server.Host = ""
		conf.Server.Port = ""
		conf.Server.Origin = ""

		conf.Elasticsearch.URL = ""
		conf.Elasticsearch.Username = ""
		conf.Elasticsearch.Password = ""

		conf.ElasticAPM.APIKey = ""
		conf.ElasticAPM.Env = ""
		conf.ElasticAPM.Version = ""
		conf.ElasticAPM.Name = ""

		return &conf
	}

	conf.Database.Host = os.Getenv("DB_HOST")
	conf.Database.Name = os.Getenv("DB_NAME")
	conf.Database.Pass = os.Getenv("DB_PASS")
	conf.Database.User = os.Getenv("DB_USER")
	conf.Database.Port = os.Getenv("DB_PORT")

	conf.Server.Secret = os.Getenv("APP_SECRET")
	conf.Server.Port = os.Getenv("APP_PORT")
	conf.Server.Host = os.Getenv("APP_HOST")
	conf.Server.Origin = os.Getenv("APP_ORIGIN")
	conf.Server.URL = os.Getenv("APP_URL")

	conf.Elasticsearch.URL = os.Getenv("ELASTICSEARCH_URL")
	conf.Elasticsearch.Username = os.Getenv("ELASTICSEARCH_USERNAME")
	conf.Elasticsearch.Password = os.Getenv("ELASTICSEARCH_PASSWORD")

	conf.ElasticAPM.APIKey = os.Getenv("ELASTIC_APM_SERVICE_API_KEY")
	conf.ElasticAPM.Name = os.Getenv("ELASTIC_APM_SERVICE_NAME")
	conf.ElasticAPM.Version = os.Getenv("ELASTIC_APM_SERVICE_VERSION")
	conf.ElasticAPM.Env = os.Getenv("ELASTIC_APM_SERVICE_ENVIRONTMENT")

	return &conf
}
