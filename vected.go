// Package vected is a component based frontend framework for golang. This
// framework delivers high performance and sleek ui's, that works both on the
// serverside and the frontend.
//
// Go templates are used as the main templating system. The framework is
// inspired by react, especially preact which I used to learn more about how
// react works.
//
// Also, this borrows from vue js. The templates are just normal go templates so
// no need to learn a different syntax.
//
// The user intrface styles is ant design
// see https://github.com/ant-design/ant-design to learn more about ant design.
package vected

import (
	"context"

	"github.com/gernest/vected/prop"
	"github.com/gernest/vected/state"
	"github.com/gernest/vected/vdom"
)

// RenderMode is a flag determining how a component is rendered.
type RenderMode uint

//supported render mode
const (
	No RenderMode = iota
	Force
	Sync
	Async
)

// Component is an interface which defines a unit of user interface.
type Component interface {

	// New must return an initialized component. This acts as a constructor, the
	// props passed to the component from the parent are passed as arguments,
	//
	// Initializing state should happen here.
	New(prop.Props) (Component, error)

	// Template this is the vected template that is rendered by the component.
	Template() string
	Render(context.Context, prop.Props, state.State) *vdom.Node
	core() *Core
}

// Core is th base struct that every struct that wants to implement Component
// interface must embed.
//
// This is used to make Props available to the component.
type Core struct {
	props           prop.Props
	state           state.State
	prevProps       prop.Props
	prevState       state.State
	disable         bool
	renderCallbacks []func()
	context         context.Context
	prevContext     context.Context
	component       Component
	base            bool
	nextBase        bool
	dirty           bool
	key             prop.NullString

	// This is a callback used to receive instance of Component or the Dom element.
	// after they have been mounted.
	ref func(interface{})
}

func (c *Core) core() *Core { return c }

// SetState updates component state and schedule re rendering.
func (c *Core) SetState(newState state.State, callback ...func()) {
	prev := c.prevState
	c.prevState = newState
	c.state = state.Merge(prev, newState)
	if len(callback) > 0 {
		c.renderCallbacks = append(c.renderCallbacks, callback...)
	}
	//TODO enqueue this for re rendering.
}

// InitState is an interface for exposing initial state.
// Component should implement this interface if they want to set initial state
// when the component is first created before being rendered.
type InitState interface {
	InitState() state.State
}

// InitProps is an interface for exposing default props. This will be merged
// with other props before being sent to render.
type InitProps interface {
	InitProps() prop.Props
}

// WillMount is an interface defining a callback which is invoked before the
// component is mounted on the dom.
type WillMount interface {
	ComponentWillMount()
}

// DidMount is an interface defining a callback that is invoked after the
// component has been mounted to the dom.
type DidMount interface {
	ComponentDidMount(context.Context, prop.Props, state.State)
}

// WillUnmount is an interface defining a callback that is invoked prior to
// removal of the rendered component from the dom.
type WillUnmount interface {
	ComponentWillUnmount(context.Context, prop.Props, state.State)
}

// WillReceiveProps is an interface defining a callback that will be called with
// the new props before they are accepted and passed to be rendered.
type WillReceiveProps interface {
	ComponentWillReceiveProps(context.Context, prop.Props)
}

// ShouldUpdate is an interface defining callback that is called before render
// determine if re render is necessary.
type ShouldUpdate interface {
	// If this returns false then re rendering for the component is skipped.
	ShouldComponentUpdate(context.Context, prop.Props, state.State) bool
}

// WillUpdate is an interface defining a callback that is called before rendering
type WillUpdate interface {
	// If returned props are not nil, then it will be merged with nextprops then
	// passed to render for rendering.
	ComponentWillUpdate(context.Context, prop.Props, state.State) prop.Props
}

// DidUpdate defines a callback that is invoked after rendering.
type DidUpdate interface {
	ComponentDidUpdate()
}

// DerivedState is an interface which can be used to derive state from props.
type DerivedState interface {
	DeriveState(prop.Props, state.State) state.State
}

// SetProps sets cmp props and possibly re renders
func SetProps(ctx context.Context, cmp Component, props prop.Props, state state.State, mode RenderMode, mountAll bool) {
	core := cmp.core()
	if core.disable {
		return
	}
	ref := props["ref"]
	if fn, ok := ref.(func(interface{})); ok {
		core.ref = fn
	}
	core.key = props.String("key")
	delete(props, "key")
	delete(props, "ref")
	_, ok := cmp.(DerivedState)
	if !ok {
		if !core.base || mountAll {
			if m, ok := cmp.(WillMount); ok {
				m.ComponentWillMount()
			}
		} else if m, ok := cmp.(WillReceiveProps); ok {
			m.ComponentWillReceiveProps(ctx, props)
		}
	}
	if core.prevProps == nil {
		core.prevProps = core.props
	}
	core.props = props
	core.disable = false
	if mode != No {
		if mode == Sync {
			// TODO render component
		} else {
			//TODO enqueue render
		}
	}
	if core.ref != nil {
		core.ref(cmp)
	}
}
