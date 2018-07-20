package conf

import (
	"encoding/json"
	"fmt"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("%v", err)
	}
	fmt.Println(dir)

	data, err := ioutil.ReadFile(dir + "/conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
