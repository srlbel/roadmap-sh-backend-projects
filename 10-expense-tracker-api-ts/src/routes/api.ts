import { Hono } from "hono";
import users from "../controllers/users";

const api = new Hono().basePath("/v1");

api.route("/users", users);

export default api;
