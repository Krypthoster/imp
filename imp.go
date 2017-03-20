package main

import(
"fmt"
"github.com/BurntSushi/toml" 
)

type serverConfig struct {
	Nick	string
	Oauth	string
	channel	string
	Title string
	Port []int
}

type server struct {
    Config serverConfig `toml:"config"`
}

type servers map[string]server

func main(){
var config servers
if _, err := toml.DecodeFile("auth_config.toml",&config); err != nil {
    panic(err)
}

for _, name := range []string{"twitch"} {
    s := config[name]
    fmt.Printf("Config: %s (title: %s) %s\n",
        name, s.Config.Title, s.Config.Nick)
    fmt.Printf("Port: %v\n", s.Config.Port)
}
}


