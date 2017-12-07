package main
import (
"fmt"
	"path/filepath"
	skia "go.skia.org/infra/go/gitinfo"
)

func main() {
	if _, err := skia.Clone("https://github.com/vijaykanthm28/vertice.git -b 1.5", filepath.Join("/home/vijay/test","vertice"), false); err != nil {
		fmt.Println( err)
	}

}
