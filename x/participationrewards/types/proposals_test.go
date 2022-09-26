package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddProtocolDataProposal_ValidateBasic(t *testing.T) {
	type fields struct {
		Title       string
		Description string
		Protocol    string
		Type        string
		Key         string
		Data        json.RawMessage
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"blank",
			fields{},
			true,
		},
		{
			"invalid_protocol",
			fields{
				Title:       "Add Test Protocol",
				Description: "A new protocol for testing protocols",
				Protocol:    "",
				Type:        "",
				Key:         "",
				Data:        nil,
			},
			true,
		},
		{
			"invalid_type",
			fields{
				Title:       "Add Test Protocol",
				Description: "A new protocol for testing protocols",
				Protocol:    "TestProtocol",
				Type:        "",
				Key:         "",
				Data:        nil,
			},
			true,
		},
		{
			"invalid_key",
			fields{
				Title:       "Add Test Protocol",
				Description: "A new protocol for testing protocols",
				Protocol:    "TestProtocol",
				Type:        "TestType",
				Key:         "",
				Data:        nil,
			},
			true,
		},
		{
			"invalid_data",
			fields{
				Title:       "Add Test Protocol",
				Description: "A new protocol for testing protocols",
				Protocol:    "TestProtocol",
				Type:        "TestType",
				Key:         "TestKey",
				Data:        nil,
			},
			true,
		},
		// TODO: add valid case
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := AddProtocolDataProposal{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Protocol:    tt.fields.Protocol,
				Type:        tt.fields.Type,
				Key:         tt.fields.Key,
				Data:        tt.fields.Data,
			}
			err := m.ValidateBasic()
			if tt.wantErr {
				t.Logf("Error:\n%v\n", err)
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}