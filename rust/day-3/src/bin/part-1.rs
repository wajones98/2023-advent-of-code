fn main() {
    println!("Hello, world!");
}

#[cfg(test)]
mod tests {
    const TEST_ENGINE_SCHEMATIC: [&str; 10] = [
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
    ];


    #[test]
    fn should_parse_line() {}
    
    #[test]
    fn should_identify_symbols() {}

    #[test]
    fn should_identify_numbers () {}

    #[test]
    fn should_calculate_total() {}
}
