package main 

import "github.com/Netflix/go-env"

type Environment struct {
  Debug int `env:"DEBUG"`
  Extras env.EnvSet
}
