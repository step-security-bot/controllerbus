package loader

import (
	"context"
	"time"

	"github.com/aperturerobotics/controllerbus/controller"
	"github.com/aperturerobotics/controllerbus/directive"
	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"
)

// resolver tracks a ExecController request
type resolver struct {
	ctx         context.Context
	directive   ExecController
	controller  *Controller
	execBackoff backoff.BackOff
	lastErr     error
}

// newResolver builds a new ExecController resolver.
func newResolver(ctx context.Context, directive ExecController, controller *Controller) *resolver {
	ebo := backoff.NewExponentialBackOff()
	ebo.InitialInterval = time.Millisecond * 500
	ebo.MaxInterval = time.Second * 30
	ebo.Multiplier = 2
	ebo.MaxElapsedTime = time.Minute
	return &resolver{
		ctx:         ctx,
		directive:   directive,
		controller:  controller,
		execBackoff: ebo,
	}
}

// resolveExecController handles a ExecController directive.
func (c *Controller) resolveExecController(
	ctx context.Context,
	dir ExecController,
) (directive.Resolver, error) {
	// Check if the ExecController is meant for / compatible with us.
	// In this case, we handle all ExecController requests.
	return newResolver(ctx, dir, c), nil
}

// Resolve resolves the values.
// Any fatal error resolving the value is returned.
// When the context is canceled valCh will not be drained anymore.
func (c *resolver) Resolve(ctx context.Context, vh directive.ResolverHandler) error {
	// Construct and attach the new controller to the bus.
	config := c.directive.GetExecControllerConfig()
	factory := c.directive.GetExecControllerFactory()

	le := c.controller.le.WithField("config", factory.GetConfigID())
	bus := c.controller.bus

	ci, err := factory.Construct(config, controller.ConstructOpts{
		Logger: le,
	})
	if err != nil {
		return err
	}

	// execute the controller
	var execNextBo time.Duration
	if c.lastErr == nil {
		c.execBackoff.Reset()
	} else {
		execNextBo = c.execBackoff.NextBackOff()
		if execNextBo == -1 {
			return errors.Wrap(c.lastErr, "backoff timeout exceeded")
		}
	}

	if execNextBo != 0 {
		le.
			WithField("backoff-duration", execNextBo.String()).
			Debug("backing off before controller re-start")
		boTimer := time.NewTimer(execNextBo)
		defer boTimer.Stop()

		// emit the value
		now := time.Now()
		vid, ok := vh.AddValue(NewExecControllerValue(
			now,
			now.Add(execNextBo),
			nil,
			c.lastErr,
		))

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-boTimer.C:
			if ok {
				vh.RemoveValue(vid)
			}
		}
	}
	c.lastErr = nil

	go func() {
		// type assertion
		t1 := time.Now()

		// emit the value
		vid, ok := vh.AddValue(NewExecControllerValue(
			t1,
			time.Time{},
			ci,
			nil,
		))
		if !ok {
			// value rejected, drop the controller on the floor.
			ci.Close()
			return
		}

		le.Debug("starting controller")
		err := bus.ExecuteController(c.ctx, ci)
		le := le.WithField("exec-time", time.Since(t1).String())
		if err != nil && err != context.Canceled {
			le.WithError(err).Warn("controller exited with error")
			c.lastErr = err
		} else {
			le.Debug("controller exited normally")
			err = nil // if context canceled, leave value there
		}
		if err != nil {
			vh.RemoveValue(vid)
		}
	}()

	return nil
}

// _ is a type assertion
var _ directive.Resolver = ((*resolver)(nil))
