package utility

import (
	"os"
	"flag"
	"strings"
)

func GetFlagValue(flag bool, trueValue interface{}, falseValue interface{}) (resp interface{}) {
	if flag {
		return trueValue
	} else {
		return falseValue
	}
}

func GetArrary(env string, value string, name string, desc string, sep string) []string {
	v := os.Getenv(env)
	if v == "" {
		v = value
	}
	var arrarystr string
	flag.StringVar(&arrarystr, name, v, desc)
	arrary := strings.Split(arrarystr, sep)
	return arrary
}
