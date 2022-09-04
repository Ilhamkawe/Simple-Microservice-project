const bcrypt = require('bcrypt')
const {User} = require('../../../models')

const Validator = require("fastest-validator")
const v = new Validator()
module.exports = async (req, res) => {

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

    const id = req.params.id
    const user = await User.findByPk(id);

    if (!user){
        return res.status(400).json({
            status : "error", 
            message : "User not found"
        })
    }

    const email = req.body.email
    if (email){
        const checkEmail = await User.findOne({where : {email}})
        if (checkEmail && email !== user.email){
            return res.json(409).json({
                status : "error", 
                message : "email already taken"
            })
        }
    }

    const password = await bcrypt.hash(req.body.password, 10)
    const {name, profession, avatar} = req.body
    
    await user.update({
        name,
        email, 
        password, 
        profession, 
        avatar
    })

    return res.status(200).json({
        status : "success", 
        data : {
            id : user.id, 
            name : user.name,
            email : user.email, 
            password : user.password, 
            profession : user.profession, 
            avatar : user.avatar 
        }
    })

}