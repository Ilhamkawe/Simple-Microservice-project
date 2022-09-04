var express = require('express');
var router = express.Router();

const verifyToken = require("../middlewares/verifytoken")

const userHandler = require("./handler/users")

const { APP_NAME } = process.env;

/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});

router.post("/register", userHandler.register)
router.post("/login", userHandler.login)
router.get("/profile", verifyToken, userHandler.profile)
router.put("/update", verifyToken, userHandler.update)
router.delete("/delete", verifyToken, userHandler.logout)

module.exports = router;
