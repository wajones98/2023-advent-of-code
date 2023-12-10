use crate::schematic::{Schematic, generate_x_lines, points_with_symbol, generate_y_lines};

fn gather_lines(schematic: Schematic) {
    let points = points_with_symbol(&schematic); 
    let x_lines = generate_x_lines(&schematic, &points);
    let y_lines = generate_y_lines(&schematic, &points);
}
