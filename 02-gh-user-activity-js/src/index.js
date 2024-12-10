import { Command } from "commander";
import { fetchUserActivity } from "./userActivity.js";

const program = new Command();

program
  .version("1.0.0")
  .description("A CLI to view the github activity")
  .argument('<string>', 'github username')
  .action((str) => {
    fetchUserActivity(str);
  })

if (!process.argv.slice(2).length) {
  program.help();
}

program.parse();
