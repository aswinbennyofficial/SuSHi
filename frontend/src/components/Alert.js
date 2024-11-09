import React, { useState } from 'react';

const Alert = ({ message, type }) => {
    const alertClass = `alert alert-${type}`;
  
    return (
      <div className={alertClass}>
        <div>
          <span>{message}</span>
        </div>
      </div>
    );
  };

export default Alert;