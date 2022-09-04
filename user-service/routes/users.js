var express = require('express');
var router = express.Router();
const userHandler = require('./handler/users');
/* GET users listing. */
router.get('/', function(req, res, next) {
  res.send('respond with a resource');
});

router.post("/register", userHandler.register)
router.post("/login",userHandler.login)
router.put("/update/:id", userHandler.update)
router.get("/:id/detail", userHandler.getUser)
router.get("/list", userHandler.listUser)
router.delete("/logout", userHandler.logout)

module.exports = router;
