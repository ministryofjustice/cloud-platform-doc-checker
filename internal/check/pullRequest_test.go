package check

import "testing"

func TestParsePR(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Expect the PR to be approved",
			args: args{
				fileName: "testdata/pr_approved.txt",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Expect the PR to be rejected",
			args: args{
				fileName: "testdata/pr_rejected.txt",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePR(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParsePR() = %v, want %v", got, tt.want)
			}
		})
	}
}
