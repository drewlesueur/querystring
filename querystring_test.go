package querystring

import "testing"

func TestMap(t *testing.T) {
	tests := []struct {
		in   interface{}
		want string
	}{
		{
			map[string]string{
				"A": "a",
				"B": "b",
			},
			"A=a&B=b",
		},
		{
			map[string]interface{}{
				"A": "a",
				"B": 1,
			},
			"A=a&B=1",
		},
	}

	for i, test := range tests {
		got, err := Stringify(test.in)
		if err != nil {
			t.Errorf("%d. Stringify(%q) returned error: %v", i, test.in, err)
		}
		if test.want != got {
			t.Errorf("expected: %s", test.want)
			t.Errorf("got     : %s", got)
		}

	}
}
