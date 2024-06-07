
const CreatePayment = async (req,res) => {
    try{
        const {amount , description} = req.body;
        console.log("CreatePayment", amount);
        return res.status(200).send("No data!");
    }catch(err){}
}

const GetPayment = async (req,res) => {
    try{}catch(err){}
}

module.exports = {
    CreatePayment,
    GetPayment
}