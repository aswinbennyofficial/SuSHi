import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import Navbar from '../components/Navbar';
import MachineList from '../components/MachineList';
import PasswordModal from '../components/PasswordModal';
import AddMachineModal from '../components/AddMachineModal';
import Alert from '../components/Alert';

const Dashboard = () => {
  const navigate = useNavigate();
  const [showPasswordModal, setShowPasswordModal] = useState(false);
  const [showAddMachineModal, setShowAddMachineModal] = useState(false);
  const [machines, setMachines] = useState([]);
  const [alerts, setAlerts] = useState([]);
  const [currentMachineId, setCurrentMachineId] = useState(null);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    fetchMachines();
  }, []);

  const fetchMachines = async () => {
    try {
      const response = await axios.get('/api/v1/machines');
      if (response.status === 200 && response.data.data === null) {
        addAlert('No machines found', 'info');
      } else if (response.status === 200) {
        setMachines(response.data.data);
      } else {
        addAlert('Failed to fetch machines', 'error');
      }
    } catch (error) {
      addAlert('Failed to fetch machines', 'error');
    }
  };

  const handleLogout = () => {
    navigate('/');
  };

  const handlePasswordModal = (machineId) => {
    setCurrentMachineId(machineId);
    setShowPasswordModal(!showPasswordModal);
  };

  const handleAddMachineModal = () => {
    setShowAddMachineModal(!showAddMachineModal);
  };

  const addMachine = async (machineData) => {
    try {
      await axios.post('/api/v1/machine', machineData);
      addAlert('Machine added successfully', 'success');
      fetchMachines();
      setShowAddMachineModal(false); // Close modal after successful addition
    } catch (error) {
      addAlert('Failed to add machine', 'error');
    }
  };

  const connectToMachine = async (password) => {
    if (!password) {
      addAlert('Password is required to connect to the machine', 'error');
      return;
    }

    setIsLoading(true);
    try {
      console.log("Machine id : ",currentMachineId);
      console.log("Connecting with password:", password);

      const response = await axios.post(`/api/v1/machine/${currentMachineId}/connect`, {
        password: password
      }, {
        headers: {
          'Content-Type': 'application/json'
        },
        withCredentials: true
      });

      if (response.data.status === 'OK') {
        const uuid = response.data.data;
        window.open(`/terminal/${uuid}`, '_blank', 'width=800,height=600');
        addAlert('Connected successfully', 'success');
        setShowPasswordModal(false);
      } else {
        addAlert(`Failed to connect to machine: ${response.data.message}`, 'error');
      }
    } catch (error) {
      addAlert(`Error connecting to machine: ${error.message}`, 'error');
    } finally {
      setIsLoading(false);
    }
  };

  const deleteMachine = async (machineId) => {
    try {
      const response = await axios.delete(`/api/v1/machine/${machineId}`);
      if (response.status === 200) {
        addAlert('Machine deleted successfully', 'success');
        // Refresh the machine list after deletion
        fetchMachines();
      } else {
        addAlert('Failed to delete machine', 'error');
      }
    } catch (error) {
      addAlert(`Error deleting machine: ${error.message}`, 'error');
    }
  };

  const addAlert = (message, type) => {
    const newAlert = { message, type, id: Date.now() };
    setAlerts(prev => [...prev, newAlert]);
    // Remove alert after 5 seconds
    setTimeout(() => {
      setAlerts(prev => prev.filter(alert => alert.id !== newAlert.id));
    }, 5000);
  };

  return (
    <div className="h-screen flex flex-col">
      <Navbar onLogout={handleLogout} onAddMachine={handleAddMachineModal} />
      <div className="container mx-auto p-4 flex-1">
        <h1 className="text-3xl font-bold mb-4">Dashboard</h1>
        <MachineList 
        machines={machines} 
        onConnect={handlePasswordModal} 
        onDelete={deleteMachine}  
      />
        
        {showPasswordModal && (
          <PasswordModal
            onClose={() => setShowPasswordModal(false)}
            onConnect={connectToMachine}
            isLoading={isLoading}
          />
        )}
        
        {showAddMachineModal && (
          <div className="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center">
            <AddMachineModal
              onClose={() => setShowAddMachineModal(false)}
              onAddMachine={addMachine}
              onAlert={addAlert}
            />
          </div>
        )}

        <div className="fixed bottom-4 right-4 z-50 flex flex-col gap-2">
          {alerts.map((alert) => (
            <Alert 
              key={alert.id} 
              message={alert.message} 
              type={alert.type} 
            />
          ))}
        </div>
      </div>
    </div>
  );
};

export default Dashboard;