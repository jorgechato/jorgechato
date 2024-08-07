### Hi there!

This is Jorge, a Senior Software Engineer and code lover.

```go
package main

import (
	"encoding/json"

	"github.com/reactivex/rxgo/v2"
)

type Stack struct {
	Category string   `json:"category"`
	Items    []string `json:"items"`
}

func main() {
	observable := rxgo.Just(
		Stack{Category: "Languages", Items: []string{"GO", "Proto", "Python", "TS"}},
		Stack{Category: "Platforms", Items: []string{"Kafka", "AWS", "Jenkins"}},
		Stack{Category: "Misc", Items: []string{"gRPC", "K8s", "Terraform", "Docker"}},
		Stack{Category: "Ongoing", Items: []string{"Monads", "RX", "日本語"}},
		Stack{Category: "Tools", Items: []string{"NeoVim", "Jetbrains"}},
	)().Marshal(json.Marshal)

	for item := range observable.Observe() {
		if item.Error() {
			panic(item.E)
		}
		println(string(item.V.([]byte)))
	}
}
```

**If you want to know more about me, check:**

- 🗺️ [Where am I Today?](https://whereisjorge.today/)
- ✍️ [What am I doing Today?](https://whatisjorgedoing.today/)
- 📦 [Are my projects alive Today?](https://jorgechato.com/status)


Seeking my next adventure in :eu:
