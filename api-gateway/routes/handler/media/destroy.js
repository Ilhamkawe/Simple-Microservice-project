const apiAdapter  = require("../../apiAdapter")

const {URL_SERVICE_MEDIA} = process.env

const api = apiAdapter(URL_SERVICE_MEDIA)

module.exports = async (req, res) => {

    try {
        const id = req.params.id;
        const media = await api.delete(`/media/${id}`)
        return res.json(media.data)
    }catch(error){

        if (error.code == "ECONNREFUSED"){
            return res.status(500).json({status : "error", message : "service unavailable"})
        }

        const {status, data} = error.response
        return res.status(status).json(data)

    }
}