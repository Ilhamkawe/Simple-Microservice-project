var express = require("express")
var router = express.Router()

const refreshTokenHandler = require("./handler/refreshtoken")

const {APP_NAME} = process.env

router.post("/", refreshTokenHandler.refreshToken)

module.exports = router

