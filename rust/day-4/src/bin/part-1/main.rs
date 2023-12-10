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

    fn get_points(self) -> u32 {
        let matching_numbers = self.matching_numbers();

        matching_numbers.iter().fold(0, |acc, _| {
            let mut total = acc;
            if total == 0 {
                total = 1;
            } else {
                total = total * 2;
            }

            total 
        }) 
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

    #[test]
    fn it_gets_correct_points() {
        let winning_numbers: HashSet<u32> = vec![41, 48, 83, 86, 17].into_iter().collect(); 
        let numbers: HashSet<u32> = vec![83, 86, 6, 31, 17, 9, 48, 53].into_iter().collect(); 

        let card = Card {
            id: 1,
            winning_numbers,
            numbers,
        };

        let expected: u32 = 8; 
        let result = card.get_points();

        assert_eq!(expected, result);
    }

    #[test]
    fn it_gets_correct_total() {
        let lines = vec![
            "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
            "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
            "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
            "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
            "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
            "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
        ];

        let expected: u32 = 13;

        let mut result: u32 = 0;

        for line in lines {
            let card = parse_line(line);
            result += card.get_points();
        }

        assert_eq!(expected, result);
    }
}
