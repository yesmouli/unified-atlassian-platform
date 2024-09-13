import api from '../utils/api';

const getTasks = () => {
  return api.get('/api/jira/tasks');
};

export default {
  getTasks
};
