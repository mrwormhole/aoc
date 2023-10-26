use std::fs;

#[derive(Debug, Eq, Ord, PartialEq, PartialOrd)]
struct Elf {
    id: u64,
    total_calories: u64,
}

fn main() {
    let path = "./input.txt";
    let content = fs::read_to_string(path).expect("should read the file");
    let mut elves: Vec<Elf> = vec![];
    let mut unique_id: u64 = 1;

    let lines: Vec<&str> = content.split("\n").collect();
    for (index, line) in lines.iter().enumerate() {
        if *line == "" {
            let mut sum: u64= 0;
            let mut temp:usize = index;
            loop {
                temp -= 1;
                sum += lines[temp].parse::<u64>().unwrap_or(0);

                if lines[temp] == "" || temp == 0 {
                    let elf: Elf = Elf { id: unique_id, total_calories: sum };
                    elves.push(elf);
                    unique_id += 1;
                    break;
                }
            }
        }
    }

    elves.sort_by(|a, b| b.total_calories.cmp(&a.total_calories));
    println!("first:{:?} second:{:?} third:{:?}", elves[0], elves[1], elves[2]);
    println!("stash:{:?}", elves[0].total_calories + elves[1].total_calories + elves[2].total_calories);
}
