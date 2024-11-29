use rand::{thread_rng, Rng};
use std::io;

fn main() {
    let mut rng = thread_rng();
    let mut difficulty = String::new();
    let random_number: i32 = rng.gen_range(1..=100);

    println!(
        "Welcome to the Number Guessing Game! \nI'm thinking of a number between 1 and 100. \n"
    );

    println!("Please select the difficulty level: \n1. Easy (10 chances) \n2. Medium (5 chances) \n3. Hard (3 chances)\n");

    print!("Enter your choice: ");
    io::Write::flush(&mut io::stdout()).expect("Failed to flush stdout");

    io::stdin()
        .read_line(&mut difficulty)
        .expect("Failed to read line");

    let difficulty = difficulty.trim();

    match difficulty {
        "1" => game(10, "Easy", random_number),
        "2" => game(5, "Medium", random_number),
        "3" => game(3, "Hard", random_number),
        _ => eprintln!("Not an option of the list"),
    }
}

fn game(attemps: i8, difficulty: &str, random: i32) {
    println!(
        "Great, you have selected the {} diffculty level. \nlet's start the game\n",
        difficulty
    );

    for attemp in 1..=attemps {
        let mut guess = String::new();

        print!("Enter your guess: ");
        io::Write::flush(&mut io::stdout()).expect("Failed to flush stdout");
        io::stdin().read_line(&mut guess).expect("Failed to read");

        let guess = guess.trim();

        match guess.parse::<i32>() {
            Ok(guess) => {
                let is_correct: bool = check_number(guess, random, attemp);
                if is_correct {
                    break;
                }
            }
            Err(_) => eprintln!("That's not a valid number \n"),
        }
    }
}

fn check_number(guess: i32, random: i32, attemps: i8) -> bool {
    if guess > random {
        println!("Incorrect! the number is less than {} \n", guess);
        return false;
    } else if guess < random {
        println!("Incorrect! the number is great than {} \n", guess);
        return false;
    } else {
        println!(
            "Congratulations! You guessed the correct number in {} attempts. \n",
            attemps
        );
        return true;
    }
}
