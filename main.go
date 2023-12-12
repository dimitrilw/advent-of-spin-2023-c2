package main

import (
	"encoding/json"
	"net/http"
	"sort"

	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
)

type Data struct {
	Kids     []int `json:"kids"`
	Weight   []int `json:"weight"`
	Capacity int   `json:"capacity"`
}
func (d Data) Len() int { return len(d.Kids) }
func (d Data) Less(i, j int) bool {
	// sort by weight per kid (ascending)
	return (d.Weight[i] / d.Kids[i]) < (d.Weight[j] / d.Kids[j])
}
func (d Data) Swap(i, j int) {
	// keep relative order of elements
	d.Kids[i], d.Kids[j] = d.Kids[j], d.Kids[i]
	d.Weight[i], d.Weight[j] = d.Weight[j], d.Weight[i]
}

type Payload struct {
	Kids int `json:"kids"`
}

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		var d Data

		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		p := Payload{ max_kids(d) }

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	})
}

func main() {}

func max_kids(d Data) int {
	res := 0

	sort.Sort(d)

	for i := 0; i < d.Len(); i++ {
		if d.Weight[i] <= d.Capacity {
			d.Capacity -= d.Weight[i]
			res += d.Kids[i]
		}
	}

	return res
}