use std::fmt::format;

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

fn valid_numbers(schematic: Schematic) -> Vec<u32> {
    let points_with_symbols: Vec<&Point> = schematic.iter().map(|line| {
        line.iter().filter(|point| -> bool {
            match &point.point_type {
                PointType::Symbol(point_type) => point_type != ".",
                PointType::Digit(_) => false,
            }
        })
    }).flatten().collect();
   
    let mut numbers: Vec<u32> = vec![]; 

    for point in points_with_symbols {
        let mut x_left_coord = point.x;
        let mut left_point: &Point = point;
        let mut left_number: String = "".to_string();

        while x_left_coord > 0 && left_point.point_type != PointType::Symbol(".".to_string()) {
            x_left_coord = left_point.x - 1;

            left_point = &schematic[left_point.y][x_left_coord];   
            if let PointType::Digit(digit) = left_point.point_type {
                left_number = format!("{}{}", digit, left_number);
            }
        }
        let left_number: u32 = left_number.parse::<u32>().expect("Expected valid u32");
        numbers.push(left_number);
    }

    numbers 
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
        let schematic = vec![parse_schematic_line(0, line)];
        
        let expected = vec![617];
        let result = valid_numbers(schematic); 
        assert_eq!(expected, result);
    }
    
    // #[test]
    // fn should_detect_number_y_adjacent() {}
    // 
    // #[test]
    // fn should_detect_number_diagonally() {}
    //
    // #[test]
    // fn should_calculate_total() {}
}
