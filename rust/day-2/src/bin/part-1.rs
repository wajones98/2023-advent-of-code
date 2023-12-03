#[derive(PartialEq, Debug)]
struct RevealedSet {
    red: u32,
    green: u32,
    blue: u32,
}

#[derive(PartialEq, Debug)]
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
    fn parse_string(line: &str) -> Self {

        let mut game_and_sets = line.split(":");
        let game_id = match game_and_sets.next() {
            Some(game) => match game.split_whitespace().last() {
                Some(id) => id.parse::<u32>().expect("Expected positive number"),
                None => panic!("Id expected"),
            },
            None => panic!("Game expected"),
        }; 
        let sets = match game_and_sets.next() {
            Some(sets) => {
                sets.split(";").map(|revealed_set| {
                    let mut output = RevealedSet{
                        red: 0,
                        green: 0,
                        blue: 0,
                    };
                    for info in revealed_set.split(",") {
                        let mut cube_info = info.split_whitespace();
                        let cube_count: u32 = cube_info.next().expect("Expected number").parse().expect("Expected number");
                        let cube_colour = cube_info.next().expect("Expected colour");
                        
                        if cube_colour == "red" {
                            output.red = cube_count;
                        } else if cube_colour == "green" {
                            output.green = cube_count;
                        } else if cube_colour == "blue" {
                            output.blue = cube_count;
                        }
                    }
                    output
                }).collect()
            }
            None => panic!("Game expected"),
        }; 
        Game::new(game_id, sets) 
    }
}

fn main() {
    println!("Hello world!");
}

#[cfg(test)]
mod tests {
    use crate::RevealedSet;

    #[test]
    fn should_parse_line() {
        let line = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green";
        let expected_output = crate::Game::new(1, vec![
            RevealedSet {
                red: 4,
                green: 0,
                blue: 3,
            },
            RevealedSet {
                red: 1,
                green: 2,
                blue: 6,
            },
            RevealedSet {
                red: 0,
                green: 2,
                blue: 0,
            }
        ]); 
        let output = crate::Game::parse_string(line);

        // assert_eq!(expected_output, output);
    }
}
