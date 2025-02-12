import { Hono } from "hono";

const users = new Hono();

type User = {
  username: string;
  email: string;
  password: string;
};

users.get("/", (c) => {
  return c.json({ method: "GET" });
});

users.get("/:id", (c) => {
  const id = c.req.param("id");

  if (!id) return c.json({ error: "id missing" });

  return c.json({ method: "GET", data: { id } });
});

users.post("/", async (c) => {
  const body = await c.req.json<User>();

  if (!body.email) return c.json({ error: "email missing" }, 400);
  if (!body.password) return c.json({ error: "password missing" }, 400);
  if (!body.username) return c.json({ error: "username missing" }, 400);

  return c.json({ method: "POST", data: body });
});

users.put("/:id", async (c) => {
  const id = c.req.param("id");
  const body = await c.req.json<User>();

  if (!id) return c.json({ error: "id missing" });
  if (!body.email) return c.json({ error: "email missing" }, 400);
  if (!body.password) return c.json({ error: "password missing" }, 400);
  if (!body.username) return c.json({ error: "username missing" }, 400);

  return c.json({ method: "PUT", data: { id, body } });
});

users.delete("/:id", (c) => {
  const id = c.req.param("id");

  if (!id) return c.json({ error: "id missing" });

  return c.json({ method: "DELETE", data: { id } });
});

export default users;
