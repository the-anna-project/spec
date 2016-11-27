package object

type WorkerExecuteConfig interface {
	Action() func(canceler <-chan struct{}) error
	Canceler() chan struct{}
	CancelOnError() bool
	NumWorkers() int
}
