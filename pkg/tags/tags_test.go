package tags_test

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/gobl/gobl/pkg/tags"
	"testing"
)

func TestNewTags(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		args args
		want tags.Tags
	}{
		{
			name: "can create an empty tags",
			args: args{},
			want: tags.Tags{},
		},
		{
			name: "can create tags with 1 value",
			args: args{tags: []string{"tag1"}},
			want: tags.Tags{"tag1": struct{}{}},
		},
		{
			name: "can create tags with multiple values",
			args: args{tags: []string{"tag1", "tag2", "tag3"}},
			want: tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}, "tag3": struct{}{}},
		},
		{
			name: "can create tags with duplicate values",
			args: args{tags: []string{"tag1", "tag1", "tag1"}},
			want: tags.Tags{"tag1": struct{}{}},
		},
		{
			name: "can create tags with multiple duplicate values",
			args: args{tags: []string{"tag1", "tag1", "tag1", "tag2", "tag2"}},
			want: tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tags.NewTags(tt.args.tags...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTags_Add(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		t    tags.Tags
		args args
		want tags.Tags
	}{
		{
			name: "can add tags to an empty set",
			t:    tags.Tags{},
			args: args{tags: []string{"tag1", "tag2"}},
			want: tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
		},
		{
			name: "can add tags to a set with existing tags",
			t:    tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag2", "tag3"}},
			want: tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}, "tag3": struct{}{}},
		},
		{
			name: "can add duplicate tags to a set",
			t:    tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag1", "tag1", "tag1"}},
			want: tags.Tags{"tag1": struct{}{}},
		},
		{
			name: "can add duplicate tags and new tags to a set",
			t:    tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag1", "tag1", "tag1", "tag2", "tag2"}},
			want: tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.Add(tt.args.tags...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTags_Remove(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		tr   tags.Tags
		args args
		want tags.Tags
	}{
		// TODO: Add test cases.
		{
			name: "can remove tags from an empty set",
			tr:   tags.Tags{},
			args: args{tags: []string{"tag1", "tag2"}},
			want: tags.Tags{},
		},
		{
			name: "can remove tags from a set with existing tags",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{tags: []string{"tag2", "tag3"}},
			want: tags.Tags{"tag1": struct{}{}},
		},
		{
			name: "can remove tags that don't exist from a set",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag2", "tag3"}},
			want: tags.Tags{"tag1": struct{}{}},
		},
		{
			name: "can remove duplicate tags from a set",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag1", "tag1", "tag1"}},
			want: tags.Tags{},
		},
		{
			name: "can remove multiple tags from a set",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{tags: []string{"tag1", "tag2"}},
			want: tags.Tags{},
		},
		{
			name: "can remove multiple duplicate tags from a set",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{tags: []string{"tag1", "tag1", "tag1", "tag2", "tag2"}},
			want: tags.Tags{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Remove(tt.args.tags...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTags_Has(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		tr   tags.Tags
		args args
		want bool
	}{
		{
			name: "can check if a tag exists in an empty set",
			tr:   tags.Tags{},
			args: args{tag: "tag1"},
			want: false,
		},
		{
			name: "can check if a tag exists in a set with the tag",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tag: "tag1"},
			want: true,
		},
		{
			name: "can check if a tag exists in a set without the tag",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tag: "tag2"},
			want: false,
		},
		{
			name: "can check if a tag exists in a set with multiple tags",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{tag: "tag2"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Has(tt.args.tag)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTags_HasAll(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		tr   tags.Tags
		args args
		want bool
	}{
		{
			name: "can check if all tags exist in an empty set",
			tr:   tags.Tags{},
			args: args{tags: []string{"tag1", "tag2"}},
			want: false,
		},
		{
			name: "can check if all tags exist in a set with all the tags",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{tags: []string{"tag1", "tag2"}},
			want: true,
		},
		{
			name: "can check if all tags exist in a set with some of the tags",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag1", "tag2"}},
			want: false,
		},
		{
			name: "can check if all tags exist in a set with none of the tags",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag2", "tag3"}},
			want: false,
		},
		{
			name: "can check if all tags exist in a set with all the tags and more",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}, "tag3": struct{}{}},
			args: args{tags: []string{"tag1", "tag2"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.HasAll(tt.args.tags...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTags_HasAny(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		tr   tags.Tags
		args args
		want bool
	}{
		{
			name: "can check if any tags exist in an empty set",
			tr:   tags.Tags{},
			args: args{tags: []string{"tag1", "tag2"}},
			want: false,
		},
		{
			name: "can check if any tags exist in a set with all the tags",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{tags: []string{"tag1", "tag2"}},
			want: true,
		},
		{
			name: "can check if any tags exist in a set with some of the tags",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag1", "tag2"}},
			want: true,
		},
		{
			name: "can check if any tags exist in a set with none of the tags",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag2", "tag3"}},
			want: false,
		},
		{
			name: "can check if any tags exist in a set with all the tags and more",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}, "tag3": struct{}{}},
			args: args{tags: []string{"tag1", "tag2"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.HasAny(tt.args.tags...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTags_HasNone(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name string
		tr   tags.Tags
		args args
		want bool
	}{
		{
			name: "can check if none of the tags exist in an empty set",
			tr:   tags.Tags{},
			args: args{tags: []string{"tag1", "tag2"}},
			want: true,
		},
		{
			name: "can check if none of the tags exist in a set with all the tags",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{tags: []string{"tag1", "tag2"}},
			want: false,
		},
		{
			name: "can check if none of the tags exist in a set with some of the tags",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag1", "tag2"}},
			want: false,
		},
		{
			name: "can check if none of the tags exist in a set with none of the tags",
			tr:   tags.Tags{"tag1": struct{}{}},
			args: args{tags: []string{"tag2", "tag3"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.HasNone(tt.args.tags...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTags_Equal(t *testing.T) {
	type args struct {
		o tags.Tags
	}
	tests := []struct {
		name string
		tr   tags.Tags
		args args
		want bool
	}{
		{
			name: "can check if two empty sets are equal",
			tr:   tags.Tags{},
			args: args{o: tags.Tags{}},
			want: true,
		},
		{
			name: "can check if two sets with the same tags are equal",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{o: tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}}},
			want: true,
		},
		{
			name: "can check if two sets with different tags are not equal",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{o: tags.Tags{"tag1": struct{}{}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Equal(tt.args.o)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTags_Equals(t *testing.T) {
	type args struct {
		o tags.Tags
	}
	tests := []struct {
		name string
		tr   tags.Tags
		args args
		want bool
	}{
		{
			name: "can check if two empty sets are equal",
			tr:   tags.Tags{},
			args: args{o: tags.Tags{}},
			want: true,
		},
		{
			name: "can check if two sets with the same tags are equal",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{o: tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}}},
			want: true,
		},
		{
			name: "can check if two sets with different tags are not equal",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}},
			args: args{o: tags.Tags{"tag1": struct{}{}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Equals(tt.args.o)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTags_Tags(t *testing.T) {
	tests := []struct {
		name string
		tr   tags.Tags
		want []string
	}{
		{
			name: "can get tags from an empty set",
			tr:   tags.Tags{},
			want: []string{},
		},
		{
			name: "can get tags from a set with 1 tag",
			tr:   tags.Tags{"tag1": struct{}{}},
			want: []string{"tag1"},
		},
		{
			name: "can get tags from a set with multiple tags",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}, "tag3": struct{}{}},
			want: []string{"tag1", "tag2", "tag3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Tags()
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestTags_Copy(t *testing.T) {
	tests := []struct {
		name string
		tr   tags.Tags
		want tags.Tags
	}{
		{
			name: "can copy an empty set",
			tr:   tags.Tags{},
			want: tags.Tags{},
		},
		{
			name: "can copy a set with 1 tag",
			tr:   tags.Tags{"tag1": struct{}{}},
			want: tags.Tags{"tag1": struct{}{}},
		},
		{
			name: "can copy a set with multiple tags",
			tr:   tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}, "tag3": struct{}{}},
			want: tags.Tags{"tag1": struct{}{}, "tag2": struct{}{}, "tag3": struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Copy()
			assert.Equal(t, tt.want, got)
		})
	}
}
