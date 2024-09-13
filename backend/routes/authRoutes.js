const express = require('express');
const router = express.Router();
const authController = require('../controllers/authController');

// Login route
router.post('/login', authController.login);

// Protected route to fetch user info
router.get('/user', authController.getUser);

module.exports = router;
