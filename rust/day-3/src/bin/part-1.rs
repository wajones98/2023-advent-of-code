use std::{num::ParseIntError, ops::Add};

enum Direction {
    Left = -1,
    Right = 1,
}

impl Add<Direction> for u32 {
    type Output = u32;

    fn add(self, rhs: Direction) -> Self::Output {
        match rhs {
            Direction::Left => {
                if self > rhs as u32 {
                    self - (rhs as u32)
                } else {
                    0 // To prevent underflow, you might want to handle this case accordingly.
                }
            }
            Direction::Right => self + (rhs as u32),
        }
    }
}

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
        if left_number != "".to_string() {
            let left_number: u32 = left_number.parse::<u32>().expect("Expected valid u32");
            numbers.push(left_number);
        }

        let mut x_right_coord = point.x;
        let mut right_point: &Point = point;
        let mut right_number: String = "".to_string();

        while x_right_coord < schematic[right_point.y].len() && right_point.point_type != PointType::Symbol(".".to_string()) {
            x_right_coord = right_point.x + 1;

            right_point = &schematic[right_point.y][x_right_coord];   
            if let PointType::Digit(digit) = right_point.point_type {
                right_number = format!("{}{}", right_number, digit);
            }
        }

        if right_number != "".to_string() {
            let right_number: u32 = right_number.parse::<u32>().expect("Expected valid u32");
            numbers.push(right_number);
        }
    }

    numbers 
}

fn walk_line(point: &Point, line: &Vec<Point>, direction: Direction) -> Result<u32, ParseIntError> {
    let mut current_coord = point.x;
    let mut point = point;
    let mut number = "".to_string();

    while current_coord < line.len() as u32 && point.point_type != PointType::Symbol(".".to_string()) {
        current_coord = point.x + direction;

        point = &line[current_coord as usize];   
        if let PointType::Digit(digit) = point.point_type {
            number = format!("{}{}", number, digit);
        }
    }
    
    number.parse()
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
    fn should_detect_number_x_adjacent_left() {
        let line = "617*......";
        let schematic = vec![parse_schematic_line(0, line)];
        
        let expected = vec![617];
        let result = valid_numbers(schematic); 
        assert_eq!(expected, result);
    }

    #[test]
    fn should_detect_number_x_adjacent_right() {
        let line = "...*12....";
        let schematic = vec![parse_schematic_line(0, line)];
        
        let expected = vec![12];
        let result = valid_numbers(schematic); 
        assert_eq!(expected, result);
    }
   
    #[test]
    fn should_detect_number_x_adjacent_multiple() {
        let line = "...*12.6*.";
        let schematic = vec![parse_schematic_line(0, line)];
        
        let expected = vec![12, 6];
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
