import React, { useState } from 'react';

const AddMachineModal = ({ onClose, onAddMachine, onAlert }) => {
  const [formData, setFormData] = useState({
    name: '',
    username: '',
    hostname: '',
    port: '22',
    private_key: '',
    passphrase: '',
    password: '',
    organization: ''
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    
    if (!formData.name || !formData.hostname || !formData.username) {
      onAlert('Name, hostname, and username are required', 'error');
      return;
    }

    if (!formData.private_key && !formData.password) {
      onAlert('Either private key or password is required', 'error');
      return;
    }

    onAddMachine(formData);
  };

  return (
    <div className="modal-box relative w-full max-w-2xl bg-base-100 p-6 rounded-lg shadow-xl">
      <h3 className="font-bold text-lg mb-4">Add New Machine</h3>
      
      <form onSubmit={handleSubmit} className="form-control w-full space-y-4">
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="label">
              <span className="label-text">Machine Name*</span>
            </label>
            <input
              type="text"
              name="name"
              value={formData.name}
              onChange={handleInputChange}
              placeholder="My Server"
              className="input input-bordered w-full"
            />
          </div>

          <div>
            <label className="label">
              <span className="label-text">Organization</span>
            </label>
            <input
              type="text"
              name="organization"
              value={formData.organization}
              onChange={handleInputChange}
              placeholder="My Organization"
              className="input input-bordered w-full"
            />
          </div>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="label">
              <span className="label-text">Hostname*</span>
            </label>
            <input
              type="text"
              name="hostname"
              value={formData.hostname}
              onChange={handleInputChange}
              placeholder="example.com"
              className="input input-bordered w-full"
            />
          </div>

          <div>
            <label className="label">
              <span className="label-text">Port</span>
            </label>
            <input
              type="number"
              name="port"
              value={formData.port}
              onChange={handleInputChange}
              placeholder="22"
              className="input input-bordered w-full"
            />
          </div>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="label">
              <span className="label-text">Username*</span>
            </label>
            <input
              type="text"
              name="username"
              value={formData.username}
              onChange={handleInputChange}
              placeholder="username"
              className="input input-bordered w-full"
            />
          </div>

          <div>
            <label className="label">
              <span className="label-text">Password</span>
            </label>
            <input
              type="password"
              name="password"
              value={formData.password}
              onChange={handleInputChange}
              placeholder="Enter password"
              className="input input-bordered w-full"
            />
          </div>
        </div>

        <div>
          <label className="label">
            <span className="label-text">Private Key</span>
          </label>
          <textarea
            name="private_key"
            value={formData.private_key}
            onChange={handleInputChange}
            placeholder="Paste your private key here"
            className="textarea textarea-bordered w-full h-24"
          />
        </div>

        <div>
          <label className="label">
            <span className="label-text">Passphrase</span>
          </label>
          <input
            type="password"
            name="passphrase"
            value={formData.passphrase}
            onChange={handleInputChange}
            placeholder="Private key passphrase (if any)"
            className="input input-bordered w-full"
          />
        </div>

        <div className="modal-action">
          <button type="submit" className="btn btn-primary">
            Add Machine
          </button>
          <button type="button" className="btn btn-ghost" onClick={onClose}>
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
};

export default AddMachineModal;