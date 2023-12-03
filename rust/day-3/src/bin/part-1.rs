fn main() {
    println!("Hello, world!");
}

#[derive(Debug, PartialEq)]
enum PointType {
    Symbol(String),
    Digit(u32),
}

#[derive(Debug, PartialEq)]
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

fn parse_schematic_line(line: &str) -> Vec<Point> {
    vec![]
}

#[cfg(test)]
mod tests {
    use crate::{Schematic, new_schematic, parse_schematic_line, Point, PointType};

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
    fn should_parse_line() {
        let line = TEST_ENGINE_SCHEMATIC[0];
        let expected: Vec<Point> = vec![
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol("4".to_string())
                },
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol("6".to_string())
                },
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol("7".to_string())
                },
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string())
                },
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string())
                },
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol("1".to_string())
                },
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol("1".to_string())
                },
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol("4".to_string())
                },
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string())
                },
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string())
                },
            ];
        let result = parse_schematic_line(line);
        assert_eq!(expected, result);
    }
    
    #[test]
    fn should_identify_symbols() {}

    #[test]
    fn should_identify_numbers () {}

    #[test]
    fn should_calculate_total() {}
}
