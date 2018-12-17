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

	cmp := func(props Object) VNode {
		return Text(props.String("text"))
	}

	testcases := []struct {
		name string
		args
		expected VNode
	}{
		{
			name: "Tag Component",
			args: args{
				tag: cmp,
				attrs: Object{
					"text": "test",
				},
				children: []VNode{
					Text("should not be included"),
				},
			},
			expected: &componentNode{
				component: cmp,
				attributes: Object{
					"text": "test",
				},
			},
		},
		{
			name: "Tag string",
			args: args{
				tag: "div",
				attrs: Object{
					"id": "hello",
				},
				children: []VNode{
					Text("should be included"),
				},
			},
			expected: &elementNode{
				name: "div",
				children: []VNode{
					&textNode{
						textContent: "should be included",
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

			if tc.expected.Type() == NodeTypeComponentNode {
				ecn, ok := tc.expected.(*componentNode)
				if !ok {
					t.Errorf("failed to convert expected to componentNode")
				}

				acn, ok := actual.(*componentNode)
				if !ok {
					t.Errorf("failed to convert actual to componentNode")
				}

				if ecn.ComponentPointer() != acn.ComponentPointer() {
					t.Errorf("component pointers not equal")
				}

				// Set components to nil to compare nodes using DeepEqual
				ecn.component = nil
				acn.component = nil
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("not matched expected: %#v, actual: %#v", tc.expected, actual)
			}
		})
	}
}
