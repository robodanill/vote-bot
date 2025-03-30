package config

import "os"

var MattermostToken = os.Getenv("MATTERMOST_TOKEN")
var TarantoolHost = os.Getenv("TARANTOOL_HOST")
var TarantoolPort = os.Getenv("TARANTOOL_PORT")
var TarantoolUser = os.Getenv("TARANTOOL_USER")
var TarantoolPass = os.Getenv("TARANTOOL_PASS")
var AppPort = os.Getenv("APP_PORT")
