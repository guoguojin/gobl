package tags

// Tags is a set of tags that can be added to events and used for filtering.
type Tags map[string]struct{}

// NewTags creates a new set of tags.
func NewTags(tags ...string) Tags {
	t := make(Tags)
	return t.Add(tags...)
}

// Add adds tags to the set.
func (t Tags) Add(tags ...string) Tags {
	for _, tag := range tags {
		t[tag] = struct{}{}
	}
	return t
}

// Remove removes tags from the set.
func (t Tags) Remove(tags ...string) Tags {
	for _, tag := range tags {
		delete(t, tag)
	}
	return t
}

// Has returns true if the set contains the tag.
func (t Tags) Has(tag string) bool {
	_, ok := t[tag]
	return ok
}

// HasAll returns true if the set contains all the tags.
func (t Tags) HasAll(tags ...string) bool {
	for _, tag := range tags {
		if !t.Has(tag) {
			return false
		}
	}
	return true
}

// HasAny returns true if the set contains any of the tags.
func (t Tags) HasAny(tags ...string) bool {
	for _, tag := range tags {
		if t.Has(tag) {
			return true
		}
	}
	return false
}

// HasNone returns true if the set contains none of the tags.
func (t Tags) HasNone(tags ...string) bool {
	return !t.HasAny(tags...)
}

// Equal returns true if the sets contain the same tags.
func (t Tags) Equal(o Tags) bool {
	if len(t) != len(o) {
		return false
	}
	for tag := range t {
		if !o.Has(tag) {
			return false
		}
	}
	return true
}

// Equals returns true if the sets contain the same tags.
func (t Tags) Equals(o Tags) bool {
	return t.Equal(o)
}

// Tags returns a slice containing the tags in the set.
func (t Tags) Tags() []string {
	tags := make([]string, 0, len(t))
	for tag := range t {
		tags = append(tags, tag)
	}
	return tags
}

// Copy returns a copy of the set.
func (t Tags) Copy() Tags {
	return NewTags(t.Tags()...)
}
