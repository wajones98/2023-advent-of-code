use std::collections::HashSet;

fn main() {
    println!("Hello world");
}

#[derive(Debug, PartialEq)]
struct Card {
    id: u32,
    winning_numbers: HashSet<u32>,
    numbers: HashSet<u32>,
}

impl Card {
    fn matching_numbers(self) -> Vec<u32> {
       self.winning_numbers.intersection(&self.numbers).cloned().collect()  
    }
}

fn parse_line(line: &str) -> Card { 
    let parts: Vec<&str> = line.split(":").collect();
    
    let id = parts[0].split_whitespace().collect::<Vec<&str>>()[1];
    let id: u32 = id.parse().unwrap();

    let parts: Vec<&str> = parts[1].split("|").collect();

    let winning_numbers: HashSet<u32> = parts[0].split_whitespace().map(|number| number.parse().unwrap()).collect();
    let numbers: HashSet<u32> = parts[1].split_whitespace().map(|number| number.parse().unwrap()).collect();
    
    Card {
        id,
        winning_numbers,
        numbers,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_parses_line() {
        let line = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53";
        
        let winning_numbers: HashSet<u32> = vec![41, 48, 83, 86, 17].into_iter().collect(); 
        let numbers: HashSet<u32> = vec![83, 86, 6, 31, 17, 9, 48, 53].into_iter().collect(); 

        let expected = Card {
            id: 1,
            winning_numbers,
            numbers,
        };

        let result = parse_line(line);

        assert_eq!(expected, result);
    }

    #[test]
    fn it_gets_matching_numbers() {
        let winning_numbers: HashSet<u32> = vec![41, 48, 83, 86, 17].into_iter().collect(); 
        let numbers: HashSet<u32> = vec![83, 86, 6, 31, 17, 9, 48, 53].into_iter().collect(); 

        let card = Card {
            id: 1,
            winning_numbers,
            numbers,
        };
        
        let expected = vec![48, 83, 17, 86].sort();
        let result = card.matching_numbers().sort();

        assert_eq!(expected, result);
    }
}
