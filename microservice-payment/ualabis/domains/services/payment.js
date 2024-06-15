const UalaApiCheckout = require('ualabis-nodejs');
require('dotenv').config();

const setUpUala = async () => {
    await UalaApiCheckout.setUp({
        userName: process.env.USER_NAME_UALA,
        clientId: process.env.CLIENT_ID_UALA,
        clientSecret: process.env.CLIENT_SECRET_ID,
        isDev: true
    });
};

setUpUala().catch(err => {
    console.error("Error setting up Uala API:", err);
    process.exit(1);
});

const CreatePayment = async (req,res) => {
    try{
        const {amount , description , succesResponse, failedResponse} = req.body;

        const order = await UalaApiCheckout.createOrder({
            amount: amount,
            description: description,
            callbackSuccess: `${succesResponse}`,
            callbackFail: `${failedResponse}`,
        });
        //const generatedOrder = await UalaApiCheckout.getOrder(order.uuid);

        //const orders = await UalaApiCheckout.getOrders({limit:'2', fromDate:'2022-08-04', toDate:'2022-08-09'});

        res.status(200).send({
            status:200,
            message: "Payment created",
            data: order
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

const VerifyUala = async (req,res) => {
    try{
        const {id} = req.params;
        const order = await UalaApiCheckout.getOrder(id);
        console.log("--------------- VerifyUala ---------------\n")
        console.log(order);
        console.log("--------------- VerifyUala ---------------")
        res.status(200).send({
             order
        });
    }catch(err){
        res.status(500).send({
            status:500,
            message: "Internal server error",
            error: err
        });
    }
}   

module.exports = {
    CreatePayment,
    GetPayment,
    VerifyUala
}