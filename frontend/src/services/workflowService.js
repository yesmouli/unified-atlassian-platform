import api from '../utils/api';

const getWorkflows = () => {
  return api.get('/api/automation/workflows');
};

export default {
  getWorkflows
};
