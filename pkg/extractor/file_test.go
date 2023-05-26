package extractor

import (
	"strings"
	"testing"
)

type extractFunctionFromFileTest struct {
	name    string
	args    args
	want    []Functions
	wantErr bool
}
type args struct {
	file string
}

func TestExtractFunctionsFromFile(t *testing.T) {

	tests := []extractFunctionFromFileTest{
		{
			name: "test 1",
			args: args{
				file: "./testfile.go.txt",
			},
			want: []Functions{
				{
					File: "./testfile.go.txt",
					Name: "func1",
					Body: `func func1() {
	fmt.Println("asdf")
}
`,
				},
				{
					File: "./testfile.go.txt",
					Name: "func2",
					Body: `func func2() {
	if true {
		fmt.Println("asdf")
	}
}
`,
				},
				{
					File: "./testfile.go.txt",
					Name: "func3",
					Body: `func func3() string {
	if true {
		fmt.Println("asdf")
	}
	return "sadf"
}
`,
				},
				{
					File: "./testfile.go.txt",
					Name: "func4",
					Body: `func func4(filePath string) (string, error) {
	if true {
		fmt.Println("asdf")
	}
	return "sadf", nil
}
`,
				},
				{
					File: "./testfile.go.txt",
					Name: "func5",
					Body: `func (c context.CancelFunc) func5(filePath string) (string, error) {
	if true {
		fmt.Println("asdf")
		if false {
			fmt.Println("{}asdf{}")
		}
	}
	return "sadf", nil
}
`,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractFunctionsFromFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractFunctionsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != 5 {
				t.Errorf("wrong number of args")
			}
			for i, _ := range tt.want {
				gcurr := got[i]
				wcurr := tt.want[i]

				if gcurr.Name != wcurr.Name {
					t.Errorf("names are different for index %d", i)
				}
				if gcurr.File != wcurr.File {
					t.Errorf("files are different for index %d", i)
				}
				if strings.TrimSpace(gcurr.Body) != strings.TrimSpace(wcurr.Body) {
					t.Errorf("bodies are different for index %d", i)

				}
			}

		})
	}
}
