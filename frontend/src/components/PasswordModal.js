import React from 'react';
import { Link } from 'react-router-dom';
import { PlusCircle, LogOut } from 'lucide-react';

const PasswordModal = ({ onClose, onConnect, onAlert }) => {
    const handleConnect = (machineId) => {
      // Get the entered password from the input field
      const password = 'your_password';
      onConnect(machineId, password);
      onClose();
    };
  
    return (
      <div className="modal modal-open">
        <div className="modal-box">
          <h3 className="font-bold text-lg">Enter Password</h3>
          <div className="form-control w-full max-w-xs">
            <label className="label">
              <span className="label-text">Password for the machine:</span>
            </label>
            <input type="password" className="input input-bordered w-full max-w-xs" />
          </div>
          <div className="modal-action">
            <button className="btn btn-primary" onClick={() => handleConnect(1)}>
              Connect
            </button>
            <button className="btn btn-ghost" onClick={onClose}>
              Cancel
            </button>
          </div>
        </div>
      </div>
    );
  };

export default PasswordModal;
