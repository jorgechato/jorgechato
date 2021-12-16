### Yo!

This is Jorge, a SRE expert and Green Hat Hacker.

```python
import json
from dataclasses import asdict, dataclass


@dataclass
class Stack:
    languages   : tuple = ("Python", "Go", "JS", "Java")
    platforms   : tuple = ("Kafka", "AWS", "Jenkins")
    misc        : tuple = ("Docker", "K8s", "Terraform")
    ongoing     : tuple = ("AsyncAPI", "RestAPI", "日本語")
    tools       : tuple = ("NeoVim", "Jetbrains")

    def serialize(self):
        return json.dumps(asdict(self), indent=4)


stack = Stack()
print(stack.serialize())
```

---

Seeking my next adventure in :jp: or :kr:
