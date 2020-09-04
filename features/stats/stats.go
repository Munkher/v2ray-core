package stats

//go:generate errorgen

import "v2ray.com/core/features"

// Counter is the interface for stats counters.
//
// v2ray:api:stable
type Counter interface {
	// Value is the current value of the counter.
	Value() int64
	// Set sets a new value to the counter, and returns the previous one.
	Set(int64) int64
	// Add adds a value to the current counter value, and returns the previous value.
	Add(int64) int64
}

// Channel is the interface for stats channel
//
// v2ray:api:stable
type Channel interface {
	// Channel returns the underlying go channel.
	Channel() chan interface{}
	// SubscriberCount returns the number of the subscribers.
	Subscribers() []chan interface{}
	// Subscribe registers for listening to channel stream and returns a new listener channel.
	Subscribe() chan interface{}
	// Unsubscribe unregisters a listener channel from current Channel object.
	Unsubscribe(chan interface{})
}

// Manager is the interface for stats manager.
//
// v2ray:api:stable
type Manager interface {
	features.Feature

	// RegisterCounter registers a new counter to the manager. The identifier string must not be empty, and unique among other counters.
	RegisterCounter(string) (Counter, error)
	// GetCounter returns a counter by its identifier.
	GetCounter(string) Counter

	// RegisterChannel registers a new channel to the manager. The identifier string must not be empty, and unique among other channels.
	RegisterChannel(string) (Channel, error)
	// GetChannel returns a channel by its identifier.
	GetChannel(string) Channel
}

// GetOrRegisterCounter tries to get the StatCounter first. If not exist, it then tries to create a new counter.
func GetOrRegisterCounter(m Manager, name string) (Counter, error) {
	counter := m.GetCounter(name)
	if counter != nil {
		return counter, nil
	}

	return m.RegisterCounter(name)
}

// GetOrRegisterChannel tries to get the StatChannel first. If not exist, it then tries to create a new channel.
func GetOrRegisterChannel(m Manager, name string) (Channel, error) {
	channel := m.GetChannel(name)
	if channel != nil {
		return channel, nil
	}

	return m.RegisterChannel(name)
}

// ManagerType returns the type of Manager interface. Can be used to implement common.HasType.
//
// v2ray:api:stable
func ManagerType() interface{} {
	return (*Manager)(nil)
}

// NoopManager is an implementation of Manager, which doesn't has actual functionalities.
type NoopManager struct{}

// Type implements common.HasType.
func (NoopManager) Type() interface{} {
	return ManagerType()
}

// RegisterCounter implements Manager.
func (NoopManager) RegisterCounter(string) (Counter, error) {
	return nil, newError("not implemented")
}

// GetCounter implements Manager.
func (NoopManager) GetCounter(string) Counter {
	return nil
}

// RegisterChannel implements Manager.
func (NoopManager) RegisterChannel(string) (Channel, error) {
	return nil, newError("not implemented")
}

// GetChannel implements Manager.
func (NoopManager) GetChannel(string) Channel {
	return nil
}

// Start implements common.Runnable.
func (NoopManager) Start() error { return nil }

// Close implements common.Closable.
func (NoopManager) Close() error { return nil }
