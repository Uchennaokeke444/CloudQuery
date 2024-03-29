package spec

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cloudquery/filetypes/v4"
	"github.com/cloudquery/plugin-sdk/v4/configtype"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/require"
)

func TestSpec_SetDefaults(t *testing.T) {
	dur30 := configtype.NewDuration(30 * time.Second)

	cases := []struct {
		Give Spec
		Want Spec
	}{

		{
			Give: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ", "}}},
			Want: Spec{Path: "test/path/{{TABLE}}.json", FileSpec: filetypes.FileSpec{Format: "json", FormatSpec: map[string]any{"delimiter": ", "}},
				BatchSize: ptr(int64(10000)), BatchSizeBytes: ptr(int64(50 * 1024 * 1024)), BatchTimeout: &dur30},
		},
	}
	for _, tc := range cases {
		got := tc.Give
		got.SetDefaults()
		if diff := cmp.Diff(tc.Want, got, cmpopts.IgnoreUnexported(filetypes.FileSpec{}, configtype.Duration{})); diff != "" {
			t.Errorf("SetDefaults() mismatch (-want +got):\n%s", diff)
		}
		require.Equal(t, tc.Want.BatchTimeout, got.BatchTimeout)
	}
}

func TestSpec_Validate(t *testing.T) {
	zero, one, dur0 := int64(0), int64(1), configtype.NewDuration(0)
	cases := []struct {
		Give    Spec
		WantErr bool
	}{
		{Give: Spec{Path: "test/path", FileSpec: filetypes.FileSpec{Format: "json"}, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: false, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: false},
		{Give: Spec{Path: "test/path/{{TABLE}}.{{UUID}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true}, // can't have no_rotate and {{UUID}}
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: &zero, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: false},         // norotate with zero batchsize
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true}, WantErr: false},                                                                       // norotate with default batchsize
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: true, BatchSize: &one}, WantErr: true},                                                       // norotate with non zero batchsize
		{Give: Spec{Path: "test/path/{{TABLE}}", FileSpec: filetypes.FileSpec{Format: "json"}, NoRotate: false, BatchSize: &one, BatchSizeBytes: &zero, BatchTimeout: &dur0}, WantErr: true},          // can't have nonzero batch size and no {{UUID}}
	}
	for i, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			err := tc.Give.Validate()
			if tc.WantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSpecUnmarshalJSON(t *testing.T) {
	data := `{
	"format": "csv",
	"format_spec": {
		"skip_header": true,
		"delimiter": "#"
	},
	"path": "abc"
}`
	var s Spec
	require.NoError(t, json.Unmarshal([]byte(data), &s))
	require.Exactly(t, Spec{
		FileSpec: filetypes.FileSpec{
			Format: filetypes.FormatTypeCSV,
			FormatSpec: map[string]any{
				"skip_header": true,
				"delimiter":   "#",
			},
		},
		Path: "abc",
	}, s)
}
