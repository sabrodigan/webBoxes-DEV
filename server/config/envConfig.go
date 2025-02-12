package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
)

var env Conf1gDto

type Conf1gDto struct {
	port        string
	secretKey   string
	databaseURL string
}

func init() {
	if env.port == "" {
		LoadEnvironmentVariable()
	}
	Conf1gEnv()
}

func Conf1gEnv() {
	env = Conf1gDto{
		port:        os.Getenv("PORT"),
		secretKey:   os.Getenv("SECRET_KEY"),
		databaseURL: os.Getenv("MONGODB_URI"),
	}
}

func LoadEnvironmentVariable() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func accessField(key string) (string, error) {
	v := reflect.ValueOf(env)
	t := reflect.TypeOf(env)
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("env is not a struct")
	}
	_, ok := t.FieldByName(key)
	if !ok {
		return "", fmt.Errorf("field %s not found", key)
	}
	f := v.FieldByName(key)
	return f.String(), nil
}

func GetEnvProperty(key string) (string, error) {
	if env.port == "" {
		Conf1gEnv()
	}
	value, err := accessField(key)
	if err != nil {
		fmt.Println("Error loading .env file")
		return value, err
	}
	return accessField(key)
}
