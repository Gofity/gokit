package gokit

import (
	"os"
	"path/filepath"
	"runtime"
)

type PathFn func() (path String, err error)

type xPath struct{}

func (x xPath) Delim() string {
	return string([]rune{os.PathSeparator})
}

func (x xPath) Join(paths ...any) String {
	delim := x.Delim()

	return JoinFn(delim, paths, func(v String) String {
		return v.TrimAffix(delim)
	})
}

func (x xPath) JoinPrefixed(paths ...any) String {
	result := x.Join(paths...)
	return String(x.Delim()) + result
}

func (x xPath) FromExecutable(paths ...any) (file String, err error) {
	return x.get(
		x.getExecPathFromArgs(paths...),
		x.getExecPathFromCaller(0, paths...),
		x.getExecPathFromCaller(1, paths...),
		x.getExecPathFromSource(paths...),
	)
}

func (x xPath) IsDir(file string) bool {
	info, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

func (x xPath) IsFile(file string) bool {
	info, err := os.Stat(file)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func (x xPath) get(fns ...PathFn) (path String, err error) {
	for _, fn := range fns {
		path, err = fn()

		if err != nil {
			continue
		}

		if x.IsFile(string(path)) {
			break
		}
	}

	return
}

func (x xPath) getExecPathFromCaller(skip int, paths ...any) PathFn {
	return func() (path String, err error) {
		_, file, _, _ := runtime.Caller(skip)
		return x.getAbsolutePathFromFile(file, paths...)
	}
}

func (x xPath) getExecPathFromArgs(paths ...any) PathFn {
	return func() (path String, err error) {
		return x.getAbsolutePathFromFile(os.Args[0], paths...)
	}
}

func (x xPath) getExecPathFromSource(paths ...any) PathFn {
	return func() (path String, err error) {
		return x.getAbsolutePathFromDir("./", paths...)
	}
}

func (x xPath) getAbsolutePathFromFile(fromFile string, paths ...any) (path String, err error) {
	dir := filepath.Dir(fromFile)
	return x.getAbsolutePathFromDir(dir, paths...)
}

func (x xPath) getAbsolutePathFromDir(dir string, paths ...any) (path String, err error) {
	if dir, err = filepath.Abs(dir); err != nil {
		return
	}

	allPaths := append([]any{dir}, paths...)

	path = x.Join(allPaths...)
	path = x.sanitize(path)

	return
}

var Path xPath
