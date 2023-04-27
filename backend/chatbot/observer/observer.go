package observer

// Observer
type Observer interface {
	Update(data interface{})
}

// Subject Observer interface
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(data interface{})
}
