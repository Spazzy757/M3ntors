import { Knex } from "knex";
import {addUpdatedAt} from "../triggers/updatedat"

const tableName = "courses"

export async function up(knex: Knex): Promise<void> {
  await knex.schema
    .createTable(tableName, (table) => {
      table.increments('id').primary();
      table.string('name', 255).notNullable();
      table.string('link', 255).notNullable();
      table.boolean('reviewed').defaultTo(false);
      table.timestamp('created_at')
        .defaultTo(knex.fn.now());
    })
    .then(() => Promise.all([addUpdatedAt(knex, tableName)]))
}

export async function down(knex: Knex): Promise<void> {
  await knex.schema.dropTable(tableName);
}

