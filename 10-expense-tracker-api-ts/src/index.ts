import { Elysia, t } from "elysia";
import db from "./config/db";
import { table, usersTable } from "./database/schema";
import { createInsertSchema } from "drizzle-typebox";
import swagger from "@elysiajs/swagger";

const _createUser = createInsertSchema(table.usersTable, {
  email: t.String({ format: "email" }),
  username: t.String(),
  password: t.String(),
});

const app = new Elysia()
  .use(swagger())
  .get("/", async () => {
    return await db.select().from(usersTable);
  })
  .post(
    "/",
    async ({ body }) => {
      await db.insert(usersTable).values(body);
      return body;
    },
    {
      body: t.Omit(_createUser, ["id", "createdAt"]),
    }
  )
  .listen(3000);

console.log(
  `🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`
);
