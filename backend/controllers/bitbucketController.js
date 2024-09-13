const axios = require('axios');

// Function to get repositories from Bitbucket
exports.getRepositories = async (req, res) => {
  try {
    const response = await axios.get(`${process.env.BITBUCKET_API_URL}/repositories`, {
      headers: {
        Authorization: `Bearer ${process.env.BITBUCKET_API_TOKEN}`
      }
    });

    res.status(200).json(response.data);
  } catch (error) {
    res.status(500).json({ message: 'Error fetching Bitbucket repositories' });
  }
};
