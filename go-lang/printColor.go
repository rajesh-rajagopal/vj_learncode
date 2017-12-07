package main 
import (
  "fmt"
  "text/tabwriter"
  "strings"
  "bytes"
  "github.com/megamsys/libgo/cmd"
)


func main() {
     fmt.Println(showURL("POST","localhost:9005","/login"))
}
func showURL(m,h,url string) string {
	w := new(tabwriter.Writer)
	var b bytes.Buffer
	w.Init(&b, 0, 8, 0, '\t', 0)
	b.Write([]byte("\t[" + cmd.Colorfy(m, "cyan", "", "bold") + "]\t" + cmd.Colorfy("http://"+ h + url, "white", "", "bold") + ".\n"))
	fmt.Fprintln(w)
	w.Flush()
	return strings.TrimSpace(b.String())
}
