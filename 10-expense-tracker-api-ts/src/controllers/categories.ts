import { Hono } from "hono";

const categories = new Hono();

// TODO: complete routes
categories.get("/", (c) => c.text("todo"));
categories.get("/:id", (c) => c.text("todo"));
categories.post("/", (c) => c.text("todo"));
categories.put("/:id", (c) => c.text("todo"));
categories.delete("/:id", (c) => c.text("todo"));

export default categories;
