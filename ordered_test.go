package orderedjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Map
	}{
		{
			name:  "basecase",
			input: `{"0": null, "1": 0, "2": "s", "3": [null, 0, "string", [], {}], "4": {"0": null, "1": 0, "2": "s", "3": [], "4": {}}}`,
			want: Map{
				{Key: json.RawMessage(`"0"`), Value: json.RawMessage(`null`)},
				{Key: json.RawMessage(`"1"`), Value: json.RawMessage(`0`)},
				{Key: json.RawMessage(`"2"`), Value: json.RawMessage(`"s"`)},
				{Key: json.RawMessage(`"3"`), Value: json.RawMessage(`[null, 0, "string", [], {}]`)},
				{Key: json.RawMessage(`"4"`), Value: json.RawMessage(`{"0": null, "1": 0, "2": "s", "3": [], "4": {}}`)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Map
			err := json.Unmarshal([]byte(tt.input), &got)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
