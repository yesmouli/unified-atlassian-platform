const axios = require('axios');

// Function to get documents from Confluence
exports.getDocuments = async (req, res) => {
  try {
    const response = await axios.get(`${process.env.CONFLUENCE_API_URL}/documents`, {
      headers: {
        Authorization: `Bearer ${process.env.CONFLUENCE_API_TOKEN}`
      }
    });

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ message: 'Error fetching Confluence documents' });
  }
};
