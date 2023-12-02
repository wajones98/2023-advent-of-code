struct RevealedSet {
    red: u32,
    green: u32,
    blue: u32,
}

impl RevealedSet {
    fn new(red: u32, green: u32, blue: u32) -> Self {
       RevealedSet { red, green, blue } 
    }
}

struct Game {
    id: u32,
    sets_revealed: Vec<RevealedSet>
}

impl Game {
    fn new(id: u32, sets_revealed: Vec<RevealedSet>) -> Self {
       Game {
            id,
            sets_revealed
        } 
    }
    
    //TODO: Parse game line from input
    fn parse_string() -> Self {
        let id = 1;
        let sets_revealed = vec![];
        Game::new(id, sets_revealed) 
    }
}

fn main() {
    println!("Hello world!");
}
