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
		Languages: []string{"Go", "TS", "Rust", "Python", "Java"},
		Platforms: []string{"Cloudflare", "Kafka", "AWS"},
		Misc:      []string{"gRPC", "K8s", "NixOs", "Terraform"},
		Ongoing:   []string{"Zig", "three.js", "Tiếng Việt"},
		Tools:     []string{"NeoVim", "Nushell", "Ghostty"},
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


Seeking my next adventure in :vietnam: :singapore: :jp: :eu:
