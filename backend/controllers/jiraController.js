const axios = require('axios');

// Function to get tasks from Jira
exports.getTasks = async (req, res) => {
  try {
    const response = await axios.get(`${process.env.JIRA_API_URL}/tasks`, {
      headers: {
        Authorization: `Bearer ${process.env.JIRA_API_TOKEN}`
      }
    });

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ message: 'Error fetching Jira tasks' });
  }
};
