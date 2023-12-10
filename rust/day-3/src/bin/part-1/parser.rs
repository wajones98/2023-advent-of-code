use crate::point;
use crate::schematic::*;
use crate::point::*;

fn gather_lines(schematic: Schematic) {
    let points = points_with_symbol(&schematic); 
    let x_lines = generate_x_lines(&schematic, &points);
    let y_lines = generate_y_lines(&schematic, &points);
    let diagonal_left = generate_diagonal_lines_left(&schematic, &points);
    let diagonal_right = generate_diagonal_lines_right(&schematic, &points);


}

fn parse_line(line: &Vec<Point>, schematic: &Schematic) -> Vec<u32> {
    let mut numbers: Vec<u32> = vec![];

    for (i, _) in line.iter().enumerate() {
        if i > 0 {
            let prev_point = &line[i - 1];            

            if let PointType::Digit(_) = prev_point.point_type {
                let x_points = &schematic[prev_point.y as usize];   
                numbers.append(&mut valid_numbers(x_points.to_vec()));  
            }
        } 

        if i < line.len() - 1 {
            let next_point = &line[i + 1];            
            let x_points = &schematic[next_point.y as usize];   
            numbers.append(&mut valid_numbers(x_points.to_vec()));  
        }
    }

    numbers
}

#[cfg(test)]
mod tests {
    use super::*;
    
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
    fn it_extracts_numbers_from_x_lines() {
        let line = vec!["617*......"];
        let schematic = new_schematic(line);
        let points = points_with_symbol(&schematic); 

        let x_lines = generate_x_lines(&schematic, &points);
        println!("{:?}", x_lines);
        let result = parse_line(&x_lines[0], &schematic);
        println!("{:?}", x_lines);

        let expected = vec![617];

        // assert_eq!(expected, result);
    }
}
