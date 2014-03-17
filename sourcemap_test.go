package sourcemap

import (
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	file := `{
    "version": 3,
    "file": "min.js",
    "names": ["bar", "baz", "n"],
    "sources": ["one.js", "two.js"],
    "sourceRoot": "/the/root",
    "mappings": "CAAC,IAAI,IAAM,SAAUA,GAClB,OAAOC,IAAID;CCDb,IAAI,IAAM,SAAUE,GAClB,OAAOA"
  }`
	m, err := Read(strings.NewReader(file))
	if err != nil {
		t.Fatal(err)
	}
	if m.File != "min.js" || m.SourceRoot != "/the/root" || len(m.Sources) != 2 || m.Sources[0] != "one.js" || len(m.Names) != 3 || m.Names[0] != "bar" {
		t.Error(m)
	}
	mappings := m.DecodedMappings()
	if len(mappings) != 13 {
		t.Error(m)
	}
	assertMapping := func(got, expected *Mapping) {
		if got.OriginalFile != expected.OriginalFile || got.OriginalLine != expected.OriginalLine || got.OriginalColumn != expected.OriginalColumn || got.GeneratedLine != expected.GeneratedLine || got.GeneratedColumn != expected.GeneratedColumn {
			t.Errorf("expected %v, got %v", expected, got)
		}
	}
	assertMapping(mappings[0], &Mapping{"/the/root/one.js", 1, 1, 1, 1})
	assertMapping(mappings[1], &Mapping{"/the/root/one.js", 1, 5, 1, 5})
	assertMapping(mappings[2], &Mapping{"/the/root/one.js", 1, 11, 1, 9})
	assertMapping(mappings[3], &Mapping{"/the/root/one.js", 1, 21, 1, 18})
	assertMapping(mappings[4], &Mapping{"/the/root/one.js", 2, 3, 1, 21})
	assertMapping(mappings[5], &Mapping{"/the/root/one.js", 2, 10, 1, 28})
	assertMapping(mappings[6], &Mapping{"/the/root/one.js", 2, 14, 1, 32})
	assertMapping(mappings[7], &Mapping{"/the/root/two.js", 1, 1, 2, 1})
	assertMapping(mappings[8], &Mapping{"/the/root/two.js", 1, 5, 2, 5})
	assertMapping(mappings[9], &Mapping{"/the/root/two.js", 1, 11, 2, 9})
	assertMapping(mappings[10], &Mapping{"/the/root/two.js", 1, 21, 2, 18})
	assertMapping(mappings[11], &Mapping{"/the/root/two.js", 2, 3, 2, 21})
	assertMapping(mappings[12], &Mapping{"/the/root/two.js", 2, 10, 2, 28})
}
