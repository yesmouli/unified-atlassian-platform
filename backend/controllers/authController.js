const jwt = require('jsonwebtoken');
const bcrypt = require('bcrypt');
const users = require('../data/users');  // This can be replaced by a database

// Login function
exports.login = async (req, res) => {
  const { email, password } = req.body;
  const user = users.find((u) => u.email === email);

  if (user && bcrypt.compareSync(password, user.password)) {
    const token = jwt.sign({ id: user.id }, process.env.JWT_SECRET, { expiresIn: '1h' });
    return res.status(200).json({ token });
  }
  
  return res.status(401).json({ message: 'Invalid credentials' });
};

// Get user function (protected route)
exports.getUser = (req, res) => {
  const user = users.find((u) => u.id === req.userId);
  
  if (user) {
    return res.status(200).json(user);
  }

  return res.status(404).json({ message: 'User not found' });
};
