package graceful

import "context"

type initializable interface {
	initialize()
}

func d[T initializable](x T) T {
	x.initialize()
	return x
}

type Runner func(grace Grace) (err error)

type Grace interface {
	context.Context
}

func Run(callback ...Runner) (err error) {
	for _, fn := range callback {
		err = instance.Run(fn)

		if err != nil {
			break
		}
	}

	return
}

func IsExiting() bool {
	return instance.isExiting()
}
