const bcrypt = require('bcrypt')
const {User} = require("../../../models")

const Validator = require("fastest-validator")
const v = new Validator();
module.exports = async (req, res) => {

    //  buat skema 
    const schema = {
        name : "string|empty:false",
        profession : "string|optional", 
        password: "string|min:6|max:16|empty:false",
        email : "email|empty:false" 
    }

    const check = v.validate(req.body, schema)

    if (check.length > 0){
        return res.status(400).json({
            status : "error", 
            message : check
        })
    }

    // cek apakah email sudah tersedia
    const user  = await User.findOne({where : {email :req.body.email}})

    if (user){
        return res.status(409).json({
            status : "error", 
            message : "Email already exist" 
        });
    }

    const password = await bcrypt.hash(req.body.password, 10)
    const data = {
        name : req.body.name, 
        password,
        email : req.body.email, 
        profession : req.body.profession, 
        role : "student" 
    }

    const createdUser = await User.create(data)
    return res.status(200).json({
       status : "success" , 
       data : {
        id : createdUser.id,
        name : createdUser.name, 
        profession : createdUser.profession,
        role :  createdUser.role,
       }
    })

}