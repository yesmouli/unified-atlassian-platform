const axios = require('axios');

// Function to get workflows from the workflow automation API
exports.getWorkflows = async (req, res) => {
  try {
    const response = await axios.get(`${process.env.WORKFLOW_API_URL}/workflows`, {
      headers: {
        Authorization: `Bearer ${process.env.WORKFLOW_API_TOKEN}`
      }
    });

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ message: 'Error fetching workflows' });
  }
};
