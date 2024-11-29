use chrono::NaiveDateTime;
use serde::{Deserialize, Serialize};
use std::{collections::HashMap, fs, path::PathBuf};

const PATH: &str = "./task.json";

#[derive(Serialize, Deserialize, Debug, PartialEq)]
pub enum TaskStatus {
    Todo,
    InProgress,
    Done,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Task {
    id: usize,
    description: String,
    pub status: TaskStatus,
    created_at: NaiveDateTime,
    updated_at: NaiveDateTime,
}

pub trait TaskImpl {
    fn new(id: usize, description: String) -> Self;
    fn update(&mut self, description: String);
    fn update_status(&mut self, status: TaskStatus);
}

impl TaskImpl for Task {
    fn new(id: usize, description: String) -> Self {
        Self {
            id,
            description,
            status: TaskStatus::Todo,
            created_at: chrono::Local::now().naive_local(),
            updated_at: chrono::Local::now().naive_local(),
        }
    }
    fn update(&mut self, description: String) {
        self.description = description;
        self.updated_at = chrono::Local::now().naive_local();
    }
    fn update_status(&mut self, status: TaskStatus) {
        self.status = status;
        self.updated_at = chrono::Local::now().naive_local();
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct TaskList {
    pub list: HashMap<usize, Task>,
    pub next_id: usize,
}

pub trait TaskListImpl {
    fn new() -> Self;
    fn load() -> Self;
    fn remove_item(&mut self, id: usize);
    fn save(&self);
}

impl TaskListImpl for TaskList {
    fn new() -> Self {
        Self {
            list: HashMap::new(),
            next_id: 1,
        }
    }

    fn load() -> Self {
        let path = PathBuf::from(PATH);
        if !path.exists() {
            return TaskList::new();
        }

        let data: &str = &fs::read_to_string(path).unwrap();
        serde_json::from_str(data).unwrap()
    }

    fn save(&self) {
        let data = serde_json::to_string_pretty(&self).unwrap();
        fs::write(PATH, data).unwrap();
    }

    fn remove_item(&mut self, id: usize) {
        self.list.remove(&id).expect("No ID matched");
    }
}
