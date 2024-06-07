const express = require("express");
const  routers = require("../routes");

const server = express()

routers(server)

module.exports = server