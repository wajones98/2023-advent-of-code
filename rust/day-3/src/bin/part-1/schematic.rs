use std::collections::HashSet;

use crate::point::{Point, PointType};

pub type Schematic = Vec<Vec<Point>>;

pub fn new_schematic(lines: Vec<&str>) -> Schematic {
    lines.iter().enumerate().map(|(i, line)| {
        parse_schematic_line(i, line)
    }).collect()
}

pub fn parse_schematic_line(y: usize, line: &str) -> Vec<Point> {
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

pub fn points_with_symbol(schematic: &Schematic) -> Vec<&Point> {
   schematic.iter().map(|line| {
        line.iter().filter(|point| -> bool {
            match &point.point_type {
                PointType::Symbol(point_type) => point_type != ".",
                PointType::Digit(_) => false,
            }
        })
    }).flatten().collect()
}

pub fn valid_numbers(points: Vec<Point>) -> Vec<u32> {
    let mut numbers: Vec<u32> = vec![];     

    for (i, point) in points.iter().enumerate() {
        if let PointType::Symbol(symbol) = &point.point_type {
            if symbol != "." {
                let mut left_current_index: isize = i as isize; 

                if left_current_index > 0 {
                    let mut number = String::from(""); 

                    loop {
                        if left_current_index < 0 {
                            break;
                        }

                        let left_point = &points[left_current_index as usize];
                        match &left_point.point_type {
                            PointType::Digit(digit) => {
                                number = format!("{}{}", digit, number);
                            }
                            PointType::Symbol(symbol) => if symbol == "." {
                                break
                            },
                        }
                        left_current_index -= 1;
                    }

                    match number.parse::<u32>() {
                        Ok(value) => {
                            numbers.push(value); 
                        },
                        Err(_) => {},
                    };
                }
                
                let mut right_current_index: isize = i as isize; 

                if right_current_index < points.len() as isize {
                    let mut number = String::from(""); 

                    loop {
                        if right_current_index >= points.len() as isize {
                            break;
                        }

                        let right_point = &points[right_current_index as usize];
                        match &right_point.point_type {
                            PointType::Digit(digit) => {
                                number = format!("{}{}", number, digit);
                            }
                            PointType::Symbol(symbol) => if symbol == "." {
                                break
                            },
                        }
                        right_current_index += 1;
                    }

                    match number.parse::<u32>() {
                        Ok(value) => {
                            numbers.push(value); 
                        },
                        Err(_) => {},
                    };
                }
            }
        }
    } 

    numbers
}

pub fn generate_x_lines(schematic: &Schematic, points: Vec<&Point>) -> Vec<Vec<Point>> {
    let x_lines_with_symbols: Vec<u32> = points.iter()
        .map(|point| point.y)
        .collect::<HashSet<_>>()
        .into_iter()
        .collect();
    
    let mut lines: Vec<Vec<Point>> = vec![];
    for y_coord in x_lines_with_symbols {
        let line = &schematic[y_coord as usize];
        lines.push(line.to_vec());
    }

    lines
}

pub fn generate_y_lines(schematic: &Schematic, points: Vec<&Point>) -> Vec<Vec<Point>> {
    let y_lines_with_symbols: Vec<u32> = points.iter()
        .map(|point| point.x)
        .collect::<HashSet<_>>()
        .into_iter()
        .collect();
    
    let mut lines: Vec<Vec<Point>> = vec![];
    for x_coord in y_lines_with_symbols {
        let mut line: Vec<Point> = vec![];
        for x_line in schematic {
           line.push(x_line[x_coord as usize].clone()); 
        }
        lines.push(line);
    }
    
    lines
}

pub fn generate_diagonal_lines_left(schematic: &Schematic, points: Vec<&Point>) -> Vec<Vec<Point>> {
    let mut lines: Vec<Vec<Point>> = vec![];
    for point in points {
        let mut line: Vec<Point> = vec![];
        let mut x = point.x as i32;
        let mut y = point.y as i32;

        while x > 0 && y < (schematic.len() - 1) as i32 {
            x -= 1;
            y -= 1;
            line.push(schematic[y as usize][x as usize].clone()); 
        }

        x = point.x as i32;
        y = point.y as i32;
        line.push(schematic[y as usize][x as usize].clone()); 

        while x < (schematic[y as usize].len() - 1) as i32  && y < (schematic.len() - 1) as i32 {
            x += 1;
            y += 1;
            line.push(schematic[y as usize][x as usize].clone()); 
        }

        lines.push(line);
    }

    lines.dedup();
    lines
}

pub fn generate_diagonal_lines_right(schematic: &Schematic, points: Vec<&Point>) -> Vec<Vec<Point>> {
    let mut lines: Vec<Vec<Point>> = vec![];
    for point in points {
        let mut line: Vec<Point> = vec![];
        let mut x = point.x as i32;
        let mut y = point.y as i32;

        while x < (schematic[y as usize].len() - 1) as i32 && y > 0 {
            x += 1;
            y -= 1;
            line.push(schematic[y as usize][x as usize].clone()); 
        }

        x = point.x as i32;
        y = point.y as i32;
        line.push(schematic[y as usize][x as usize].clone()); 

        while x > 0 && y < (schematic.len() - 1) as i32 {
            x -= 1;
            y += 1;
            line.push(schematic[y as usize][x as usize].clone()); 
        }

        lines.push(line);
    }

    lines.dedup();
    lines
}


#[cfg(test)]
mod tests {
    use crate::{point::{Point, PointType}, schematic::new_schematic};

    use super::*;
    
    #[test]
    fn it_parses_line() {
        let line = "4.51.";
        let expected = vec![
            Point {
                x: 0,
                y: 0,
                point_type: PointType::Digit(4),
            },
            Point {
                x: 1,
                y: 0,
                point_type: PointType::Symbol(".".to_string()),
            },
            Point {
                x: 2,
                y: 0,
                point_type: PointType::Digit(5),
            },
            Point {
                x: 3,
                y: 0,
                point_type: PointType::Digit(1),
            },
            Point {
                x: 4,
                y: 0,
                point_type: PointType::Symbol(".".to_string()),
            },
        ];       
        let result = parse_schematic_line(0, line);

        assert_eq!(expected, result);
    }

    #[test]
    fn it_parses_schematic() {
        let lines = vec!["4..", ".5."];

        let expected = vec![
            vec![
                Point {
                    x: 0,
                    y: 0,
                    point_type: PointType::Digit(4),
                },
                Point {
                    x: 1,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string()),
                },
                Point {
                    x: 2,
                    y: 0,
                    point_type: PointType::Symbol(".".to_string()),
                },
            ],
            vec![
                Point {
                    x: 0,
                    y: 1,
                    point_type: PointType::Symbol(".".to_string()),
                },
                Point {
                    x: 1,
                    y: 1,
                    point_type: PointType::Digit(5),
                },
                Point {
                    x: 2,
                    y: 1,
                    point_type: PointType::Symbol(".".to_string()),
                },
            ],
        ];

        let result = new_schematic(lines);
        assert_eq!(expected, result);
    }

    #[test]
    fn it_identifies_points_with_symbols() {
        let lines = vec!["*..", ".$."];
        let schematic = new_schematic(lines);
        
        let expected = vec![
            &schematic[0][0],
            &schematic[1][1],
        ];

        let result = points_with_symbol(&schematic);
        assert_eq!(expected, result);
    }

    #[test]
    fn it_extracts_correct_numbers() {
        let line = "617*...$1.";
        let line = parse_schematic_line(0, line); 

        let expected = vec![617, 1]; 
        let result = valid_numbers(line);

        assert_eq!(expected, result);
    }

    #[test]
    fn it_generates_x_lines() {
        let lines = vec!["*$.", "..."];
        let schematic = new_schematic(lines);

        let expected = vec![vec![
            Point {
                x: 0,
                y: 0,
                point_type: PointType::Symbol("*".to_string()),
            },
            Point {
                x: 1,
                y: 0,
                point_type: PointType::Symbol("$".to_string()),
            },
            Point {
                x: 2,
                y: 0,
                point_type: PointType::Symbol(".".to_string()),
            },
        ]];
        
        let symbols = points_with_symbol(&schematic);
        let result = generate_x_lines(&schematic, symbols); 

        assert_eq!(expected, result);
    }

    #[test]
    fn it_generates_y_lines() {
        let lines = vec![".*.", ".*."];
        let schematic = new_schematic(lines);

        let expected = vec![vec![
            Point {
                x: 1,
                y: 0,
                point_type: PointType::Symbol("*".to_string()),
            },
            Point {
                x: 1,
                y: 1,
                point_type: PointType::Symbol("*".to_string()),
            },
        ]]; 

        let symbols = points_with_symbol(&schematic);
        let result = generate_y_lines(&schematic, symbols);
        
        assert_eq!(expected, result);
    }
    
    #[test]
    fn it_generates_diaganol_lines_left() {
        let lines = vec![
            "*..", 
            ".*.", 
            "..1"
        ];
        let schematic = new_schematic(lines);

        let expected = vec![vec![
            Point {
                x: 0,
                y: 0,
                point_type: PointType::Symbol("*".to_string()),
            },
            Point {
                x: 1,
                y: 1,
                point_type: PointType::Symbol("*".to_string()),
            },
            Point {
                x: 2,
                y: 2,
                point_type: PointType::Digit(1),
            },
        ]]; 

        let symbols = points_with_symbol(&schematic);
        let result = generate_diagonal_lines_left(&schematic, symbols);
        
        assert_eq!(expected, result);

    }
    
    #[test]
    fn it_generates_diaganol_lines_right() {
        let lines = vec![
            ".2.", 
            "$..", 
            "..."
        ];
        let schematic = new_schematic(lines);

        let expected = vec![vec![
            Point {
                x: 1,
                y: 0,
                point_type: PointType::Digit(2),
            },
            Point {
                x: 0,
                y: 1,
                point_type: PointType::Symbol("$".to_string()),
            },
        ]]; 

        let symbols = points_with_symbol(&schematic);
        let result = generate_diagonal_lines_right(&schematic, symbols);
        println!("{:?}", result); 
        assert_eq!(expected, result);
    }
}
