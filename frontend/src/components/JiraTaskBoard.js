import React, { useState, useEffect } from 'react';
import jiraService from '../services/jiraService';

const JiraTaskBoard = () => {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    jiraService.getTasks()
      .then(response => setTasks(response.data))
      .catch(error => console.error('Error fetching Jira tasks:', error));
  }, []);

  return (
    <div className="jira-task-board">
      <h2>Jira Task Board</h2>
      <ul>
        {tasks.map(task => (
          <li key={task.id}>
            {task.summary} - {task.status}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default JiraTaskBoard;
