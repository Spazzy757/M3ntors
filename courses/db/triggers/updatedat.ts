import { Knex } from "knex";

export const addUpdatedAt = (knex: Knex, tableName: string) =>
  knex.raw(`
  DROP TRIGGER IF EXISTS set_timestamp ON "${tableName}";
  CREATE TRIGGER set_timestamp
  BEFORE INSERT OR UPDATE
  ON "${tableName}"
  FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
`);
