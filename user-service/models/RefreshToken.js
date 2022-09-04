module.exports = (Sequelize, DataTypes) => {

    const RefreshToken = Sequelize.define('RefreshToken', {
        id : {
            type : DataTypes.INTEGER, 
            allowNull : false, 
            primaryKey : true,
            autoIncrement : true,
        },
        token : {
            type : DataTypes.STRING, 
            allowNull : false, 
        }, 
        user_id : {
            type : DataTypes.INTEGER, 
            allowNull : false,
        }, 
        createdAt : {
            field : "created_at",
            type : DataTypes.DATE, 
            allowNull : false,
        }, updatedAt : {
            field : "updated_at",
            type : DataTypes.DATE, 
            allowNull : false
        }
    }, { 
        tableName : "refresh_tokens", 
        timestamps : true
    });


    return RefreshToken
}