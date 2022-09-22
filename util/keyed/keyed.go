package keyed

import (
	"context"
	"sort"
	"sync"
)

// Routine is a function called as a goroutine.
// If nil is returned, exits cleanly permanently.
// If an error is returned, can be restarted later.
type Routine func(ctx context.Context) error

// Keyed manages a set of goroutines with associated Keys.
type Keyed[T comparable] struct {
	// ctorCb is the constructor callback
	ctorCb func(key string) (Routine, T)

	// mtx guards below fields
	mtx sync.Mutex
	// ctx is the current root context
	ctx context.Context
	// routines is the set of running routines
	routines map[string]*runningRoutine[T]
}

// NewKeyed constructs a new Keyed execution manager.
// Note: routines won't start until SetContext is called.
func NewKeyed[T comparable](ctorCb func(key string) (Routine, T)) *Keyed[T] {
	if ctorCb == nil {
		ctorCb = func(key string) (Routine, T) {
			var empty T
			return nil, empty
		}
	}
	return &Keyed[T]{
		ctorCb:   ctorCb,
		routines: make(map[string]*runningRoutine[T], 1),
	}
}

// SetContext updates the root context, restarting all running routines.
// if restart is true, all errored routines also restart
func (k *Keyed[T]) SetContext(ctx context.Context, restart bool) {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	sameCtx := k.ctx == ctx
	if sameCtx && !restart {
		return
	}

	k.ctx = ctx
	for _, rr := range k.routines {
		if sameCtx && rr.err == nil {
			continue
		}
		if rr.err == nil || restart {
			if rr.ctxCancel != nil {
				rr.ctxCancel()
				rr.ctx, rr.ctxCancel = nil, nil
			}
			rr.start(ctx)
		}
	}
}

// GetKeys returns the list of keys registered with the Keyed instance.
// Note: this is an instantaneous snapshot.
func (k *Keyed[T]) GetKeys() []string {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	keys := make([]string, 0, len(k.routines))
	for k := range k.routines {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// KeyWithData is a key with associated data.
type KeyWithData[T comparable] struct {
	// Key is the key.
	Key string
	// Data is the associated data.
	Data T
}

// GetKeysWithData returns the keys and the data for the keys.
// Note: this is an instantaneous snapshot.
func (k *Keyed[T]) GetKeysWithData() []KeyWithData[T] {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	out := make([]KeyWithData[T], 0, len(k.routines))
	for k, v := range k.routines {
		out = append(out, KeyWithData[T]{
			Key:  k,
			Data: v.data,
		})
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].Key < out[j].Key
	})
	return out
}

// SyncKeys synchronizes the list of running routines with the given list.
// If restart=true, restarts any failed routines in the list.
func (k *Keyed[T]) SyncKeys(keys []string, restart bool) {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	if k.ctx != nil {
		select {
		case <-k.ctx.Done():
			k.ctx = nil
		default:
		}
	}

	routines := make(map[string]*runningRoutine[T], len(keys))
	for _, key := range keys {
		v := routines[key]
		if v != nil {
			continue
		}
		v, existed := k.routines[key]
		if !existed {
			routine, data := k.ctorCb(key)
			v = newRunningRoutine(k, routine, data)
			k.routines[key] = v
		}
		routines[key] = v
		if !existed || restart {
			if k.ctx != nil {
				v.start(k.ctx)
			}
		}
	}
	for key, rr := range k.routines {
		if _, ok := routines[key]; ok {
			continue
		}
		if rr.ctxCancel != nil {
			rr.ctxCancel()
		}
		delete(k.routines, key)
	}
}

// GetKey returns the routine for the given key.
// Note: this is an instantaneous snapshot.
func (k *Keyed[T]) GetKey(key string) (Routine, T) {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	v, existed := k.routines[key]
	if !existed {
		var empty T
		return nil, empty
	}

	return v.routine, v.data
}

// ResetRoutine resets the given routine after checking the condition functions.
// If any return true, resets the instance.
//
// If len(conds) == 0, always resets the given key.
func (k *Keyed[T]) ResetRoutine(key string, conds ...func(T) bool) (existed bool, reset bool) {
	k.mtx.Lock()
	defer k.mtx.Unlock()

	if k.ctx != nil {
		select {
		case <-k.ctx.Done():
			k.ctx = nil
		default:
		}
	}

	v, existed := k.routines[key]
	if !existed {
		return false, false
	}
	anyMatched := len(conds) == 0
	for _, cond := range conds {
		if cond != nil && cond(v.data) {
			anyMatched = true
			break
		}
	}
	if !anyMatched {
		return true, false
	}

	if v.ctxCancel != nil {
		v.ctxCancel()
	}
	routine, data := k.ctorCb(key)
	v = newRunningRoutine(k, routine, data)
	k.routines[key] = v
	if k.ctx != nil {
		v.start(k.ctx)
	}

	return true, true
}
