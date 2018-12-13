package hyperscript

import (
	"reflect"
	"testing"
)

func TestH(t *testing.T) {
	type args struct {
		tag      interface{}
		attrs    Object
		children []VNode
	}

	testcases := []struct {
		name string
		args
		expected VNode
	}{
		{
			"Tag StatelessComponent",
			args{
				StatelessComponent(func(props Object) VNode {
					return Text(props.String("text"))
				}),
				Object{
					"text": "test",
				},
				[]VNode{
					Text("should not be included"),
				},
			},
			Text("test"),
		},
		{
			"Tag func(Object) VNode",
			args{
				func(props Object) VNode {
					return Text(props.String("text"))
				},
				Object{
					"text": "test",
				},
				[]VNode{
					Text("should not be included"),
				},
			},
			Text("test"),
		},
		{
			"Tag string",
			args{
				"div",
				Object{
					"id": "hello",
				},
				[]VNode{
					Text("should be included"),
				},
			},
			&Element{
				NodeName: "div",
				Children: []VNode{
					Text("should be included"),
				},
				Attributes: Object{
					"id": "hello",
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := H(tc.tag, tc.attrs, tc.children...)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("not matched expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}
}
