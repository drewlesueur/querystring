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
		{
			map[string]interface{}{
				"A": "a",
				"B": map[string]interface{}{
					"C": "c",
					"D": "d",
				},
			},
			"A=a&B[C]=c&B[D]=d",
		},
		{
			map[string]interface{}{
				"A": "a",
				"B": []interface{}{
					"b1",
					[]interface{}{"b2", "b3"},
					[]string{"b4", "b5"},
					map[string]interface{}{"B6": "b6"},
					map[string]interface{}{"B7": "b7"},
				},
			},
			"A=a&B[0]=b1&B[1][0]=b2&B[1][1]=b3&B[2][0]=b4&B[2][1]=b5&B[3][B6]=b6&B[4][B7]=b7",
		},
		{
			map[string]interface{}{
				"A": "a",
				"B": [5]interface{}{
					"b1",
					[2]interface{}{"b2", "b3"},
					[2]string{"b4", "b5"},
					map[string]interface{}{"B6": "b6"},
					map[string]interface{}{"B7": "b7"},
				},
			},
			"A=a&B[0]=b1&B[1][0]=b2&B[1][1]=b3&B[2][0]=b4&B[2][1]=b5&B[3][B6]=b6&B[4][B7]=b7",
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
