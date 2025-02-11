import {
  int,
  sqliteTable,
  text,
  real,
  AnySQLiteColumn,
} from "drizzle-orm/sqlite-core";

export const usersTable = sqliteTable("users", {
  id: int().primaryKey({ autoIncrement: true }),
  username: text().notNull().unique(),
  email: text().notNull().unique(),
  password: text().notNull(),
  createdAt: int("created_at", { mode: "timestamp" }).$defaultFn(
    () => new Date()
  ),
});

export const expensesTable = sqliteTable("expenses", {
  id: int().primaryKey({ autoIncrement: true }),
  name: text().notNull(),
  amount: real().notNull().default(0.0),
  categoryId: int().references((): AnySQLiteColumn => categoryTable.id),
  ownerId: int().references((): AnySQLiteColumn => usersTable.id),
});

export const categoryTable = sqliteTable("categories", {
  id: int().primaryKey({ autoIncrement: true }),
  name: text().notNull(),
  expenseId: int().references((): AnySQLiteColumn => expensesTable.id),
  ownerId: int().references((): AnySQLiteColumn => usersTable.id),
});

export const table = {
  usersTable,
  expensesTable,
  categoryTable,
} as const;

export type Table = typeof table;
