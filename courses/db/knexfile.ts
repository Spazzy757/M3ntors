import type { Knex } from 'knex';

const config: { [key: string]: Knex.Config } = {
  setup: {
    client: 'postgresql',
    connection: {
      host: 'localhost',
      database: 'postgres',
      user: 'postgres',
      password: 'postgres',
    },
    pool: {
      min: 2,
      max: 10,
    },
    seeds: {
      directory: 'migrations/setup/',
    },
  },
  migrate: {
    client: 'postgresql',
    connection: {
      host: 'localhost',
      database: 'courses',
      user: 'postgres',
      password: 'postgres',
    },
    pool: {
      min: 2,
      max: 10,
    },
  },
};

module.exports = config;
