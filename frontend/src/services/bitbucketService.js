import api from '../utils/api';

const getRepositories = () => {
  return api.get('/api/bitbucket/repos');
};

export default {
  getRepositories
};
