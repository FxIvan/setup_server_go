const {Router} = require('express');
const { CreatePayment } = require('../domains/services/payment');


const router = Router();

router.route("/create/payment").post(CreatePayment);

module.exports = router;