use clap::{Parser, Subcommand};

#[derive(Parser)]
#[command(version = "1")]
pub struct Cli {
    #[clap(subcommand)]
    pub command: Command,
}

#[derive(Subcommand)]
pub enum Command {
    Add { description: String },
    Update { id: usize, description: String },
    Delete { id: usize },
    ChangeStatus { id: usize, status: String },
    List { status: Option<String> },
}
