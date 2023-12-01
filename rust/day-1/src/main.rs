fn main() {
    println!("Hello, world!");
}

#[cfg(test)]
mod tests {
    #[test]
    fn finds_two_digits() {}

    #[test]
    fn finds_two_digits_when_multiple_present() {}

    #[test]
    fn finds_two_digits_when_single_digit() {}

    #[test]
    fn correctly_sums_digits() {}
}
