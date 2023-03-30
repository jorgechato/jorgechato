### Yo!

This is Jorge, a Software Engineer, SRE expert and Grey Hat Hacker.

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
        languages: vec!["Python", "Go", "JS", "Java"],
        platforms: vec!["Kafka", "AWS", "Jenkins"],
        misc: vec!["Docker", "K8s", "Terraform"],
        ongoing: vec!["Rust", "AsyncAPI", "日本語"],
        tools: vec!["NeoVim", "Jetbrains"],
    };

    match serde_json::to_string(&stack) {
        Ok(s) => println!("{}", s),
        Err(e) => panic!("{:?}",e),
    }
}
```


Seeking my next adventure in :jp:
