package bus

import (
	"context"
	"sync"

	"github.com/aperturerobotics/controllerbus/directive"
	"github.com/aperturerobotics/util/broadcast"
)

// ExecOneOff executes a one-off directive.
//
// If returnIfIdle is set, returns nil, nil, nil if idle.
// If any resolvers return an error, returns that error.
// valDisposeCb is called if the value is no longer valid.
// valDisposeCb might be called multiple times.
// If err != nil, ref == nil.
func ExecOneOff(
	ctx context.Context,
	bus Bus,
	dir directive.Directive,
	returnIfIdle bool,
	valDisposeCallback func(),
) (directive.AttachedValue, directive.Instance, directive.Reference, error) {
	return ExecOneOffWithFilter(ctx, bus, dir, returnIfIdle, valDisposeCallback, nil)
}

// ExecOneOffWithFilter executes a one-off directive with a filter cb.
//
// Waits until the callback returns true before returning a value.
// If returnIfIdle is set, returns nil, nil, nil if idle.
// If any resolvers return an error, returns that error.
// valDisposeCb is called if the value is no longer valid.
// valDisposeCb might be called multiple times.
// If err != nil, ref == nil.
func ExecOneOffWithFilter(
	ctx context.Context,
	bus Bus,
	dir directive.Directive,
	returnIfIdle bool,
	valDisposeCallback func(),
	filterCb func(val directive.AttachedValue) (bool, error),
) (directive.AttachedValue, directive.Instance, directive.Reference, error) {
	// mtx, bcast guard these variables
	var mtx sync.Mutex
	var bcast broadcast.Broadcast

	var val directive.AttachedValue
	var resErr error
	var idle bool

	di, ref, err := bus.AddDirective(
		dir,
		NewCallbackHandler(
			func(v directive.AttachedValue) {
				mtx.Lock()
				if val == nil && resErr == nil {
					ok := filterCb == nil
					if !ok {
						ok, resErr = filterCb(v)
						if resErr != nil {
							bcast.Broadcast()
						}
					}
					if ok && resErr == nil {
						val = v
						bcast.Broadcast()
					}
				}
				mtx.Unlock()
			},
			func(v directive.AttachedValue) {
				if valDisposeCallback == nil {
					return
				}
				mtx.Lock()
				valueRemoved := val != nil && val.GetValueID() == v.GetValueID()
				mtx.Unlock()
				if valueRemoved {
					valDisposeCallback()
				}
			},
			func() {
				mtx.Lock()
				if resErr != nil && !idle {
					resErr = ErrDirectiveDisposed
					bcast.Broadcast()
				}
				mtx.Unlock()
				if valDisposeCallback != nil {
					valDisposeCallback()
				}
			},
		),
	)
	if err != nil {
		if ref != nil {
			ref.Release()
		}
		return nil, nil, nil, err
	}

	defer di.AddIdleCallback(func(errs []error) {
		mtx.Lock()
		defer mtx.Unlock()
		if resErr != nil {
			return
		}
		for _, err := range errs {
			if err != nil {
				resErr = err
				break
			}
		}
		// idle
		if resErr == nil {
			idle = true
		}
		bcast.Broadcast()
	})()

	for {
		mtx.Lock()
		if val != nil {
			val, ref := val, ref // copy
			mtx.Unlock()
			return val, di, ref, nil
		}
		if resErr != nil || (idle && returnIfIdle) {
			err := resErr
			mtx.Unlock()
			ref.Release()
			return nil, nil, nil, err
		}
		waitCh := bcast.GetWaitCh()
		mtx.Unlock()

		select {
		case <-ctx.Done():
			ref.Release()
			return nil, nil, nil, context.Canceled
		case <-waitCh:
		}
	}
}
