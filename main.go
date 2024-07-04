package auwebswaggerui

import (
	"github.com/shurcooL/vfsgen"
	"log"
	"net/http"
)

var fs http.FileSystem = http.Dir("third_party/swagger_ui")

func main() {
	err := vfsgen.Generate(fs, vfsgen.Options{})
	if err != nil {
		log.Fatalln(err)
	}
}
