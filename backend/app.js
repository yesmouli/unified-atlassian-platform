const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const authRoutes = require('./routes/authRoutes');
const jiraRoutes = require('./routes/jiraRoutes');
const confluenceRoutes = require('./routes/confluenceRoutes');
const bitbucketRoutes = require('./routes/bitbucketRoutes');
const workflowRoutes = require('./routes/workflowRoutes');
const errorHandler = require('./middleware/errorHandler');
const verifyToken = require('./middleware/verifyToken');

const app = express();

// Middleware
app.use(cors());
app.use(bodyParser.json());

// Routes
app.use('/api/auth', authRoutes);
app.use('/api/jira', verifyToken, jiraRoutes);
app.use('/api/confluence', verifyToken, confluenceRoutes);
app.use('/api/bitbucket', verifyToken, bitbucketRoutes);
app.use('/api/automation', verifyToken, workflowRoutes);

// Error handling
app.use(errorHandler);

// Start the server
const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
