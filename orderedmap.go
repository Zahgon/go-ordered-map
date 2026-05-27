// Package orderedmap implements an ordered map, i.e. a map that also keeps track of
// the order in which keys were inserted.
//
// All operations are constant-time.
//
// Github repo: https://github.com/wk8/go-ordered-map
package orderedmap

import (
	"iter"

	list "github.com/bahlo/generic-list-go"
)

type Pair[K comparable, V any] struct {
	Key   K
	Value V

	element *list.Element[*Pair[K, V]]
}

type OrderedMap[K comparable, V any] struct {
	pairs             map[K]*Pair[K, V]
	list              *list.List[*Pair[K, V]]
	disableHTMLEscape bool
}

type initConfig[K comparable, V any] struct {
	capacity          int
	initialData       []Pair[K, V]
	disableHTMLEscape bool
}

type InitOption[K comparable, V any] func(config *initConfig[K, V])

// WithCapacity allows giving a capacity hint for the map, akin to the standard make(map[K]V, capacity).
func WithCapacity[K comparable, V any](capacity int) InitOption[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// WithInitialData allows passing in initial data for the map.
func WithInitialData[K comparable, V any](initialData ...Pair[K, V]) InitOption[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// WithDisableHTMLEscape disables HTMl escaping when marshalling to JSON
func WithDisableHTMLEscape[K comparable, V any]() InitOption[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// New creates a new OrderedMap.
// options can either be one or several InitOption[K, V], or a single integer,
// which is then interpreted as a capacity hint, à la make(map[K]V, capacity).
func New[K comparable, V any](options ...any) *OrderedMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

const invalidOptionMessage = `when using orderedmap.New[K,V]() with options, either provide one or several InitOption[K, V]; or a single integer which is then interpreted as a capacity hint, à la make(map[K]V, capacity).` //nolint:lll

func invalidOption() { _ = "STUB: not implemented"; return }

func (om *OrderedMap[K, V]) initialize(capacity int, disableHTMLEscape bool) {
	_ = "STUB: not implemented"
	return
}

// Get looks for the given key, and returns the value associated with it,
// or V's nil value if not found. The boolean it returns says whether the key is present in the map.
func (om *OrderedMap[K, V]) Get(key K) (val V, present bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Load is an alias for Get, mostly to present an API similar to `sync.Map`'s.
func (om *OrderedMap[K, V]) Load(key K) (V, bool) {
	_ = "STUB: not implemented"

	// Value returns the value associated with the given key or the zero value.
	return *new(V), false
}

func (om *OrderedMap[K, V]) Value(key K) (val V) { _ = "STUB: not implemented"; return *new(V) }

// GetPair looks for the given key, and returns the pair associated with it,
// or nil if not found. The Pair struct can then be used to iterate over the ordered map
// from that point, either forward or backward.
func (om *OrderedMap[K, V]) GetPair(key K) *Pair[K, V] { _ = "STUB: not implemented"; return nil }

// Set sets the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Set`.
func (om *OrderedMap[K, V]) Set(key K, value V) (val V, present bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// AddPairs allows setting multiple pairs at a time. It's equivalent to calling
// Set on each pair sequentially.
func (om *OrderedMap[K, V]) AddPairs(pairs ...Pair[K, V]) { _ = "STUB: not implemented"; return }

// Store is an alias for Set, mostly to present an API similar to `sync.Map`'s.
func (om *OrderedMap[K, V]) Store(key K, value V) (V, bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Delete removes the key-value pair, and returns what `Get` would have returned
// on that key prior to the call to `Delete`.
func (om *OrderedMap[K, V]) Delete(key K) (val V, present bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Len returns the length of the ordered map.
func (om *OrderedMap[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

// Oldest returns a pointer to the oldest pair. It's meant to be used to iterate on the ordered map's
// pairs from the oldest to the newest, e.g.:
// for pair := orderedMap.Oldest(); pair != nil; pair = pair.Next() { fmt.Printf("%v => %v\n", pair.Key, pair.Value) }
func (om *OrderedMap[K, V]) Oldest() *Pair[K, V] { _ = "STUB: not implemented"; return nil }

// Newest returns a pointer to the newest pair. It's meant to be used to iterate on the ordered map's
// pairs from the newest to the oldest, e.g.:
// for pair := orderedMap.Newest(); pair != nil; pair = pair.Prev() { fmt.Printf("%v => %v\n", pair.Key, pair.Value) }
func (om *OrderedMap[K, V]) Newest() *Pair[K, V] { _ = "STUB: not implemented"; return nil }

// Next returns a pointer to the next pair.
func (p *Pair[K, V]) Next() *Pair[K, V] { _ = "STUB: not implemented"; return nil }

// Prev returns a pointer to the previous pair.
func (p *Pair[K, V]) Prev() *Pair[K, V] { _ = "STUB: not implemented"; return nil }

func listElementToPair[K comparable, V any](element *list.Element[*Pair[K, V]]) *Pair[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// KeyNotFoundError may be returned by functions in this package when they're called with keys that are not present
// in the map.
type KeyNotFoundError[K comparable] struct {
	MissingKey K
}

func (e *KeyNotFoundError[K]) Error() string { _ = "STUB: not implemented"; return "" }

// MoveAfter moves the value associated with key to its new position after the one associated with markKey.
// Returns an error iff key or markKey are not present in the map. If an error is returned,
// it will be a KeyNotFoundError.
func (om *OrderedMap[K, V]) MoveAfter(key, markKey K) error { _ = "STUB: not implemented"; return nil }

// MoveBefore moves the value associated with key to its new position before the one associated with markKey.
// Returns an error iff key or markKey are not present in the map. If an error is returned,
// it will be a KeyNotFoundError.
func (om *OrderedMap[K, V]) MoveBefore(key, markKey K) error { _ = "STUB: not implemented"; return nil }

func (om *OrderedMap[K, V]) getElements(keys ...K) ([]*list.Element[*Pair[K, V]], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MoveToBack moves the value associated with key to the back of the ordered map,
// i.e. makes it the newest pair in the map.
// Returns an error iff key is not present in the map. If an error is returned,
// it will be a KeyNotFoundError.
func (om *OrderedMap[K, V]) MoveToBack(key K) error { _ = "STUB: not implemented"; return nil }

// MoveToFront moves the value associated with key to the front of the ordered map,
// i.e. makes it the oldest pair in the map.
// Returns an error iff key is not present in the map. If an error is returned,
// it will be a KeyNotFoundError.
func (om *OrderedMap[K, V]) MoveToFront(key K) error { _ = "STUB: not implemented"; return nil }

// GetAndMoveToBack combines Get and MoveToBack in the same call. If an error is returned,
// it will be a KeyNotFoundError.
func (om *OrderedMap[K, V]) GetAndMoveToBack(key K) (val V, err error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

// GetAndMoveToFront combines Get and MoveToFront in the same call. If an error is returned,
// it will be a KeyNotFoundError.
func (om *OrderedMap[K, V]) GetAndMoveToFront(key K) (val V, err error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

// FromOldest returns an iterator over all the key-value pairs in the map, starting from the oldest pair.
func (om *OrderedMap[K, V]) FromOldest() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }

// FromNewest returns an iterator over all the key-value pairs in the map, starting from the newest pair.
func (om *OrderedMap[K, V]) FromNewest() iter.Seq2[K, V] { _ = "STUB: not implemented"; return nil }

// KeysFromOldest returns an iterator over all the keys in the map, starting from the oldest pair.
func (om *OrderedMap[K, V]) KeysFromOldest() iter.Seq[K] { _ = "STUB: not implemented"; return nil }

// KeysFromNewest returns an iterator over all the keys in the map, starting from the newest pair.
func (om *OrderedMap[K, V]) KeysFromNewest() iter.Seq[K] { _ = "STUB: not implemented"; return nil }

// ValuesFromOldest returns an iterator over all the values in the map, starting from the oldest pair.
func (om *OrderedMap[K, V]) ValuesFromOldest() iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// ValuesFromNewest returns an iterator over all the values in the map, starting from the newest pair.
func (om *OrderedMap[K, V]) ValuesFromNewest() iter.Seq[V] { _ = "STUB: not implemented"; return nil }

// From creates a new OrderedMap from an iterator over key-value pairs.
func From[K comparable, V any](i iter.Seq2[K, V]) *OrderedMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (om *OrderedMap[K, V]) Filter(predicate func(K, V) bool) { _ = "STUB: not implemented"; return }
