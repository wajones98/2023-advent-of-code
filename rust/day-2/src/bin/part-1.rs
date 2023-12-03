#[derive(PartialEq, Debug)]
struct Set {
    red: u32,
    green: u32,
    blue: u32,
}

#[derive(PartialEq, Debug)]
struct Game {
    id: u32,
    sets_revealed: Vec<Set>
}

impl Game {
    fn new(id: u32, sets_revealed: Vec<Set>) -> Self {
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
                    let mut output = Set{
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

fn validate_set(expected_set: Set, revealed_sets: Vec<Set>) -> bool {
    revealed_sets.into_iter().fold(false, |_, set| {
        set.red <= expected_set.red && set.green <= expected_set.green && set.blue <= expected_set.blue 
    })
}

#[cfg(test)]
mod tests {
    use crate::{Set, validate_set};
    
    const VALID_SET: Set = Set {
        red: 12,
        green: 13,
        blue: 14,
    }; 

    #[test]
    fn should_parse_line() {
        let line = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green";
        let expected_output = crate::Game::new(1, vec![
            Set {
                red: 4,
                green: 0,
                blue: 3,
            },
            Set {
                red: 1,
                green: 2,
                blue: 6,
            },
            Set {
                red: 0,
                green: 2,
                blue: 0,
            }
        ]); 
        let output = crate::Game::parse_string(line);

        assert_eq!(expected_output, output);
    }

    #[test]
    fn should_correctly_determine_valid_set() {
        let sets = vec![
            Set {
                red: 4,
                green: 0,
                blue: 3,
            },
            Set {
                red: 1,
                green: 2,
                blue: 6,
            },
            Set {
                red: 0,
                green: 2,
                blue: 0,
            }
        ];
         
        assert_eq!(true, validate_set(VALID_SET, sets))
    }
}
