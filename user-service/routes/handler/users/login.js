const bcrypt = require('bcrypt')
const {User} = require('../../../models')

const Validator = require("fastest-validator")
const v = new Validator();

module.exports = async (req, res) => {

    const schema = {
        email : "email|empty:false",
        password : "string|empty:false"
    }

    const check = v.validate(req.body, schema)

    if (check.length == 0){
        return res.status(400).json({
            status : "error", 
            message : check
        })
    }

    const user = await User.findOne({where : {email : req.body.email}})

    if (!user){
        return res.status(404).json({
            status : "error", 
            message : "User not found"
        })
    }

    const isPasswordValid = await bcrypt.compare(req.body.password, user.password)

    if (!isPasswordValid){
        return res.status(404).json({
            status : "error", 
            message : "User not found"
        })
    }

    res.json({
        status : "success", 
        data : {
            id : user.id, 
            name : user.name, 
            email : user.email,
            role : user.role, 
            profession : user.profession,
        }
    })
    
}