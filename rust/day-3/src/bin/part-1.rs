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

fn valid_numbers(schematic: &Schematic) -> Vec<u32> {
    vec![]
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
    use crate::{Schematic, new_schematic, parse_schematic_line, Point, PointType, valid_numbers};

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
    fn should_detect_number_x_adjacent() {
        let line = "617*......";
        let schematic_line = vec![parse_schematic_line(0, line)];
        
        let expected = vec![617];
        let result = valid_numbers(&schematic_line); 
        assert_eq!(expected, result);
    }
    
    #[test]
    fn should_detect_number_y_adjacent() {}
    
    #[test]
    fn should_detect_number_diagonally() {}

    #[test]
    fn should_calculate_total() {}
}
