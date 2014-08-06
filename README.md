pprof
=====

Middleware pprof provides pprof service for [Macaron](https://github.com/Unknwon/macaron).

[API Reference](https://gowalker.org/github.com/macaron-contrib/pprof)

### Installation

	go get github.com/macaron-contrib/pprof
	
## Usage

```go
// main.go
import (
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/pprof"
)

func main() {
  	m := macaron.Classic()
  	m.Use(pprof.Pprofer())
	m.Run()
}
```

## License

This project is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text.