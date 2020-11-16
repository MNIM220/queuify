package workers

type Worker struct {
	Process []func(entrance []string) (exit []string)
}
