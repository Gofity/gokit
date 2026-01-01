//go:build !windows

package gokit

import (
	"os"
	"os/user"
	"strings"
)

func (x xPath) Expand(path String) String {
	if path == "" {
		return path
	}

	u, _ := user.Current()

	if u == nil || u.HomeDir == "" {
		return path
	}

	var home string

	if char := rune(path[0]); char == '~' {
		delim := string([]rune{os.PathSeparator})
		home = strings.TrimRight(u.HomeDir, delim)
		path = path[1:]
	}

	return String(home + os.ExpandEnv(string(path)))
}

func (x xPath) sanitize(path String) (value String) {
	defer func() {
		recover()
		value = path
	}()

	if path[:1] == "/" {
		return
	}

	prefixes := Array[String]{"~/", "./", ".."}

	if !prefixes.Contains(path[:2]) {
		path = "/" + path
	}

	return
}
