const express = require('express');
const router = express.Router();
const confluenceController = require('../controllers/confluenceController');

// Get Confluence documents
router.get('/docs', confluenceController.getDocuments);

module.exports = router;
