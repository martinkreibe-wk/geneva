package elements

// Builder will create an element.
type Builder interface {

	// Build this object into an element.
	Build() (elem Element, err error)
}

// CollectionBuilder will create a collection of elements.
type CollectionBuilder interface {

	// BuildCollection this object into a collection of elements.
	BuildCollection() (elem CollectionElement, err error)
}
