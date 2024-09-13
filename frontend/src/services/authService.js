import api from '../utils/api';

const login = (credentials) => {
  return api.post('/api/auth/login', credentials);
};

export default {
  login
};
