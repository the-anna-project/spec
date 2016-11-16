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
	// FSService returns a file system service. It is used to operate on file
	// system abstractions of a certain type.
	FS() FSService
	// IDService returns an ID service. It is used to create IDs of a certain
	// type.
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
	SetConnectionService(connectionService ConnectionService)
	SetEndpointCollection(endpointCollection EndpointCollection)
	SetFeatureService(featureService FeatureService)
	SetForwarderService(forwarderService ForwarderService)
	SetFSService(fsService FSService)
	SetIDService(idService IDService)
	SetInstrumentorService(instrumentorService InstrumentorService)
	SetLogService(logService LogService)
	SetNetworkService(networkService NetworkService)
	SetPermutationService(permutationService PermutationService)
	SetRandomService(randomService RandomService)
	SetStorageCollection(storageCollection StorageCollection)
	// TODO move to InputCollection/OutputCollection => InputService/OutputService
	SetTextInputService(textInputService TextInputService)
	SetTextOutputService(textOutputService TextOutputService)
	SetTrackerService(trackerService TrackerService)
	// Shutdown ends all processes of the service collection like shutting down a
	// machine. The call to Shutdown blocks until the service collection is
	// completely shut down, so you might want to call it in a separate goroutine.
	Shutdown()
	Storage() StorageCollection
	// TextInput returns an text output service. It is used to send text
	// responses back to the client.
	TextInput() TextInputService
	// TextOutput returns an text output service. It is used to send text
	// responses back to the client.
	TextOutput() TextOutputService
	Tracker() TrackerService
}
