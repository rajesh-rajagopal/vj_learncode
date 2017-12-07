//You can get the top-level keys of a JSON structure by doing:

package main

import (
    "encoding/json"
    "fmt"
)

// your JSON structure as a byte slice
var j = []byte(`{"foo":1,"bar":2,"baz":[3,4]}`)

func main() {

    // a map container to decode the JSON structure into
    c := make(map[string]interface{})

    // unmarschal JSON
    e := json.Unmarshal(j, &c)

    // panic on error
    if e != nil {
        panic(e)
    }

    // a string slice to hold the keys
    k := make([]string, len(c))

    // iteration counter
    i := 0

    // copy c's keys into k
    for s, _ := range c {
        k[i] = s
        i++
    }

    // output result to STDOUT
    fmt.Printf("%#v\n", k)

}
