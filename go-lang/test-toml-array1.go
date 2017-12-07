package main

import(
 "github.com/BurntSushi/toml"
 "fmt"
)

type One struct {
 Songs songs
}

type song struct {
  Name     string
  Duration string 
}
type songs struct {
  Song []song
}

func main() {

var tomlBlob =`
[one]
        [[song]]
        name = "Thunder Road"
        duration = "4m49s"

        [[song]]
        name = "Stairway to Heaven"
        duration = "8m03s"
`

var favorites One
if _, err := toml.Decode(tomlBlob, &favorites); err != nil {
  fmt.Println(err)
}

fmt.Println(favorites)


}
