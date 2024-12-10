use chrono::NaiveDate;
use serde::{Deserialize, Serialize};
use std::{collections::HashMap, fs, path::PathBuf};

const PATH: &str = "./expenses.json";

#[derive(Serialize, Deserialize, Clone, Debug)]
pub struct Expense {
    pub id: usize,
    pub description: String,
    pub amount: f32,
    pub date: NaiveDate,
    updated: NaiveDate,
}

pub trait ExpenseImpl {
    fn new(id: usize, description: String, amount: f32) -> Self;
    fn update(&mut self, description: Option<String>, amount: Option<f32>);
}

impl ExpenseImpl for Expense {
    fn new(id: usize, description: String, amount: f32) -> Self {
        Self {
            id,
            description,
            amount,
            date: chrono::Local::now().date_naive(),
            updated: chrono::Local::now().date_naive(),
        }
    }

    fn update(&mut self, description: Option<String>, amount: Option<f32>) {
        if let Some(description) = description {
            self.description = description;
        };

        if let Some(amount) = amount {
            self.amount = amount;
        };

        self.updated = chrono::Local::now().date_naive();
    }
}

#[derive(Serialize, Deserialize)]
pub struct ExpenseList {
    pub list: HashMap<usize, Expense>,
    pub next_id: usize,
}

pub trait ExpenseListImpl {
    fn new() -> Self;
    fn load() -> Self;
    fn remove_item(&mut self, id: usize);
    fn save(&self);
}

impl ExpenseListImpl for ExpenseList {
    fn new() -> Self {
        Self {
            list: HashMap::new(),
            next_id: 1,
        }
    }

    fn load() -> Self {
        let path = PathBuf::from(PATH);

        if !path.exists() {
            return ExpenseList::new();
        }

        let data: &str = &fs::read_to_string(path).unwrap();
        serde_json::from_str(data).unwrap()
    }

    fn remove_item(&mut self, id: usize) {
        self.list.remove(&id).expect("No ID Matched");
    }

    fn save(&self) {
        let data = serde_json::to_string_pretty(&self).unwrap();
        fs::write(PATH, data).unwrap();
    }
}
