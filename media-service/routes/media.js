var express = require("express");
var router = express.Router();
var isBase64 = require("is-Base64")
var base64Img = require("base64-img")

const {Media} = require("../models");

const fs = require("fs");

// const { json } = require("sequelize/types");
// const { now } = require("sequelize/types/utils");

// menghapus gambar berdasarkan id 
router.delete("/:id", async (req, res)=>{
    const id = req.params.id;
    console.log(id);

    const media = await Media.findByPk(id);
    if (!media){
        return res.status(400).json({
            status  : "error", 
            message : "Invalid Id"
        });
    }
    // console.log(media.Image)
    fs.unlink(`./public/${media.Image}`, async (err) => {
        if (err){
            return  res.status(400).json({
                    status : "error", 
                    message : "Cannot Delete Image"
                });
        } 
    });

    await media.destroy();

    return res.status(200).json({
        status : "sukses", 
        message : "Data Media Berhasil Dihapus"
    });

})

//menampilkan gambar
router.get("/", async (req, res) => {
    const media = await Media.findAll({attributes : ["id", "image"]});
    
    const mappedMedia = media.map((m) => {
        m.dataValues.image = `${req.get("host")}/${m.dataValues.image}`;
        return m;
    });


    return res.status(200).json({
        status : "sukses", 
        images : mappedMedia
    });
});

// upload gambar
router.post("/", (req, res) => {
    const image = req.body.image;

    if(!isBase64(image, {mimeRequired : true})){
        return res.status(400).json({status : "error", message: "Invalid Base64"});
    }

    base64Img.img(image,"./public/images", Date.now(),async (err, filePath) => {

        if (err){
            return res.status(400).json({status : "error", message : err.message})
        }

        const fileName = filePath.split("\\").pop().split("/").pop();
        
        const media = await Media.create({Image: `images/${fileName}`});

        return res.json({
            status : "Sukses",
            data : {
                id : media.id, 
                image : `${req.get('host')}/images/${fileName}`
            }
        });

    })
    


})

module.exports = router;