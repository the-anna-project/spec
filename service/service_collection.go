package spec

// ServiceCollection represents a collection of factories. This scopes different
// service implementations in a simple container, which can easily be passed
// around.
type ServiceCollection interface {
	Activator() ActivatorService
	Boot()
	Connection() ConnectionService
	Endpoint() EndpointCollection
	Feature() FeatureService
	Forwarder() ForwarderService
	// FS returns a file system service. It is used to operate on file system
	// abstractions of a certain type.
	FS() FSService
	// ID returns an ID service. It is used to create IDs of a certain type.
	ID() IDService
	Instrumentor() InstrumentorService
	// Log returns a log service. It is used to print log messages.
	Log() LogService
	Network() NetworkService
	// Permutation returns a permutation service. It is used to permute instances
	// of type PermutationList.
	Permutation() PermutationService
	// Random returns a random service. It is used to create random numbers.
	Random() RandomService
	SetActivatorService(activatorService ActivatorService)
	SetConnectionService(connection ConnectionService)
	SetEndpointCollection(endpointCollection EndpointCollection)
	SetFeatureService(feature FeatureService)
	SetForwarderService(forwarder ForwarderService)
	SetFSService(FS FSService)
	SetIDService(ID IDService)
	SetInstrumentorService(instrumentor InstrumentorService)
	SetLogService(log LogService)
	SetNetworkService(network NetworkService)
	SetPermutationService(permutation PermutationService)
	SetRandomService(random RandomService)
	SetStorageCollection(storageCollection StorageCollection)
	// TODO move to InputCollection/OutputCollection => InputService/OutputService
	SetTextInputService(textInput TextInputService)
	SetTextOutputService(textOutput TextOutputService)
	SetTrackerService(tracker TrackerService)
	// Shutdown ends all processes of the service collection like shutting down a
	// machine. The call to Shutdown blocks until the service collection is
	// completely shut down, so you might want to call it in a separate goroutine.
	Shutdown()
	Storage() StorageCollection
	// TextInput returns an text output service. It is used to send text
	// responses back to the client.
	TextInput() TextInpuService)
	// TextOutput returns an text output service. It is used to send text
	// responses back to the client.
	TextOutput() TextOutpuService)
	Tracker() TrackeService)
}
