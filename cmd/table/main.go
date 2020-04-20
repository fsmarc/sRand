// sRand/table is a small command-line-programm to generate distributance-tables of
// Charsets for further analysis.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sRand/rand"
	"strings"
	"text/tabwriter"
	"time"
)

var n int
var charset string
var seed int64
var verbose bool
var out string

func init() {
	flag.IntVar(&n, "n", 1000, "Number of Iterations used to estimate distributance.")
	flag.StringVar(&charset, "charset", rand.DefaultCharset, "Charset used for rand.")
	flag.Int64Var(&seed, "seed", time.Now().UnixNano(), "Seed used for rand. Default to time.Now().UnixNano()")
	flag.BoolVar(&verbose, "v", false, "Print Table on os.Stdout.")
	flag.StringVar(&out, "json", "", "Save measurement as json-File.")
}

func main() {
	flag.Parse()

	rnd := rand.New(rand.NewSource(seed))

	data := &Data{n, seed, allocate(n, rnd)}

	log.Printf("Seed: %d\n", seed)
	log.Printf("Total: %d\n", n)

	if verbose {
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 0, '\t', 0)
		data.asTable(w, "\t")
	}

	if out != "" {
		f, err := os.Create(out)
		if err != nil {
			log.Fatalf("Creating File: %v\n", err)
		}

		defer f.Close()

		enc := json.NewEncoder(f)
		enc.Encode(data)
	}

}

type Value struct {
	Char string
	Occ  int
	Perc float32
}

type Data struct {
	Total  int
	Seed   int64
	Values []Value
}

func (d *Data) asTable(tw *tabwriter.Writer, delemiter string) (n int) {
	defer tw.Flush()

	fmt.Fprintf(tw, "Char%sOcc%sPerc\n", delemiter, delemiter)
	n++

	for _, value := range d.Values {
		fmt.Fprintf(tw, "%s%s%d%s%3.4f\n", value.Char, delemiter, value.Occ, delemiter, value.Perc)
		n++
	}

	return n
}

func allocate(n int, r *rand.Rand) []Value {
	//initialize measurement slice
	dist := make([]Value, len(r.Charset))

	for i, v := range dist {
		v.Char = string(r.Charset[i])
		dist[i] = v
	}

	for i := 0; i < n; i++ {
		c := string(r.Char())
		j := strings.Index(r.Charset, c)
		v := dist[j]
		v.Occ++
		v.Perc = (float32(v.Occ) / float32(n)) * 100
		dist[j] = v
	}

	return dist
}
