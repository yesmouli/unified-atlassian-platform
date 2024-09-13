import React, { useState, useEffect } from 'react';
import JiraTaskBoard from './JiraTaskBoard';
import ConfluenceDocs from './ConfluenceDocs';
import BitbucketRepoView from './BitbucketRepoView';
import WorkflowAutomation from './WorkflowAutomation';

const Dashboard = () => {
  const [userData, setUserData] = useState(null);

  useEffect(() => {
    // Example call to the backend to fetch user data
    fetch('/api/auth/user', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`
      }
    })
      .then(response => response.json())
      .then(data => setUserData(data))
      .catch(error => console.error('Error fetching user data:', error));
  }, []);

  if (!userData) return <div>Loading...</div>;

  return (
    <div className="dashboard">
      <h1>Welcome, {userData.name}</h1>
      <JiraTaskBoard />
      <ConfluenceDocs />
      <BitbucketRepoView />
      <WorkflowAutomation />
    </div>
  );
};

export default Dashboard;
