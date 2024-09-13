const express = require('express');
const router = express.Router();
const workflowController = require('../controllers/workflowController');

// Get workflow automations
router.get('/workflows', workflowController.getWorkflows);

module.exports = router;
