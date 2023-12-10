#[derive(Debug, PartialEq, Clone)]
pub enum PointType {
    Symbol(String),
    Digit(u32),
}

#[derive(Debug, PartialEq, Clone)]
pub struct Point {
    pub x: u32,
    pub y: u32,
    pub point_type: PointType,
}
