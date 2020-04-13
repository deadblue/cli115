package cli115

import (
	"dead.blue/cli115/core"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deadblue/elevengo"
	"github.com/skip2/go-qrcode"
	"os"
	"path"
)

type CookieData struct {
	Uid  string `json:"uid"`
	Cid  string `json:"cid"`
	Seid string `json:"seid"`
}

func initAgent(opts *Options) (agent *elevengo.Agent, err error) {
	agent = elevengo.Default()
	// try load cookie
	if cr, err := loadCookie(opts); err == nil {
		if err = agent.CredentialsImport(cr); err == nil {
			return agent, nil
		}
	}
	// prompt user to sign in
	err = signIn(agent)
	return
}

func loadCookie(opts *Options) (cr *elevengo.Credentials, err error) {
	// make credentials by arguments
	if opts.Uid != "" && opts.Cid != "" && opts.Seid != "" {
		cr = &elevengo.Credentials{
			UID:  opts.Uid,
			CID:  opts.Cid,
			SEID: opts.Seid,
		}
		return
	}
	// use default cookie path if not specify.
	// default path is $CONFIG_DIR/cli115/cookie.json
	if opts.CookieFile == "" {
		if dir, err := os.UserConfigDir(); err == nil {
			opts.CookieFile = path.Join(dir, "cli115", "cookie.json")
		}
	}
	// try load cookie file
	file, err := os.Open(opts.CookieFile)
	if err != nil {
		return
	}
	defer core.QuietlyClose(file)
	// decode cookie
	jd, data := json.NewDecoder(file), &CookieData{}
	if err = jd.Decode(data); err == nil {
		cr = &elevengo.Credentials{
			UID:  data.Uid,
			CID:  data.Cid,
			SEID: data.Seid,
		}
	}
	return
}

func signIn(agent *elevengo.Agent) (err error) {
	// TODO: handle QRcode expired
	// load cookie failed, start QRcode signin
	session, err := agent.QrcodeStart()
	if err != nil {
		return
	}
	// Print QRcode
	code, err := qrcode.New(session.Content, qrcode.Medium)
	if err != nil {
		return
	}
	fmt.Println("Please scan the QRcode on mobile, add allow this sign-in:")
	fmt.Print(code.ToSmallString(false))
	// Wait sign-in
	for wait := true; wait; {
		if status, serr := agent.QrcodeStatus(session); serr != nil {
			err = serr
		} else {
			if status.IsAllowed() {
				err = agent.QrcodeLogin(session)
				// TODO: store cookie into cookie file
				wait = false
			} else if status.IsCanceled() {
				err = errors.New("user canceled")
				wait = false
			}
		}

	}
	return
}
