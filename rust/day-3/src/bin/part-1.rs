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
    x: usize,
    y: usize,
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

fn parse_schematic_line(y: usize, line: &str) -> Vec<Point> {
    line.chars().enumerate().map(|(x, c)| {
        let point_type: PointType = match c.to_digit(10) {
            Some(digit) => PointType::Digit(digit),
            None => PointType::Symbol(c.to_string()),
        };
        
        Point {
            x,
            y,
            point_type,
        }
    }).collect() 
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
                    point_type: PointType::Digit(4)
                },
                Point {
                    x: 1,
                    y: 0,
                    point_type: PointType::Digit(6)
                },
                Point {
                    x: 2,
                    y: 0,
                    point_type: PointType::Digit(7)
                },
                Point {
                    x: 3,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string())
                },
                Point {
                    x: 4,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string())
                },
                Point {
                    x: 5,
                    y: 0,
                    point_type: PointType::Digit(1)
                },
                Point {
                    x: 6,
                    y: 0,
                    point_type: PointType::Digit(1)
                },
                Point {
                    x: 7,
                    y: 0,
                    point_type: PointType::Digit(4)
                },
                Point {
                    x: 8,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string())
                },
                Point {
                    x: 9,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string())
                },
            ];
        let result = parse_schematic_line(0, line);
        assert_eq!(expected, result);
    }
    
    #[test]
    fn should_identify_symbols() {}

    #[test]
    fn should_identify_numbers () {}

    #[test]
    fn should_calculate_total() {}
}
