package main

import "testing"

func TestAreaForPackage(t *testing.T) {
	tests := []struct {
		name        string
		pkgPath     string
		stripPrefix string
		areaPrefix  string
		want        string
	}{
		{
			name:        "application package",
			pkgPath:     "github.com/acme/server/internal/view/render",
			stripPrefix: "github.com/acme/server/internal",
			areaPrefix:  "app",
			want:        "app.view.render",
		},
		{
			name:        "library package",
			pkgPath:     "github.com/acme/ble/rx/decoder",
			stripPrefix: "github.com/acme/ble",
			areaPrefix:  "lib.ble",
			want:        "lib.ble.rx.decoder",
		},
		{
			name:        "root package",
			pkgPath:     "github.com/acme/server/internal",
			stripPrefix: "github.com/acme/server/internal",
			areaPrefix:  "app",
			want:        "app",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := areaForPackage(tt.pkgPath, tt.stripPrefix, tt.areaPrefix)
			if err != nil {
				t.Fatalf("areaForPackage returned error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("areaForPackage() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestAreaForPackageRejectsOutsideStripPrefix(t *testing.T) {
	if _, err := areaForPackage("github.com/acme/other/pkg", "github.com/acme/server", "app"); err == nil {
		t.Fatalf("expected strip-prefix error")
	}
}
