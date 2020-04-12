package cli115

import (
	"github.com/spf13/pflag"
	"os"
)

type Options struct {
	Uid, Cid, Seid string
	CookieFile     string
}

func FromCommandLine() (opts *Options) {
	opts = &Options{}
	fs := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
	fs.StringVarP(&opts.Uid, "uid", "u", "", "Credential UID")
	fs.StringVarP(&opts.Cid, "cid", "c", "", "Credential CID")
	fs.StringVarP(&opts.Seid, "seid", "s", "", "Credential SEID")
	fs.StringVarP(&opts.CookieFile, "cookie-file", "C", "", "Cookie file")
	_ = fs.Parse(os.Args[1:])
	return
}
