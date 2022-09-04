module.exports = (Sequelize, DataTypes) => {

    const Media = Sequelize.define('Media', {
      id : {
        type : DataTypes.INTEGER, 
        primaryKey : true, 
        autoIncrement : true, 
        allowNUll : false,
      }, 
      Image : {
        type : DataTypes.STRING, 
        allowNull : false, 
      }, 
      createdAt: {
        field : 'created_at',
        type : DataTypes.DATE, 
        allowNull : false, 
      }, 
      updatedAt : {
        field : 'updated_at', 
        type : DataTypes.DATE, 
        allowNull : false
      },
    }, {
        tableName : 'media'
      })

    return Media

}