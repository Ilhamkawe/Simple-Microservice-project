'use strict';

module.exports = {
  async up (queryInterface, Sequelize) {
    /**
     * Add altering commands here.
     *
     * Example:
     * await queryInterface.createTable('users', { id: Sequelize.INTEGER });
     */

      await queryInterface.createTable('refresh_tokens', {
        id : {
          type : Sequelize.INTEGER, 
          primaryKey : true,
          autoIncrement : true,
          allowNull : false,
        }, 
        token : {
          type : Sequelize.TEXT, 
          allowNull : false, 
        }, 
        user_id : {
          type : Sequelize.INTEGER, 
          allowNull : false,
        }, 
        created_at : {
          type : Sequelize.DATE, 
          allowNull : false, 
        }, 
        updated_at : {
          type : Sequelize.DATE, 
          allowNull : false,
        }
      })

      await queryInterface.addConstraint('refresh_tokens', {
        type : "foreign key", 
        name : "REFRESH_TOKEN__USER_ID", 
        fields : ["user_id"], 
        references : {
          table : "users",
          field : "id"
        }
      });
  },

  async down (queryInterface, Sequelize) {
    /**
     * Add reverting commands here.
     *
     * Example:
     * await queryInterface.dropTable('users');
     */
  }
};
