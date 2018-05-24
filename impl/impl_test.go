package impl

import "testing"

func TestValidGames(t *testing.T) {
	for k, v := range Map() {
		if err := v.Valid(); err != nil {
			t.Error(err)
			continue
		}
		t.Logf("%q: %q", k, v.Build().SVG(true))
		// TODO: validate mutable SVG is setup to not cause problems on the client
		// TODO: validate move slugs are unique
	}
}
