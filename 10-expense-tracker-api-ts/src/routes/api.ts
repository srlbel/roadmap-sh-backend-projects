import { Hono } from "hono";
import users from "../controllers/users";
import login from "../controllers/login";
import expenses from "../controllers/expenses";
import categories from "../controllers/categories";

const api = new Hono().basePath("/v1");

api.route("/users", users);
api.route("/expenses", expenses);
api.route("/categories", categories);
api.route("/", login);

export default api;
