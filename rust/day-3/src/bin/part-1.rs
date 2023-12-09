use std::ops::Add;

#[derive(Debug, Copy, Clone)]
enum Direction {
    LeftUp = -1,
    RightDown = 1,
}

impl Add<&Direction> for u32 {
    type Output = u32;

    fn add(self, rhs: &Direction) -> Self::Output {
        match *rhs {
            Direction::LeftUp => {
                let rhs_value = *rhs as i32; // Get the associated integer value
                if self >= (-rhs_value) as u32 {
                    self - (-rhs_value) as u32
                } else {
                    9999 // Handle underflow as needed
                }
            }
            Direction::RightDown => {
                let rhs_value = *rhs as u32; // Get the associated integer value
                self + rhs_value
            }
        }
    }
}

fn main() {
    println!("Hello, world!");
}

#[derive(Debug, PartialEq, Eq, Clone)]
enum PointType {
    Symbol(String),
    Digit(u32),
}

#[derive(Debug, PartialEq, Eq, Clone)]
struct Point {
    x: u32,
    y: u32,
    point_type: PointType,
}

type Schematic = Vec<Vec<Point>>;

fn new_schematic(lines: Vec<&str>) -> Schematic {
    lines.iter().enumerate().map(|(i, line)| {
        parse_schematic_line(i, line)
    }).collect()
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
    let mut left_found_points: Vec<Point> = vec![];
    let mut right_found_points: Vec<Point> = vec![];

    for point in points_with_symbols {
        let left_point: &Point = point;
        let left_number = walk_line(left_point, &schematic[left_point.y as usize], &Direction::LeftUp, &mut left_found_points, &mut right_found_points);
        if let Some(value) = left_number {
            println!("1");
            numbers.push(value);
        }

        let right_point: &Point = point;
        let right_number = walk_line(right_point, &schematic[right_point.y as usize], &Direction::RightDown, &mut left_found_points, &mut right_found_points);
        if let Some(value) = right_number {
            println!("2");
            numbers.push(value);
        }
         
        // Handle y
        // let above_y_coord = (point.y + &Direction::LeftUp) as usize; 
        // if above_y_coord != 9999 {
        //     let point_above = &schematic[above_y_coord][point.x as usize];
        // }

        let below_y_coord = (point.y + &Direction::RightDown) as usize; 
        if below_y_coord < schematic.len() {
            let point_below = &schematic[below_y_coord][point.x as usize];

            let left_number = walk_line(point_below, &schematic[point_below.y as usize], &Direction::LeftUp, &mut left_found_points, &mut right_found_points);
            if let Some(value) = left_number {
                numbers.push(value);
            }

            let right_number = walk_line(point_below, &schematic[point_below.y as usize], &Direction::RightDown, &mut left_found_points, &mut right_found_points);
            if let Some(value) = right_number {
                numbers.push(value);
            }
        }
    }

    numbers 
}

fn walk_line(point: &Point, line: &Vec<Point>, direction: &Direction, left_found_points: &mut Vec<Point>, right_found_points: &mut Vec<Point>) -> Option<u32> {
    let mut point = point;
    let mut number = "".to_string();
    let mut current_coord: usize = point.x.try_into().expect("Expected u32 to parse into usize");
    
    while point_coord_is_valid(point, line) {
        if current_coord == 9999 {
            break;
        }

        point = &line[current_coord];  

        if let PointType::Digit(digit) = point.point_type {
            match direction {
                Direction::LeftUp => { 
                    if !right_found_points.contains(point) {
                        number = format!("{}{}", digit, number);
                        right_found_points.push(point.clone());
                    }
                },
                Direction::RightDown => { 
                    if !left_found_points.contains(point) {
                        number = format!("{}{}", number, digit);
                        left_found_points.push(point.clone());
                    }
                },
            };        
        }
        
        current_coord = (point.x + direction).try_into().expect("Expected u32 to parse into usize");
    }

    match number.parse::<u32>() {
        Ok(value) => Some(value),
        Err(_) => None,
    }
}

fn point_coord_is_valid(point: &Point, line: &Vec<Point>) -> bool {
   point.x < line.len() as u32 && point.point_type != PointType::Symbol(".".to_string()) 
}

fn parse_schematic_line(y: usize, line: &str) -> Vec<Point> {
    line.chars().enumerate().map(|(x, c)| {
        let point_type: PointType = match c.to_digit(10) {
            Some(digit) => PointType::Digit(digit),
            None => PointType::Symbol(c.to_string()),
        };
        
        Point {
            x: (x as u32),
            y: (y as u32),
            point_type,
        }
    }).collect() 
}

#[cfg(test)]
mod tests {
    use crate::{Schematic, new_schematic, parse_schematic_line, Point, PointType, valid_numbers, Direction};

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
    fn should_add_direction_left() {
        let point = Point {
            x: 5,
            y: 0,
            point_type: PointType::Digit(1)
        };
        let direction = &Direction::LeftUp; 
        let result = point.x + direction; 
        
        let expected = 4;
        assert_eq!(expected, result);
    }

    #[test]
    fn should_add_direction_right() {
        let point = Point {
            x: 5,
            y: 0,
            point_type: PointType::Digit(1)
        };
        let direction = &Direction::RightDown; 
        let result = point.x + direction; 
        
        let expected = 6;
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

    #[test]
    fn should_detect_number_y_adjacent() {
        let lines = vec![
            "...$.*....", 
            ".664.598.."
        ];
        let schematic = new_schematic(lines); 

        let expected = vec![664, 598];
        let result = valid_numbers(schematic);
        assert_eq!(expected, result);
    }

    // #[test]
    // fn should_detect_number_diagonally() {}
    //
    // #[test]
    // fn should_calculate_total() {}
}
