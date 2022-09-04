var express = require('express');
var router = express.Router();

const refreshTokenHandler = require('./handler/refreshtokens')

router.get("/", refreshTokenHandler.getToken)
router.post("/", refreshTokenHandler.create)

module.exports = router;
