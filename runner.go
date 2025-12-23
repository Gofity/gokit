package gokit

type Runner func() (err error)

func Run(runners ...Runner) (err error) {
	for _, runner := range runners {
		err = runner()

		if err != nil {
			break
		}
	}

	return
}
