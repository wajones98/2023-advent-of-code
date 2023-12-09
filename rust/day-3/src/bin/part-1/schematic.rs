pub type Schematic = Vec<Vec<Point>>;

fn new_schematic(lines: Vec<&str>) -> Schematic {
    lines.iter().enumerate().map(|(i, line)| {
        parse_schematic_line(i, line)
    }).collect()
}

