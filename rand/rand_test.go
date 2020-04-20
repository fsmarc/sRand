package rand

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"text/tabwriter"
	"time"
)

func TestStringn(t *testing.T) {
	l := 10  //length of random strings
	n := 100 //number of strings generated
	gen := make(map[string]bool)
	var duplicates int

	time := time.Now().UnixNano()
	t.Logf("Seed: %d", time)

	rnd := New(NewSource(time))

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
	n := 1000

	//initialize rand
	time := time.Now().UnixNano()
	t.Logf("Seed: %d", time)
	rnd := New(NewSource(time))

	//initialize measurement slice
	dist := make([]int, len(rnd.Charset))

	for i := 0; i < n; i++ {
		c := string(rnd.Char())
		j := strings.Index(rnd.Charset, c)
		dist[j]++
	}

	//initialize table
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintln(w, "Char\tOcc.\t%")

	for i, occ := range dist {
		var perc float32 = float32(occ) / float32(n)
		fmt.Fprintf(w, "%s\t%d\t%3.2f\n", string(rnd.Charset[i]), occ, perc*100)
	}
}

func BenchmarkChar(b *testing.B) {
	time := time.Now().UnixNano()
	b.Logf("Seed: %d", time)

	rnd := New(NewSource(time))

	for i := 0; i < b.N; i++ {
		rnd.Char()
	}
}

func ExampleNew() {
	rnd := New(NewSource(1))
	fmt.Println(rnd.Stringn(10))
}

func ExampleNewWithCharset() {
	rnd := NewWithCharset(NewSource(1), "Hello World!")
	fmt.Println(rnd.Char())
}
