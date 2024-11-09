import React from 'react';
import { Link } from 'react-router-dom';
import { PlusCircle, LogOut } from 'lucide-react';


const Navbar = ({ onLogout, onAddMachine }) => {
    return (
      <div className="navbar bg-base-100 shadow-lg">
        <div className="navbar-start">
          <Link to="/" className="btn btn-ghost normal-case text-xl">
            SuSHi
          </Link>
        </div>
        <div className="navbar-end">
          <button className="btn btn-ghost" onClick={onAddMachine}>
            <PlusCircle />
            Add machine
          </button>
          <button className="btn btn-ghost" onClick={onLogout}>
            <LogOut />
            Logout
          </button>
        </div>
      </div>
    );
  };

export default Navbar;