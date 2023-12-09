#[derive(Debug, PartialEq)]
pub enum PointType {
    Symbol(String),
    Digit(u32),
}

#[derive(Debug, PartialEq)]
pub struct Point {
    pub x: u32,
    pub y: u32,
    pub point_type: PointType,
}
