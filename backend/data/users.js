// This is a mock data structure; replace with a database
const bcrypt = require('bcrypt');

const users = [
  {
    id: 1,
    email: 'john.doe@example.com',
    password: bcrypt.hashSync('password123', 10),  // Store hashed passwords
    name: 'John Doe'
  },
  {
    id: 2,
    email: 'jane.smith@example.com',
    password: bcrypt.hashSync('mypassword', 10),
    name: 'Jane Smith'
  }
];

module.exports = users;
