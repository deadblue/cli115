package cli115

import (
	"github.com/spf13/pflag"
	"os"
	"path"
)

type Options struct {
	Uid, Cid, Seid string
	CookieFile     string
}

func FromCommandLine() (opts *Options) {
	opts = &Options{}
	fs := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
	fs.StringVarP(&opts.Uid, "uid", "u", "", "UID value in Cookie.")
	fs.StringVarP(&opts.Cid, "cid", "c", "", "CID value in Cookie.")
	fs.StringVarP(&opts.Seid, "seid", "s", "", "SEID value in Cookie.")
	fs.StringVarP(&opts.CookieFile, "cookie-file", "C", "", "Cookie file path.")
	_ = fs.Parse(os.Args[1:])

	if opts.CookieFile == "" {
		// Assume the config dir will always be available.
		dir, _ := os.UserConfigDir()
		opts.CookieFile = path.Join(dir, "cli115", "cookie.json")
	}
	return
}
