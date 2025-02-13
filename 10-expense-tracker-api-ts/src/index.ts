import { Hono } from "hono";
import { logger } from "hono/logger";
import api from "./routes/api";

const app = new Hono();

app.use(logger());
app.get("/", (c) => c.text("Expense Tracker root"));
app.route("/api", api);

export default app;
// TODO: Fix return http status code
