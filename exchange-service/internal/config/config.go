package config

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

const (
	CONFIG_FILE_TYPE    = "json"
	DEFAULT_CONFIG_PATH = "config.json"
	LOG_JSON_FORMAT     = "json"
	LOG_CONSOLE_FORMAT  = "console"
	LOG_COMPOSE_FORMAT  = "compose"
)

type myFormatter struct {
	log.TextFormatter
}

type Envs struct {
	AppEnv           string
	ServiceName      string
	RestPort         string
	GrpcPort         string
	AuthenticatePort string
	LogLevel         string
	LogFormat        string

	Services struct {
		User struct {
			RestPort string
			GrpcPort string
		}
		Product struct {
			RestPort string
			GrpcPort string
		}
	}
}

var Env Envs

func Initialize(args []string) error {
	configPath := flag.String("config-path","","define the config.json file path to use")
	flag.Parse()

	Env = Envs{}



	log.Info("configPath: ", *configPath)
	if configPath != nil && *configPath != ""  {
		viper.SetConfigFile(*configPath)
	} else {
		viper.SetConfigFile(DEFAULT_CONFIG_PATH)
	}
	viper.SetConfigType(CONFIG_FILE_TYPE)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file json : ", err)
		return err
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
		return err
	}

	if Env.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	configureLogger()

	return nil
}

func configureLogger() {

	logLevel, err := log.ParseLevel(Env.LogLevel)
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.Printf("Log level: %s", logLevel.String())
	log.SetLevel(logLevel)
	log.Infof("LOG FORMAT: %s", Env.LogFormat)

	switch Env.LogFormat {
	case LOG_JSON_FORMAT:
		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: "02-01-2006 15:04:05",
		})
	case LOG_CONSOLE_FORMAT:
		log.SetFormatter(
			&myFormatter{log.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "02-01-2006 15:04:05",
				ForceColors:            true,
				DisableLevelTruncation: false,
			}})
	case LOG_COMPOSE_FORMAT:
		log.SetFormatter(
			&myFormatter{
				log.TextFormatter{
					FullTimestamp:          true,
					TimestampFormat:        "02-01-2006 15:04:05",
					ForceColors:            true,
					DisableLevelTruncation: false,
				}})
	default:
		log.SetFormatter(
			&log.TextFormatter{
				FullTimestamp:          true,
				TimestampFormat:        "02-01-2006 15:04:05",
				ForceColors:            false,
				DisableLevelTruncation: false,
			})

	}

	// TODO: log - add formatters
	// TODO: log - add hooks
}

func (f *myFormatter) Format(entry *log.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 32 // gray
	case log.WarnLevel:
		levelColor = 33 // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("\u001B[%dm%s\u001B[0m - [%s] - %s\n %v",
		levelColor, strings.ToUpper(entry.Level.String()), entry.Time.Format(f.TimestampFormat), entry.Message, entry.Data)), nil
}
