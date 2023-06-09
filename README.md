### Hi there!

This is Jorge, a Software Engineer, code lover and Grey Hat Hacker.

```rust
use serde::Serialize;


#[derive(Serialize)]
struct Stack<'a> {
    languages: Vec<&'a str>,
    platforms: Vec<&'a str>,
    misc: Vec<&'a str>,
    ongoing: Vec<&'a str>,
    tools: Vec<&'a str>,
}

fn main() {
    let stack: Stack = Stack {
        languages: vec!["GO", "Python", "TS", "Java"],
        platforms: vec!["Kafka", "AWS", "Jenkins"],
        misc: vec!["Docker", "K8s", "Terraform"],
        ongoing: vec!["Monads", "RX", "gRPC", "日本語"],
        tools: vec!["NeoVim", "Jetbrains"],
    };

    match serde_json::to_string(&stack) {
        Ok(s) => println!("{}", s),
        Err(e) => panic!("{:?}",e),
    }
}
```

**If you want to know more about me, check:**

- 🗺️ [Where am I Today?](https://whereisjorge.today/)
- ✍️ [What am I doing Today?](https://whatisjorgedoing.today/)
- 📦 [Are my projects alive Today?](https://jorgechato.com/status)


Seeking my next adventure in :jp: or :eu:
