package deepcopy

import (
	"reflect"
	"strings"
)

// fieldDetail stores field copying detail parsed from a struct field
type fieldDetail struct {
	field    *reflect.StructField
	key      string
	ignored  bool
	required bool

	done         bool
	index        []int
	nestedFields []*fieldDetail
}

// markDone sets the `done` flag of a field detail and all of its nested fields recursively
func (detail *fieldDetail) markDone() {
	detail.done = true
	for _, f := range detail.nestedFields {
		f.markDone()
	}
}

// parseTag parses struct tag for getting copying detail and configuration
func parseTag(detail *fieldDetail) {
	tagValue, ok := detail.field.Tag.Lookup(defaultTagName)
	detail.key = detail.field.Name
	if !ok {
		detail.ignored = !detail.field.IsExported()
		return
	}

	tags := strings.Split(tagValue, ",")
	switch {
	case tags[0] == "-":
		detail.ignored = true
	case tags[0] != "":
		detail.key = tags[0]
	}

	for _, tagOpt := range tags[1:] {
		if tagOpt == "required" && !detail.ignored {
			detail.required = true
		}
	}

	if !detail.required && !detail.field.IsExported() {
		detail.ignored = true
	}
}
