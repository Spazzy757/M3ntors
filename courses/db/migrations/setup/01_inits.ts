import { Knex } from 'knex';

const triggerSetTimestamp = `
CREATE OR REPLACE FUNCTION update_updated_at_column() RETURNS trigger
  LANGUAGE plpgsql
  AS $$
  BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
  END;
  $$;
`

export async function seed(knex: Knex): Promise<void[]> {
  const dbName = 'courses'
  const result = await knex.raw(
    `SELECT 1 AS result FROM pg_database WHERE datname='${dbName}'`
  );
  if (result.rowCount) {
    return;
  }
  console.info(`${dbName} database does not exist, creating it`);
  await knex.raw(`CREATE DATABASE ${dbName};`).then();
  const result1 = await knex.raw(
    `SELECT 1 AS result FROM pg_database WHERE datname='${dbName}'`
  );
  const config: Knex.Config = knex.client.config;
  const newConfig: Knex.Config = {
      ...config,
      connection:
        typeof config.connection != 'string'
          ? { ...config.connection, database: dbName }
          : config.connection,
  };
  const prefixedKnex = require('knex')(newConfig);
  await prefixedKnex.raw(triggerSetTimestamp)
}
