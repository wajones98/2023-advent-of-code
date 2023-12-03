fn main() {
    println!("Hello, world!");
}

enum PointType {
    Symbol(String),
    Digit(u32),
}

struct Point {
    x: u32,
    y: u32,
    point_type: PointType,
}

type Schematic = Vec<Vec<Point>>;

fn new_schematic(lines: Vec<&str>) -> Schematic {
    vec![
        vec![
            Point{
                x: 0,
                y: 0,
                point_type: PointType::Symbol(".".to_string()), 
            }
        ]
    ] 
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
