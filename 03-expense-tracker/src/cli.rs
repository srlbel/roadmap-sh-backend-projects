use clap::{Parser, Subcommand};

#[derive(Parser)]
#[command(version = "1")]
pub struct Cli {
    #[clap(subcommand)]
    pub command: Command,
}

#[derive(Subcommand)]
pub enum Command {
    Add {
        #[clap(short, long)]
        description: String,
        #[clap(short, long)]
        amount: f32,
    },
    List {
        #[clap(short, long)]
        format: Option<String>,
    },
    Summary {
        #[clap(short, long)]
        month: Option<u8>,
    },
    Delete {
        #[clap(short, long)]
        id: usize,
    },
    Update {
        #[clap(short, long)]
        id: usize,
        #[clap(short, long)]
        description: Option<String>,
        #[clap(short, long)]
        amount: Option<f32>,
    },
}
