package file

import (
	"reflect"
	"testing"
)

func TestExists(t *testing.T) {
	type args struct {
		fpath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "no exists",
			args: args{
				fpath: "/tmp/file/debug.json",
			},
			want: false,
		},
		{
			name: "exists",
			args: args{
				fpath: "/etc/hosts",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.fpath); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFileChanged(t *testing.T) {
	type args struct {
		src  string
		dest string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "no change",
			args: args{
				src:  "/etc/hosts",
				dest: "/etc/hosts",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "change",
			args: args{
				src:  "/etc/hosts",
				dest: "/etc/profile",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFileChanged(tt.args.src, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFileChanged() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsFileChanged() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDirectory(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "yes",
			args: args{
				path: "/etc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "no",
			args: args{
				path: "/etc/hosts",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "no",
			args: args{
				path: "/tmp/file", // not exists
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDirectory(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursiveFilesLookup(t *testing.T) {
	type args struct {
		root    string
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RecursiveFilesLookup(tt.args.root, tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecursiveFilesLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecursiveFilesLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursiveDirsLookup(t *testing.T) {
	type args struct {
		root    string
		pattern string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RecursiveDirsLookup(tt.args.root, tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecursiveDirsLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecursiveDirsLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
