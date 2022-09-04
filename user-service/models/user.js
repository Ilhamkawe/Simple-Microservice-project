module.exports = (Sequelize, DataTypes) => {
    const User = Sequelize.define('User', {
        id : { 
            type : DataTypes.INTEGER,
            primaryKey : true,
            autoIncrement : true, 
            allowNull : false
        }, 
        name : {
            type : DataTypes.STRING, 
            allowNull : false, 
        }, 
        profession : {
            type : DataTypes.STRING, 
            allowNull : true,
        },
        avatar : {
            type : DataTypes.STRING, 
            alloNull : true,
        }, 
        role : {
            type : DataTypes.ENUM, 
            values : ["admin", "student"], 
            allowNull : false, 
            defaultValue : "student"
        }, 
        email : {
            type : DataTypes.STRING, 
            allowNull : true,
        }, 
        password : {
            type : DataTypes.STRING, 
            allowNull : false,
        },
        createdAt : {
            field : "created_at", 
            allowNull : false,
            type : DataTypes.DATE
        },
        updatedAt : {
            field : "updated_at", 
            allowNull : false,
            type : DataTypes.DATE
        }
    }, {
        tableName : "users", 
        timestamps : true 
    })

    return User
}