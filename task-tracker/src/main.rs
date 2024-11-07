use clap::Parser;
use cli::{Cli, Command};
use task::{Task, TaskImpl, TaskList, TaskListImpl};

mod cli;
mod task;

fn main() {
    let mut tasks = TaskList::load();
    let args = Cli::parse();

    match args.command {
        Command::Add { description } => {
            let next_id: usize = tasks.next_id;
            let task: Task = Task::new(next_id, description);
            tasks.list.insert(next_id, task);
            tasks.next_id += 1;
            tasks.save();
        }
        Command::ChangeStatus { status, id } => {
            let task = tasks.list.get_mut(&id).unwrap();
            match status.as_str() {
                "in-progress" => {
                    task.update_status(task::TaskStatus::InProgress);
                    tasks.save();
                }
                "done" => {
                    task.update_status(task::TaskStatus::Done);
                    tasks.save()
                }
                _ => panic!("Invalid status"),
            }
        }
        Command::Delete { id } => {
            tasks.remove_item(id);
            tasks.save();
        }
        Command::List { status } => {
            let tasks = tasks.list.values();

            if let Some(status) = status {
                let status = match status.as_str() {
                    "todo" => task::TaskStatus::Todo,
                    "in-progress" => task::TaskStatus::InProgress,
                    "done" => task::TaskStatus::Done,
                    _ => panic!("Invalid status provided"),
                };

                let filtered_tasks = tasks.filter(|task| task.status == status);
                for task in filtered_tasks {
                    println!("{:?}", task);
                }
            } else {
                for task in tasks {
                    println!("{:?}", task);
                }
            }
        }
        Command::Update { id, description } => {
            let task = tasks.list.get_mut(&id).unwrap();
            task.update(description);
            tasks.save();
        }
    }
}
