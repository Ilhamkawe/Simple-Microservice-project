'use strict';
const bcrypt = require('bcrypt')

module.exports = {
  async up (queryInterface, Sequelize) {
    /**
     * Add seed commands here.
     *
     * Example:
     * await queryInterface.bulkInsert('People', [{
     *   name: 'John Doe',
     *   isBetaMember: false
     * }], {});
    */

    await queryInterface.bulkInsert('users', [
      {
        name : "Muhammad Ilham Kusumawardhana", 
        profession : "Fullstack Developer", 
        role : "admin", 
        email : "kawekaweha321@gmail.com", 
        password : await bcrypt.hash('Bandung00', 10), 
        created_at : new Date(), 
        updated_at : new Date()
      }, {
        name : "Muhamamd Adityo Kusumawardhana", 
        profession : "Mahasiswa", 
        role : "student", 
        email : "kawekaweha123@gmail.com", 
        password : await bcrypt.hash("Rahasia123", 10), 
        created_at : new Date(), 
        updated_at : new Date()
      }
    ]);
  },

  async down (queryInterface, Sequelize) {
    /**
     * Add commands to revert seed here.
     *
     * Example:
     * await queryInterface.bulkDelete('People', null, {});
     */
  }
};
