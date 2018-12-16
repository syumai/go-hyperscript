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
			&textNode{
				Node: &Node{
					nodeName: "test",
				},
				textContent: "test",
			},
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
			&textNode{
				Node: &Node{
					nodeName: "test",
				},
				textContent: "test",
			},
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
			&elementNode{
				Node: &Node{
					nodeName: "div",
					children: []VNode{
						&textNode{
							Node: &Node{
								nodeName: "should be included",
							},
							textContent: "should be included",
						},
					},
				},
				attributes: Object{
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
