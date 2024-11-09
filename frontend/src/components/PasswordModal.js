import React, { useState } from 'react';

const PasswordModal = ({ onClose, onConnect, isLoading }) => {
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleConnect = (e) => {
    e.preventDefault();
    if (!password.trim()) {
      setError('Password is required');
      return;
    }
    onConnect(password);
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-base-100  rounded-lg p-6 w-full max-w-md">
        <div className="mb-4">
          <h2 className="text-xl font-semibold">Enter Password</h2>
        </div>
        
        <form onSubmit={handleConnect} className="space-y-4">
          <div>
            <input
              type="password"
              placeholder="Enter machine password"
              value={password}
              onChange={(e) => {
                setPassword(e.target.value);
                setError('');
              }}
              className={`w-full px-3 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 ${
                error ? 'border-red-500' : 'border-gray-300'
              }`}
            />
            {error && (
              <p className="mt-1 text-sm text-red-500">{error}</p>
            )}
          </div>
          
          <div className="flex justify-end space-x-2">
            <button
              type="button"
              onClick={onClose}
              disabled={isLoading}
              className="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
            >
              Cancel
            </button>
            <button
              type="submit"
              disabled={isLoading}
              className="px-4 py-2 text-sm font-medium text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50"
            >
              {isLoading ? 'Connecting...' : 'Connect'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default PasswordModal;