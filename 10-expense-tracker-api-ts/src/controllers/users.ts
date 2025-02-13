import { Hono } from "hono";
import { db } from "../db/db_config";
import { usersTable } from "../db/schema";
import { eq } from "drizzle-orm";

type User = typeof usersTable.$inferInsert;

const users = new Hono();

users.get("/", async (c) => {
  const users = await db.query.usersTable.findMany();
  return c.json({ data: users }, 200);
});

users.get("/:id", async (c) => {
  const id = c.req.param("id");

  if (!id) return c.json({ error: "id missing" });

  const user = await db.query.usersTable.findFirst({
    where: eq(usersTable.id, Number(id)),
  });

  if (!user) return c.json({ error: "no user found" }, 404);

  return c.json({ user }, 200);
});

users.post("/", async (c) => {
  const { email, password, username } = await c.req.json<User>();

  if (!email) return c.json({ error: "email missing" }, 400);
  if (!password) return c.json({ error: "password missing" }, 400);
  if (!username) return c.json({ error: "username missing" }, 400);

  const isRepeated = await db.query.usersTable.findFirst({
    where: eq(usersTable.email, email),
  });

  if (isRepeated) return c.json({ error: "user already registered" }, 403);

  const hashedPassword = await Bun.password.hash(password);

  const user = await db
    .insert(usersTable)
    .values({ email, password: hashedPassword, username })
    .returning({
      email: usersTable.email,
      username: usersTable.username,
    });

  return c.json({ status: "created", user: user[0] }, 201);
});

users.put("/:id", async (c) => {
  const id = c.req.param("id");
  const { email, password, username } = await c.req.json<User>();

  if (!id) return c.json({ error: "id missing" }, 404);
  if (!email) return c.json({ error: "email missing" }, 400);
  if (!password) return c.json({ error: "password missing" }, 400);
  if (!username) return c.json({ error: "username missing" }, 400);

  const hashedPassword = await Bun.password.hash(password);

  const updatedUser = await db
    .update(usersTable)
    .set({ email, password: hashedPassword, username })
    .where(eq(usersTable.id, Number(id)))
    .returning({
      email: usersTable.email,
      username: usersTable.username,
    });

  return c.json({ data: { updatedUser } }, 200);
});

users.delete("/:id", async (c) => {
  const id = c.req.param("id");

  if (!id) return c.json({ error: "id missing" }, 404);

  await db.delete(usersTable).where(eq(usersTable.id, Number(id)));

  return c.json({ message: "deleted" }, 202);
});

export default users;
