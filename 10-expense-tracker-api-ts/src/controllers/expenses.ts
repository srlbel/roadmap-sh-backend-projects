import { Hono } from "hono";

const expenses = new Hono();

// TODO: complete routes
expenses.get("/", (c) => c.text("todo"));
expenses.get("/:id", (c) => c.text("todo"));
expenses.post("/", (c) => c.text("todo"));
expenses.put("/:id", (c) => c.text("todo"));
expenses.delete("/:id", (c) => c.text("todo"));

export default expenses;
