import React, { useState, useEffect } from 'react';
import workflowService from '../services/workflowService';

const WorkflowAutomation = () => {
  const [workflows, setWorkflows] = useState([]);

  useEffect(() => {
    workflowService.getWorkflows()
      .then(response => setWorkflows(response.data))
      .catch(error => console.error('Error fetching workflows:', error));
  }, []);

  return (
    <div className="workflow-automation">
      <h2>Workflow Automation</h2>
      <ul>
        {workflows.map(workflow => (
          <li key={workflow.id}>
            {workflow.name} - {workflow.status}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default WorkflowAutomation;
