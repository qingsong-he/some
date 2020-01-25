package main

import (
	"github.com/go-redis/redis"
	"github.com/qingsong-he/ce"
	"os"
	"reflect"
)

func init() {
	ce.Print(os.Args[0])
}

func main() {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := cli.Ping().Result()
	ce.CheckError(err)
	defer cli.Close()

	scriptHash, err := cli.ScriptLoad(`return {KEYS[1], KEYS[2], ARGV[1], ARGV[2]}`).Result()
	ce.CheckError(err)
	ce.Print("script hash:", scriptHash)

	result, err := cli.EvalSha(scriptHash, []string{"key1", "key2"}, "v1", "v2").Result()
	ce.CheckError(err)

	ce.Print(reflect.TypeOf(result).String(), result)
}
