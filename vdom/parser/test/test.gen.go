// Code generated by vected DO NOT EDIT.
package test

import (
	"context"
	"github.com/gernest/vected/props"
	"github.com/gernest/vected/state"
	"github.com/gernest/vected/vdom"
)

// Render implements vected.Renderer interface.
func (h Hello) Render(ctx context.Context, props props.Props, state state.State) *vdom.Node {
	return &vdom.Node{
		Type: vdom.ElementNode,
		Data: "div",
		Attr: []vdom.Attribute{
			{Key: "classname", Val: props["classNames"]},
		},
	}
}
