fn main() {
    println!("Hello, world!");
}

fn extract_digits(text: String) -> u64 {
    0
}

#[cfg(test)]
mod tests {
    use crate::extract_digits;

    struct Test {
        input: String,
        output: u64,
    }

    #[test]
    fn finds_two_digits() {
        let tests = vec![
            Test {
                input: String::from("1abc2"),
                output: 12,
            },
            Test {
                input: String::from("pqr3stu8vwx"),
                output: 38,
            },
        ];

        for test in tests {
            assert_eq!(extract_digits(test.input), test.output);
        }
    }

    #[test]
    fn finds_two_digits_when_multiple_present() {}

    #[test]
    fn finds_two_digits_when_single_digit() {}

    #[test]
    fn correctly_sums_digits() {}
}
