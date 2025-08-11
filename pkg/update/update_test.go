package update

import "testing"

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"0.1.0", "0.1.0", 0},
		{"0.2.0", "0.1.0", -1},
		{"0.1.0", "0.2.0", 1},
		{"1.0.0", "0.9.9", -1},
		{"0.10.0", "0.2.5", -1},
		{"0.2", "0.2.1", 1},
	}
	for _, tt := range tests {
		got := compareVersions(tt.a, tt.b)
		if got != tt.want {
			t.Fatalf("compareVersions(%s,%s)=%d want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestArchiveSuffix(t *testing.T) {
	if got := archiveSuffix("linux", "amd64"); got == "" {
		t.Fatal("expected suffix for linux/amd64")
	}
	if got := archiveSuffix("windows", "arm64"); got == "" {
		t.Fatal("expected suffix for windows/arm64")
	}
	if got := archiveSuffix("plan9", "amd64"); got != "" {
		t.Fatal("expected empty suffix for unsupported os")
	}
}
