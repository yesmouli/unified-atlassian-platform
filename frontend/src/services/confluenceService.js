import api from '../utils/api';

const getDocuments = () => {
  return api.get('/api/confluence/docs');
};

export default {
  getDocuments
};
