package fs

import (
	"go.dead.blue/cli115/internal/pkg/util"
	"path"
)

const specialChars = "*?[]\\ "

func escape(name string) string {
	return util.StdEscape(name, specialChars)
}

func mustMatch(pattern, name string) bool {
	if ok, err := path.Match(pattern, name); err != nil {
		return false
	} else {
		return ok
	}
}
