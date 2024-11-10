import React from 'react';

const MachineList = ({ machines, onConnect, onDelete }) => {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4" id="machine-list">
      {machines.map((machine) => (
        <div key={machine.id} className="card bg-base-100 shadow-xl">
          <div className="card-body">
            <h2 className="card-title">{machine.name}</h2>
            <div className="card-actions justify-end gap-2">
              <button 
                className="btn btn-error" 
                onClick={() => onDelete(machine.id)}
              >
                Delete
              </button>
              <button 
                className="btn btn-primary" 
                onClick={() => onConnect(machine.id)}
              >
                Connect
              </button>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
};

export default MachineList;