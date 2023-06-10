### Hi there!

This is Jorge, a Software Engineer, code lover and Grey Hat Hacker.

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
		Stack{Category: "Languages", Items: []string{"GO", "Python", "TS", "Java"}},
		Stack{Category: "Platforms", Items: []string{"Kafka", "AWS", "Jenkins"}},
		Stack{Category: "Misc", Items: []string{"K8s", "Terraform", "Docker"}},
		Stack{Category: "Ongoing", Items: []string{"Monads", "RX", "gRPC", "日本語"}},
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


Seeking my next adventure in :jp: or :eu:
