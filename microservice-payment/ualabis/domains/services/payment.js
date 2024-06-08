const UalaApiCheckout = require('ualabis-nodejs');
require('dotenv').config();

const CreatePayment = async (req,res) => {
    try{
        const {amount , description} = req.body;
       
        await UalaApiCheckout.setUp({
            userName: process.env.USER_NAME_UALA,
            clientId: process.env.CLIENT_ID_UALA,
            clientSecret: process.env.CLIENT_SECRET_ID,
            isDev: true
        });

        const order = await UalaApiCheckout.createOrder({
            amount: Number(amount),
            description: description,
            callbackSuccess: 'https://www.google.com/search?q=failed',
            callbackFail: 'https://www.google.com/search?q=success',
        });

        //const generatedOrder = await UalaApiCheckout.getOrder(order.uuid);

        //const orders = await UalaApiCheckout.getOrders({limit:'2', fromDate:'2022-08-04', toDate:'2022-08-09'});

        res.status(200).send({
            status:200,
            message: "Payment created",
            data: order.links.checkoutLink
        });
    }catch(err){
        res.status(500).send({
            status:500,
            message: "Internal server error",
            error: err
        });
    }
}

const GetPayment = async (req,res) => {
    try{}catch(err){}
}

module.exports = {
    CreatePayment,
    GetPayment
}