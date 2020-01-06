package config

import (
	"fmt"

	"github.com/timest/env"
)

//Env Env
var Env *envs

type envs struct {
	GoogleURL string `env:"GOOGLE_URL" default:"https://translate.google.cn"`
}

func init() {
	Env = new(envs)
	env.IgnorePrefix()
	err := env.Fill(Env)
	fmt.Println(Env)
	if err != nil {
		panic(err)
	}
}
