import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { PlusCircle, LogOut } from 'lucide-react';


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