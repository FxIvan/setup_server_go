const {Pool} = require('pg');
require('dotenv').config();
const config = require("config-yml")

const db = new Pool({
    host: process.env.PG_HOST,
    port: config.services.db.port,
    user: process.env.PG_USER,
    password: config.services.db.environment.POSTGRES_PASSWORD,
    database: process.env.PG_DATABASE
  });

module.exports = db;