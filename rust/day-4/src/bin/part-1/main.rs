fn main() {
    println!("Hello world");
}

#[derive(Debug, PartialEq)]
struct Card {
    id: u32,
    winning_numbers: Vec<u32>,
    numbers: Vec<u32>,
}

fn parse_line(line: &str) -> Card {
    let parts: Vec<&str> = line.split(":").collect();
    
    let id = parts[0].split_whitespace().collect::<Vec<&str>>()[1];
    let id: u32 = id.parse().unwrap();

    let parts: Vec<&str> = parts[1].split("|").collect();

    let winning_numbers: Vec<u32> = parts[0].split_whitespace().map(|number| number.parse().unwrap()).collect();
    let numbers: Vec<u32> = parts[1].split_whitespace().map(|number| number.parse().unwrap()).collect();
    
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

        let expected = Card {
            id: 1,
            winning_numbers: vec![41, 48, 83, 86, 17],
            numbers: vec![83, 86, 6, 31, 17, 9, 48, 53],
        };

        let result = parse_line(line);

        assert_eq!(expected, result);
    }
}
