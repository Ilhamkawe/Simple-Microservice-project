const {User} = require("../../../models")

module.exports = async (req, res) => {
    const userIds = req.query.user_ids || []
    const sqlAtt = {attributes : ['id','email', 'name', 'profession', 'avatar', 'role']}

    if (userIds.length) {
        sqlAtt.where = {
            id : userIds
        }
    }

    const users = await User.findAll(sqlAtt)

    return res.status(200).json({
        status : "success", 
        data : users
    })

}