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
        languages: vec!["Python", "Go", "JS", "Java"],
        platforms: vec!["Kafka", "AWS", "Jenkins"],
        misc: vec!["Docker", "K8s", "Terraform"],
        ongoing: vec!["Rust", "RX", "gRPC", "日本語"],
        tools: vec!["NeoVim", "Jetbrains"],
    };

    match serde_json::to_string(&stack) {
        Ok(s) => println!("{}", s),
        Err(e) => panic!("{:?}",e),
    }
}
```

**If you want to know more about me, check:**

<p>
    
[![Where am I Today?](https://img.shields.io/static/v1?label=%F0%9F%97%BA%20WHERE&message=am%20I%20Today%3F&color=2563eb&style=flat)](https://whereisjorge.today/)
[![What am I doing Today?](https://img.shields.io/static/v1?label=%E2%9C%8D%20WHAT&message=am%20I%20doing%20Today%3F&color=059669&style=flat)](https://whatisjorgedoing.today/)

</p>
</br>

Seeking my next adventure in :jp: or :eu:
