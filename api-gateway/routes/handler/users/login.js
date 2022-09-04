const apiAdapter = require("../../apiAdapter")

const {
    URL_SERVICE_USERS,
    JWT_SECRET, 
    JWT_SECRET_REFRESH_TOKEN, 
    JWT_ACCESS_TOKEN_EXPIRED, 
    JWT_REFRESH_TOKEN_EXPIRED
    } = process.env

const jwt = require("jsonwebtoken")

const api = apiAdapter(URL_SERVICE_USERS)

module.exports = async (req, res) => {

    try {
        const user = await api.post("/users/login", req.body)
        const data = user.data.data
        const token = jwt.sign({data}, JWT_SECRET, {expiresIn : JWT_ACCESS_TOKEN_EXPIRED} )
        const refreshToken = jwt.sign({data}, JWT_SECRET_REFRESH_TOKEN, {expiresIn : JWT_REFRESH_TOKEN_EXPIRED})
        await api.post("/refresh_token", {user_id : data.id, refresh_token : refreshToken})

        return res.json({
            message : "Success", 
            data : {
                token : token, 
                refresh_token : refreshToken
            }
        })
        
    }catch (error) {

        if(error.code == "ECONNREFUSED"){
            return res.status(500).json({status : "error", message : "service tidak teresedia"})
        }
        
        const {status , data} = error.response;
        return res.status(status).json(data)
        // console.log(error)

        

    }

}