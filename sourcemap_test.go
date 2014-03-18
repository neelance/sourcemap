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
		if got.GeneratedLine != expected.GeneratedLine || got.GeneratedColumn != expected.GeneratedColumn || got.OriginalFile != expected.OriginalFile || got.OriginalLine != expected.OriginalLine || got.OriginalColumn != expected.OriginalColumn || got.OriginalName != expected.OriginalName {
			t.Errorf("expected %v, got %v", expected, got)
		}
	}
	assertMapping(mappings[0], &Mapping{1, 1, "/the/root/one.js", 1, 1, ""})
	assertMapping(mappings[1], &Mapping{1, 5, "/the/root/one.js", 1, 5, ""})
	assertMapping(mappings[2], &Mapping{1, 9, "/the/root/one.js", 1, 11, ""})
	assertMapping(mappings[3], &Mapping{1, 18, "/the/root/one.js", 1, 21, "bar"})
	assertMapping(mappings[4], &Mapping{1, 21, "/the/root/one.js", 2, 3, ""})
	assertMapping(mappings[5], &Mapping{1, 28, "/the/root/one.js", 2, 10, "baz"})
	assertMapping(mappings[6], &Mapping{1, 32, "/the/root/one.js", 2, 14, "bar"})
	assertMapping(mappings[7], &Mapping{2, 1, "/the/root/two.js", 1, 1, ""})
	assertMapping(mappings[8], &Mapping{2, 5, "/the/root/two.js", 1, 5, ""})
	assertMapping(mappings[9], &Mapping{2, 9, "/the/root/two.js", 1, 11, ""})
	assertMapping(mappings[10], &Mapping{2, 18, "/the/root/two.js", 1, 21, "n"})
	assertMapping(mappings[11], &Mapping{2, 21, "/the/root/two.js", 2, 3, ""})
	assertMapping(mappings[12], &Mapping{2, 28, "/the/root/two.js", 2, 10, "n"})
}
