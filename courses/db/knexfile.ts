import type { Knex } from 'knex';

const dbname = 'courses'

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
      database: dbname,
      user: 'postgres',
      password: 'postgres',
    },
    pool: {
      min: 2,
      max: 10,
    },
  },
  containersetup: {
    client: 'postgresql',
    connection: {
      host: '/var/run/postgresql',
      user: 'postgres',
      database: 'postgres',
    },
    pool: {
      min: 2,
      max: 10,
    },
    seeds: {
      directory: 'migrations/setup',
    },
  },
  container: {
    client: 'postgresql',
    connection: {
      host: '/var/run/postgresql',
      user: 'postgres',
      database: dbname,
    },
    pool: {
      min: 2,
      max: 10,
    },
    migrations: {
      directory: 'migrations',
    },
  },
};

module.exports = config;
