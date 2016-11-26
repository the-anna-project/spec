package spec

// LayerCollection represents a collection of layer services. This scopes
// different layer service implementations in a simple container, which can
// easily be passed around.
type LayerCollection interface {
	Boot()
	Behaviour() servicespec.LayerService
	Information() servicespec.LayerService
	SetBehaviour(behaviourService servicespec.LayerService)
	SetInformation(informationService servicespec.LayerService)
}
