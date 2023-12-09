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

#[cfg(test)]
mod tests {
    use crate::{point::{Point, PointType}, schematic::new_schematic};

    use super::parse_schematic_line;
    
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
}
