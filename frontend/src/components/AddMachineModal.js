import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { PlusCircle, LogOut } from 'lucide-react';

const AddMachineModal = ({ onClose, onAddMachine, onAlert }) => {
    const handleAddMachine = () => {
      // Get the form data from the input fields
      const machineData = {
        name: 'New Machine',
        username: 'your_username',
        hostname: 'your_hostname',
        port: 22,
        privateKey: 'your_private_key',
        passphrase: 'your_passphrase',
        password: 'your_password',
        organization: 'your_organization',
      };
      onAddMachine(machineData);
      onClose();
    };
  
    return (
      <div className="modal modal-open">
        <div className="modal-box">
          <h3 className="font-bold text-lg">Add New Machine</h3>
          <div className="form-control w-full">
            {/* Form fields for the new machine */}
          </div>
          <div className="modal-action">
            <button className="btn btn-primary" onClick={handleAddMachine}>
              Add Machine
            </button>
            <button className="btn btn-ghost" onClick={onClose}>
              Cancel
            </button>
          </div>
        </div>
      </div>
    );
  };

export default AddMachineModal;
