package hclwritex

import (
	"bytes"
	"testing"
)

func TestBlockGet(t *testing.T) {
	cases := []struct {
		name    string
		src     string
		address string
		ok      bool
		want    string
	}{
		{
			name: "simple",
			src: `
a0 = v0
b1 {
  a2 = v2
}

b2 l1 {
}
`,
			address: "b1",
			ok:      true,
			want: `b1 {
  a2 = v2
}
`,
		},
		{
			name:    "no match",
			address: "hoge",
			ok:      true,
			want:    "",
		},
		{
			name:    "empty",
			address: "",
			ok:      false,
			want:    "",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			inStream := bytes.NewBufferString(tc.src)
			outStream := new(bytes.Buffer)
			err := GetBlock(inStream, outStream, "test", tc.address)
			if tc.ok && err != nil {
				t.Fatalf("unexpected err = %s", err)
			}

			got := outStream.String()
			if !tc.ok && err == nil {
				t.Fatalf("expected to return an error, but no error, outStream: \n%s", got)
			}

			if got != tc.want {
				t.Fatalf("got:\n%s\nwant:\n%s", got, tc.want)
			}
		})
	}
}
