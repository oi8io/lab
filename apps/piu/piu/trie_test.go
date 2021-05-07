package piu

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_node_insert(t *testing.T) {
	type fields struct {
		pattern  string
		part     string
		children []*node
		isWild   bool
	}
	type args struct {
		pattern string
		parts   []string
		height  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				pattern:  tt.fields.pattern,
				part:     tt.fields.part,
				children: tt.fields.children,
				isWild:   tt.fields.isWild,
			}
			fmt.Println(n)
		})
	}
}

func Test_node_matchChild(t *testing.T) {
	type fields struct {
		pattern  string
		part     string
		children []*node
		isWild   bool
	}
	type args struct {
		part string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				pattern:  tt.fields.pattern,
				part:     tt.fields.part,
				children: tt.fields.children,
				isWild:   tt.fields.isWild,
			}
			if got := n.matchChild(tt.args.part); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_matchChildren(t *testing.T) {
	type fields struct {
		pattern  string
		part     string
		children []*node
		isWild   bool
	}
	type args struct {
		part string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				pattern:  tt.fields.pattern,
				part:     tt.fields.part,
				children: tt.fields.children,
				isWild:   tt.fields.isWild,
			}
			if got := n.matchChildren(tt.args.part); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchChildren() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_search(t *testing.T) {
	type fields struct {
		pattern  string
		part     string
		children []*node
		isWild   bool
	}
	type args struct {
		parts  []string
		height int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				pattern:  tt.fields.pattern,
				part:     tt.fields.part,
				children: tt.fields.children,
				isWild:   tt.fields.isWild,
			}
			if got := n.search(tt.args.parts, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
