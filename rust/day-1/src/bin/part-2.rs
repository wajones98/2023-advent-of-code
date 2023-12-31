use std::{
    fs::File,
    io::{prelude::*, BufReader},
    path::Path,
};

const DIGIT_WORDS: [&str; 9] = [
    "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
];

fn main() {
    let lines = lines_from_file("./input.txt");
    let result = process_document(lines);
    println!("{}", result)
}

fn lines_from_file(filename: impl AsRef<Path>) -> Vec<String> {
    let file = File::open(filename).expect("File doesn't exist");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("Failed to parse line"))
        .collect()
}

fn process_document(lines: Vec<String>) -> u64 {
    lines
        .into_iter()
        .fold(0, |total, line| total + extract_digits(line))
}

fn extract_digits(text: String) -> u64 {
    let mut current_string: String = String::from("");
    let mut digits: Vec<String> = vec![];

    for c in text.chars() {
        match c.to_digit(10) {
            Some(_) => {
                digits.push(c.to_string());
            }
            None => {}
        }

        current_string = current_string + &c.to_string();
        for (i, word) in DIGIT_WORDS.iter().enumerate() {
            if current_string.contains(word) {
                let result = (i + 1).to_string();
                current_string = current_string.replace(&word[..&word.len() - 1], &result);
                digits.push(result);
            }
        }
    }

    let result = format!("{}{}", digits[0], digits[digits.len() - 1]);
    result.parse().expect("Expected result to be integer")
}

#[cfg(test)]
mod tests {
    use crate::{extract_digits, process_document};

    struct Test {
        input: String,
        output: u64,
    }

    #[test]
    fn should_find_correct_digits() {
        let tests = vec![
            Test {
                input: String::from("1abc2"),
                output: 12,
            },
            Test {
                input: String::from("pqr3stu8vwx"),
                output: 38,
            },
            Test {
                input: String::from("a1b2c3d4e5f"),
                output: 15,
            },
            Test {
                input: String::from("treb7uchet"),
                output: 77,
            },
            Test {
                input: String::from("two1nine"),
                output: 29,
            },
            Test {
                input: String::from("eightwothree"),
                output: 83,
            },
            Test {
                input: String::from("abcone2threexyz"),
                output: 13,
            },
            Test {
                input: String::from("xtwone3four"),
                output: 24,
            },
            Test {
                input: String::from("4nineeightseven2"),
                output: 42,
            },
            Test {
                input: String::from("zoneight234"),
                output: 14,
            },
            Test {
                input: String::from("7pqrstsixteen"),
                output: 76,
            },
        ];

        for test in tests {
            assert_eq!(extract_digits(test.input), test.output);
        }
    }

    #[test]
    fn correctly_calculates_total() {
        let input = vec![
            String::from("1abc2"),
            String::from("pqr3stu8vwx"),
            String::from("a1b2c3d4e5f"),
            String::from("treb7uchet"),
        ];

        let output = 142;

        assert_eq!(process_document(input), output);
    }
}
