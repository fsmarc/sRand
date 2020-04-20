package rand_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"text/tabwriter"
	"time"

	"github.com/fsmarc/sRand/rand"
)

func TestStringn(t *testing.T) {
	l := 10  //length of random strings
	n := 100 //number of strings generated
	gen := make(map[string]bool)
	var duplicates int

	time := time.Now().UnixNano()
	t.Logf("Seed: %d", time)

	rnd := rand.New(rand.NewSource(time))

	for i := 0; i < n; i++ {
		str := rnd.Stringn(l)
		if gen[str] {
			duplicates++
		} else {
			gen[str] = true
		}
	}

	t.Logf("Generated %d strings. Got %d duplicates", n, duplicates)
	if duplicates > 0 {
		t.Errorf("Got %d duplicates", duplicates)
	}
}

func TestDistributance(t *testing.T) {
	buf := new(bytes.Buffer)
	n := 1000

	//initialize rand
	time := time.Now().UnixNano()
	t.Logf("Seed: %d", time)
	rnd := rand.New(rand.NewSource(time))

	//initialize measurement slice
	dist := make([]int, len(rnd.Charset))

	for i := 0; i < n; i++ {
		c := string(rnd.Char())
		j := strings.Index(rnd.Charset, c)
		dist[j]++
	}

	//initialize table
	w := new(tabwriter.Writer)
	w.Init(buf, 0, 8, 0, '\t', 0)

	fmt.Fprintln(w, "Char\tOcc.\t%")

	for i, occ := range dist {
		var perc float32 = float32(occ) / float32(n)
		fmt.Fprintf(w, "%s\t%d\t%3.2f\n", string(rnd.Charset[i]), occ, perc*100)
	}

	w.Flush()

	t.Logf("\n%s", buf)
}

func BenchmarkChar(b *testing.B) {
	time := time.Now().UnixNano()
	b.Logf("Seed: %d", time)

	rnd := rand.New(rand.NewSource(time))

	for i := 0; i < b.N; i++ {
		rnd.Char()
	}
}
