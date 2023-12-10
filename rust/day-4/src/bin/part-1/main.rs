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
    Card {
        id: 0,
        winning_numbers: vec![],
        numbers: vec![],
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
