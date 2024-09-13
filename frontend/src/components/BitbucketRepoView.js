import React, { useState, useEffect } from 'react';
import bitbucketService from '../services/bitbucketService';

const BitbucketRepoView = () => {
  const [repos, setRepos] = useState([]);

  useEffect(() => {
    bitbucketService.getRepositories()
      .then(response => setRepos(response.data))
      .catch(error => console.error('Error fetching Bitbucket repositories:', error));
  }, []);

  return (
    <div className="bitbucket-repo-view">
      <h2>Bitbucket Repositories</h2>
      <ul>
        {repos.map(repo => (
          <li key={repo.id}>
            <a href={repo.link}>{repo.name}</a>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default BitbucketRepoView;
