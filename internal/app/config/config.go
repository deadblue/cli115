package config

import (
	"go.dead.blue/cli115/internal/pkg/util"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
	"path"
)

const (
	appName = "cli115"
)

type Aria2Conf struct {
	// Full path of aria2
	Path string `yaml:"path"`
	// RPC mode flag
	Rpc bool `yaml:"rpc"`
	// RPC endpoint
	Endpoint string `yaml:"endpoint"`
	// RPC token
	Token string `yaml:"token"`
	// Download directory
	Dir string `yaml:"dir"`
}

type CurlConf struct {
	// Full path of curl
	Path string `yaml:"path"`
}

type MpvConf struct {
	// Full path of mpv
	Path string `yaml:"path"`
	// Start in full-screen mode
	Fs bool `yaml:"fs"`
}

type Conf struct {
	// Aria2 config
	Aria2 *Aria2Conf `json:"aria2"`
	// Curl config
	Curl *CurlConf `yaml:"curl"`
	// Mpv config
	Mpv *MpvConf `yaml:"mpv"`
}

func Load() (conf *Conf) {
	conf = &Conf{}
	// Load conf from file
	confDir, _ := os.UserConfigDir()
	confFile := path.Join(confDir, appName, "conf.yaml")
	if file, err := os.Open(confFile); err == nil {
		defer util.QuietlyClose(file)
		_ = yaml.NewDecoder(file).Decode(conf)
	}
	// Default config
	if conf.Aria2 == nil {
		conf.Aria2 = &Aria2Conf{Path: "", Rpc: false}
		if exe, err := exec.LookPath("aria2c"); err == nil {
			conf.Aria2.Path = exe
		}
	}
	if conf.Curl == nil {
		conf.Curl = &CurlConf{}
		if exe, err := exec.LookPath("curl"); err == nil {
			conf.Curl.Path = exe
		}
	}
	if conf.Mpv == nil {
		conf.Mpv = &MpvConf{
			Fs: false,
		}
		if exe, err := exec.LookPath("mpv"); err == nil {
			conf.Mpv.Path = exe
		}
	}
	return
}
