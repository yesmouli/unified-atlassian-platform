import React, { useState, useEffect } from 'react';
import confluenceService from '../services/confluenceService';

const ConfluenceDocs = () => {
  const [docs, setDocs] = useState([]);

  useEffect(() => {
    confluenceService.getDocuments()
      .then(response => setDocs(response.data))
      .catch(error => console.error('Error fetching Confluence documents:', error));
  }, []);

  return (
    <div className="confluence-docs">
      <h2>Confluence Documents</h2>
      <ul>
        {docs.map(doc => (
          <li key={doc.id}>
            <a href={doc.link}>{doc.title}</a>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ConfluenceDocs;
