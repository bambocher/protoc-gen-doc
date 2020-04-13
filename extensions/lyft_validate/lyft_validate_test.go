package extensions_test

import (
	"testing"

	"github.com/envoyproxy/protoc-gen-validate/validate"
	"github.com/golang/protobuf/proto"
	"github.com/bambocher/protoc-gen-doc/extensions"
	. "github.com/bambocher/protoc-gen-doc/extensions/lyft_validate"
	"github.com/stretchr/testify/require"
)

func TestTransform(t *testing.T) {
	fieldRules := &validate.FieldRules{
		Type: &validate.FieldRules_String_{
			String_: &validate.StringRules{
				MinLen: proto.Uint64(1),
			},
		},
	}

	transformed := extensions.Transform(map[string]interface{}{"validate.rules": fieldRules})
	require.NotEmpty(t, transformed)

	rules := transformed["validate.rules"].(ValidateExtension).Rules()
	require.Equal(t, rules, []ValidateRule{
		{Name: "string.min_len", Value: uint64(1)},
	})
}
