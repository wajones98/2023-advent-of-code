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

pub fn valid_numbers(line: Vec<&Point>) -> Vec<u32> {
   vec![0] 
}

#[cfg(test)]
mod tests {
    use crate::{point::{Point, PointType}, schematic::new_schematic};

    use super::{parse_schematic_line, points_with_symbol, valid_numbers};
    
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
        let line = "617*......";
        let line = parse_schematic_line(0, line); 

        let expected = vec![617]; 
        let result = valid_numbers(line);

        assert_eq!(expected, result);
    }
}
