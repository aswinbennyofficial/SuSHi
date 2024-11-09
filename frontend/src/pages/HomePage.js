import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGoogle, faGithub } from '@fortawesome/free-brands-svg-icons';
import axios from 'axios';

const HomePage = () => {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);

  const api = axios.create({
    baseURL: 'http://localhost:8080',  // or whatever your base URL is
    withCredentials: true  // if you're using cookies for authentication
  });

  const handleGoogleAuth = async () => {
    setIsLoading(true);
    try {
      const response = await axios.get('/api/v1/auth/url?to=google');
      console.log(response);
      window.location.href = response.data;
    } catch (error) {
      console.error('Error during Google authentication:', error);
      setIsLoading(false);
    }
  };

  const handleGithubAuth = async () => {
    setIsLoading(true);
    try {
      const response = await axios.get('/api/v1/auth/url?to=github');
     
      window.location.href = response.data;
    } catch (error) {
      console.error('Error during GitHub authentication:', error);
      setIsLoading(false);
    }
  };

  return (
    <div className="bg-base-200 min-h-screen">
      {/* Navbar */}
      <nav className="navbar bg-base-100 shadow-lg p-4">
        <div className="navbar-start">
          <span className="btn btn-ghost normal-case text-xl">SuSHi</span>
        </div>
        <div className="navbar-end">
          <button
            className="btn btn-primary"
            onClick={() => document.getElementById('authModal').showModal()}
          >
            Login / Sign Up
          </button>
        </div>
      </nav>

      {/* Hero Section */}
      <div
        className="hero min-h-screen bg-cover bg-center"
        style={{ backgroundImage: 'url("/images/terminal.png")' }}
      >
        <div className="hero-overlay bg-opacity-60"></div>
        <div className="hero-content text-neutral-content text-center flex flex-col justify-center items-center h-full p-6">
          <div className="max-w-md">
            <h1 className="text-5xl font-bold mb-5">Hello there</h1>
            <p className="mb-5 text-lg">
              Welcome to SuSHi, web-based platform that helps to make SSH connections
              to remote machines from any location, with a browser based terminal.
            </p>
            <button
              className="btn btn-primary"
              onClick={() => document.getElementById('authModal').showModal()}
            >
              Get started
            </button>
          </div>
        </div>
      </div>

      {/* Login/Sign Up Modal */}
      <dialog id="authModal" className="modal">
        <div className="modal-box">
          <h3 className="font-bold text-lg mb-4">Login / Sign Up</h3>
          <div className="flex flex-col gap-4">
            <button
              id="googleAuth"
              className="btn btn-outline btn-primary"
              onClick={handleGoogleAuth}
              disabled={isLoading}
            >
              <FontAwesomeIcon icon={faGoogle} className="h-6 w-6 mr-2" />
              {isLoading ? 'Logging in with Google...' : 'Login with Google'}
            </button>
            <button
              id="githubAuth"
              className="btn btn-outline btn-secondary"
              onClick={handleGithubAuth}
              disabled={isLoading}
            >
              <FontAwesomeIcon icon={faGithub} className="h-6 w-6 mr-2" />
              {isLoading ? 'Logging in with GitHub...' : 'Login with GitHub'}
            </button>
          </div>
        </div>
        <form method="dialog" className="modal-backdrop">
          <button>Close</button>
        </form>
      </dialog>

      {/* Content Section */}
      <div className="container mx-auto px-4 py-8">
        <p className="text-lg mb-8">
          The private keys are encrypted in the database using PBKDF2 derived keys
          derived from the machine-specific password provided by the user and an
          owner-specific salt.
        </p>
      </div>

      {/* Screenshot Slider Section */}
      <div className="container mx-auto px-4 py-8">
        <h2 className="text-3xl font-bold mb-4">Screenshots</h2>
        <div className="carousel w-full">
          <div id="slide1" className="carousel-item relative w-full">
            <img src="/images/dashboard.png" className="w-full" />
            <div className="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
              <Link to="#slide4" className="btn btn-circle">
                ❮
              </Link>
              <Link to="#slide2" className="btn btn-circle">
                ❯
              </Link>
            </div>
          </div>
          {/* Add more carousel items */}
        </div>
      </div>

      {/* Content Section */}
      <div className="container mx-auto px-4 py-8">
        <p className="text-lg mb-8">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam in dui
          mauris. Vivamus hendrerit arcu sed erat molestie vehicula. Sed auctor
          neque eu tellus rhoncus ut eleifend nibh porttitor. Ut in nulla enim.
          Phasellus molestie magna non est bibendum non venenatis nisl tempor.
        </p>
      </div>

      {/* Footer */}
      <footer className="footer bg-neutral text-neutral-content p-10">
        <nav>
          <h6 className="footer-title">Services</h6>
          <Link to="#" className="link link-hover">
            Branding
          </Link>
          <Link to="#" className="link link-hover">
            Design
          </Link>
          <Link to="#" className="link link-hover">
            Marketing
          </Link>
          <Link to="#" className="link link-hover">
            Advertisement
          </Link>
        </nav>
        <nav>
          <h6 className="footer-title">Company</h6>
          <Link to="#" className="link link-hover">
            About us
          </Link>
          <Link to="#" className="link link-hover">
            Contact
          </Link>
          <Link to="#" className="link link-hover">
            Jobs
          </Link>
          <Link to="#" className="link link-hover">
            Press kit
          </Link>
        </nav>
        <nav>
          <h6 className="footer-title">Legal</h6>
          <Link to="#" className="link link-hover">
            Terms of use
          </Link>
          <Link to="#" className="link link-hover">
            Privacy policy
          </Link>
          <Link to="#" className="link link-hover">
            Cookie policy
          </Link>
        </nav>
      </footer>
    </div>
  );
};

export default HomePage;