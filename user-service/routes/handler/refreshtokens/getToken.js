const {RefreshToken} = require('../../../models')

module.exports = async (req, res) => {
    const refreshToken = req.query.refresh_token

    if (!refreshToken){
        return res.status(400).json({
            status : "error", 
            message : "invalid token"
        })
    }   

    const token = await RefreshToken.findOne({where : {
        token : refreshToken
    }})

    if (!token){
        return res.status(400).json({
            status : "error", 
            message : "invalid token"
        })
    }

    return res.status(200).json({
        status : "success", 
        data : token
    })
}