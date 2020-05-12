package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"path"
)

type Options struct {
	Uid, Cid, Seid string
	CookieFile     string
}

const usage = `Usage: %s [options...]

Options:
  -u, --uid <UID>:   UID value in cookie.
  -c, --cid <CID>:   CID value in cookie.
  -s, --seid <SEID>: SEID value in cookie.
  -C, --cookie-file <file>: Json format cookie file.

`

func ParseCommandLine() (opts *Options) {
	opts, help := &Options{}, false
	// Parse arguments
	fs := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
	fs.StringVarP(&opts.Uid, "uid", "u", "", "")
	fs.StringVarP(&opts.Cid, "cid", "c", "", "")
	fs.StringVarP(&opts.Seid, "seid", "s", "", "")
	fs.StringVarP(&opts.CookieFile, "cookie-file", "C", "", "")
	fs.BoolVarP(&help, "help", "h", false, "")
	fs.Usage = func() {
		fmt.Printf(usage, os.Args[0])
	}
	_ = fs.Parse(os.Args[1:])
	// Exit for help
	if help {
		fs.Usage()
		os.Exit(0)
	}
	// Default options
	if opts.CookieFile == "" {
		dir, _ := os.UserConfigDir()
		opts.CookieFile = path.Join(dir, appName, "cookie.json")
	}
	return
}
