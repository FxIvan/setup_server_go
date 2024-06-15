const {Router} = require('express');
const { CreatePayment , VerifyUala } = require('../domains/services/payment');


const router = Router();

router.route("/create/payment").post(CreatePayment);
router.route("/verify/uala/:id").get(VerifyUala);

module.exports = router;