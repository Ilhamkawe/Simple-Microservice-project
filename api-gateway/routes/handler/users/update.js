const apiAdapter = require("../../apiAdapter")

const {
    URL_SERVICE_USERS
} = process.env

const api = apiAdapter(URL_SERVICE_USERS)

module.exports = async (req, res) => {
    try { 
        const id = req.user.data.id
        const user = await api.put(`/users/update/${id}`, req.body)
        return res.json({
            message : "success", 
            data : user.data
        })
    } catch (error) {
        if(error.code == "ECONNREFUSED"){
            return res.status(500).json({status : "error", message : "service tidak teresedia"})
        }
        
        const {status , data} = error.response;
        return res.status(status).json(data)
    }
}