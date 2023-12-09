pub enum PointType {
    Symbol(String),
    Digit(u32),
}

pub struct Point {
    x: u32,
    y: u32,
    point_type: PointType,
}
