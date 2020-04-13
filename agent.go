package cli115

import (
	"errors"
	"fmt"
	"github.com/deadblue/elevengo"
	"github.com/skip2/go-qrcode"
)

func initAgent(opts *Options) (agent *elevengo.Agent, err error) {
	agent = elevengo.Default()
	// TODO: import credentials from options
	// Start QRCode sign-in process
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
				wait = false
			} else if status.IsCanceled() {
				err = errors.New("user canceled")
				wait = false
			}
		}
	}
	// TODO: handle QRcode expired
	return
}
