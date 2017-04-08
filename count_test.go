package syllables

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_CountIn(t *testing.T) {
	data, err := ioutil.ReadFile(filepath.Join("testdata", "syllables.json"))
	if err != nil {
		panic(err)
	}

	tests := make(map[string]int)
	json.Unmarshal(data, &tests)
	if err != nil {
		panic(err)
	}

	failures := 0
	for word, count := range tests {
		got := In(word)
		if got != count {
			failures++
			t.Errorf("syllables.In(%q) == %v, expected %v", word, got, count)
		}
	}
	fmt.Printf("Total failures: %d\n", failures)
}

func Test_CountInBytes(t *testing.T) {
	cases := []struct {
		want int
		in   []byte
	}{
		// Original test cases
		{3, []byte("syllable")},
		{3, []byte("unicorn")},
		{1, []byte("hi")},
		{2, []byte("hihi")},
		{1, []byte("mmmmmmmmmmmmmmmm")},
		{2, []byte("hoopty")},

		// Additional test cas)es
		{6, []byte("syllables in this phrase")},
		{31, []byte("Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.")},
		{108, []byte("If you have just untarred a binary Go distribution, you need to set the environment variable $GOROOT to the full path of the go directory (the one containing this file). You can omit the variable if you unpack it into /usr/local/go, or if you rebuild from sources by running all.bash (see doc/install-source.html). You should also add the Go binary directory $GOROOT/bin to your shell's path.")},
		{2, []byte("template")},
		{4, []byte("abalone")},
		{5, []byte("aborigine")},
		{3, []byte("simile")},
		{4, []byte("facsimile")},
		{3, []byte("syncope")},
		{0, []byte("")},
		{1, []byte("yo")},
		{3, []byte("kilogram")},
		{112, []byte("If you have just untarred a binary Go distribution, you need to set the environment variable $GOROOT to the full path of the go directory (the one containing this file). You can omit the variable if you unpack it into /usr/local/go, or if you rebuild from sources by running all.bash (see doc/install-source.html). You should also add the Go binary directory $GOROOT/bin to your Scientology shell's path. ")},
		{115, []byte("If you have just untarred a binary Go distribution, you need to set the environment variable $GOROOT to the full path of the go directory (the one containing this file). You can omit the variable if you unpack it into /usr/local/go, or if you rebuild from sources by running all.bash (see doc/install-source.html). You should also add the Go binary directory $GOROOT/bin to your Scientology ology shell's path. ")},
		{11, []byte("The quick brown fox jumps over the lazy dog.")},
	}

	for _, c := range cases {
		got := InBytes(c.in)
		if got != c.want {
			t.Errorf("syllables.In(%q) == %v, expected %v", c.in, got, c.want)
		}
	}
}
