package spec

// WorkerService implements a service to process work concurrently.
type WorkerService interface {
	// Execute runs all workers concurrently configured for the current pool. The
	// call to Execute blocks until all workers within the pool have finished
	// their work.
	//
	// The returned error channel is never nil. In case of an error it will
	// always at least contain one error, depending on the worker pool's
	// configuration.
	Execute(config objectspec.WorkerExecuteConfig) chan error
	// ExecuteConfig provides a default configuration for Execute.
	ExecuteConfig() objectspec.WorkerExecuteConfig
}