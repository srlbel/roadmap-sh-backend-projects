import { Hono } from "hono";

const login = new Hono();

//  TODO: Implement JWT Token
login.post("/login", (c) => c.text("todo"));

// TODO: Create user
login.post("/register", (c) => c.text("todo"));

export default login;
