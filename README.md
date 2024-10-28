### Hi there!

This is Jorge, a Senior Software Engineer and code lover.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	stack := &Stack{
		Languages: []string{"Go", "TS", "Python", "Java"},
		Platforms: []string{"Cloudflare", "Kafka", "AWS"},
		Misc:      []string{"gRPC", "K8s", "Terraform"},
		Ongoing:   []string{"three.js", "NixOs", "日本語"},
		Tools:     []string{"NeoVim", "Jetbrains"},
	}

	jsonData, err := json.Marshal(stack)
	if err != nil {
		log.Fatalf("Failed to marshal stack: %v", err)
	}

	fmt.Println(string(jsonData))
}
```

**If you want to know more about me, check:**

- 🦾 [My Page](https://jorgechato.com/)
- 🗺️ [Where am I Today?](https://whereisjorge.today/)
- ✍️  [What am I doing Today?](https://whatisjorgedoing.today/)
- 📦 [Are my projects alive Today?](https://2023.jorgechato.com/status)


Seeking my next adventure in :eu:
