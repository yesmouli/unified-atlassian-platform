const express = require('express');
const router = express.Router();
const jiraController = require('../controllers/jiraController');

// Get Jira tasks
router.get('/tasks', jiraController.getTasks);

module.exports = router;
