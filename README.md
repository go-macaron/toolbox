toolbox
=======

Middleware toolbox provides health chcek, pprof, profile and statistic services for [Macaron](https://github.com/Unknwon/macaron).

[API Reference](https://gowalker.org/github.com/macaron-contrib/toolbox)

### Installation

	go get github.com/macaron-contrib/toolbox
	
## Usage

```go
// main.go
import (
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/toolbox"
)

func main() {
  	m := macaron.Classic()
  	m.Use(toolbox.Toolboxer(m))
	m.Run()
}
```

## License

This project is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text.