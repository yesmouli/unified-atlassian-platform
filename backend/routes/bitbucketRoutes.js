const express = require('express');
const router = express.Router();
const bitbucketController = require('../controllers/bitbucketController');

// Get Bitbucket repositories
router.get('/repos', bitbucketController.getRepositories);

module.exports = router;
