const {User, RefreshToken} = require('../../../models')


module.exports = async (req, res) => {
    const userId = req.body.user_id
    const user = await User.findByPk(userId)


    if (!user){
        return res.status(404).json({
            sxtatus : "error", 
            message : "User not found"
        })
    }

    await RefreshToken.destroy({
        where : {
            user_id : user.id
        }
    });


    return res.json({
        status : 'success', 
        message : 'refresh token deleted'
    });
}