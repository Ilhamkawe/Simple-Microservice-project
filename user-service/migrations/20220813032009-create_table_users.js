'use strict';

module.exports = {
  async up (queryInterface, Sequelize) {
    /**
     * Add altering commands here.
     *
     * Example:
     * await queryInterface.createTable('users', { id: Sequelize.INTEGER });
     */

    await queryInterface.createTable('users', {
      id : {
        type : Sequelize.INTEGER, 
        alloNull : false, 
        primaryKey : true, 
        autoIncrement : true,
      }, 
      name : {
        type : Sequelize.STRING, 
        allowNull : false, 
      }, 
      profession : {
        type : Sequelize.STRING, 
        allowNull : true,
      }, 
      avatar : {
        type : Sequelize.STRING, 
        allowNull : true,
      },
      role : {
        type : Sequelize.ENUM, 
        values : ["admin", "student"],
        allowNull : false,
      }, 
      email : {
        type : Sequelize.STRING, 
        allowNull : false,
      }, 
      password  : {
        type : Sequelize.STRING, 
        allowNull : false,
      },
      created_at :{
        type : Sequelize.DATE, 
        allowNull : false,
      }, 
      updated_at :{
        type : Sequelize.DATE, 
        allowNull : false,
      }
    })

    await queryInterface.addConstraint('users', {
      type : "unique",
      name : "UNIQUE_USER_EMAIL",
      fields : ["email"] 
    })

  },

  async down (queryInterface, Sequelize) {
    /**
     * Add reverting commands here.
     *
     * Example:
     */
    await queryInterface.dropTable('users');
  }
};
