const express = require('express');

const routers = (app) => {
    //ruta principal
    app.get('/', async (_req, res) => {
        res.send("Run!");
    });
    app.use(express.json());
    //rutas de api
    app.use('/api/', require('../controllers'));
};

module.exports = routers;

