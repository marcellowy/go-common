package tools

import (
	"encoding/xml"
	"testing"
)

func TestXMLMarshalString(t *testing.T) {
	type args struct {
		v interface{}
	}

	type Person struct {
		XMLName xml.Name `xml:"root"`
		Name    string   `xml:"name"`
		Age     int      `xml:"age"`
	}

	type School struct {
		Student string `xml:"student"`
		Teacher string `xml:"teacher"`
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				v: Person{
					Name: "bob",
					Age:  18,
				},
			},
			want: `<root><name>bob</name><age>18</age></root>`,
		},
		{
			name: "test1",
			args: args{
				v: &School{
					Student: "bob",
					Teacher: "Lily",
				},
			},
			want: `<School><student>bob</student><teacher>Lily</teacher></School>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XMLMarshalString(tt.args.v); got != tt.want {
				t.Errorf("XMLMarshalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
