import React, { useEffect, useRef, useState } from 'react';
import { Terminal } from 'xterm';
import 'xterm/css/xterm.css';
import { useParams } from 'react-router-dom';

const TerminalComponent = () => {
  const [terminals, setTerminals] = useState({});
  const [wsConnections, setWsConnections] = useState({});
  const [activeTab, setActiveTab] = useState(0);
  const [tabCount, setTabCount] = useState(1);
  const terminalRefs = useRef({});
  const { uuid } = useParams();

  const createTerminal = (tabId) => {
    if (terminalRefs.current[tabId]) {
      return;
    }

    const term = new Terminal();
    const container = document.getElementById(`terminal-${tabId}`);
    if (!container) return;

    term.open(container);

    const host = window.location.host;
    const wsURL = `ws://${host}/api/v1/ssh?uuid=${uuid || noUUID()}`;
    const ws = new WebSocket(wsURL);

    ws.onopen = () => {
      term.write('Connected to SuSHI v0...\r\n\r\n');
      term.prompt = '> ';
      term.onData(data => {
        ws.send(JSON.stringify({ type: 'data', data: data }));
      });
      startHeartbeat(ws);
    };

    ws.onmessage = event => {
      term.write(event.data);
    };

    ws.onerror = event => {
      term.write(`WebSocket error: ${event}\r\n`);
    };

    ws.onclose = () => {
      term.write('Connection closed.\r\n');
    };

    terminalRefs.current[tabId] = term;
    setTerminals(prev => ({ ...prev, [tabId]: term }));
    setWsConnections(prev => ({ ...prev, [tabId]: ws }));
  };

  const startHeartbeat = (ws) => {
    setInterval(() => {
      if (ws.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify({ type: 'heartbeat', data: '' }));
      }
    }, 5 * 60 * 1000); // 5 minutes
  };

  const noUUID = () => {
    console.log('no uuid');
    return '';
  };

  const addTab = () => {
    const newTabId = tabCount;
    setTabCount(prev => prev + 1);
    setTimeout(() => {
      createTerminal(newTabId);
      setActiveTab(newTabId);
    }, 0);
  };

  const switchTab = (tabId) => {
    setActiveTab(tabId);
  };

  useEffect(() => {
    createTerminal(0);
    
    return () => {
      // Cleanup WebSocket connections when component unmounts
      Object.values(wsConnections).forEach(ws => {
        if (ws.readyState === WebSocket.OPEN) {
          ws.close();
        }
      });
    };
  }, []);

  return (
    <div className="bg-base-200 h-screen flex flex-col">
      <div className="flex items-center p-2 bg-[#333] text-white">
        <div className="flex">
          {[...Array(tabCount)].map((_, index) => (
            <a
              key={index}
              className={`tab px-3 py-1 cursor-pointer hover:bg-[#555] ${
                activeTab === index ? 'bg-[#444]' : ''
              }`}
              onClick={() => switchTab(index)}
            >
              Terminal {index + 1}
            </a>
          ))}
        </div>
        <button
          onClick={addTab}
          className="ml-auto px-3 py-1 bg-[#444] hover:bg-[#555]"
        >
          +
        </button>
      </div>
      <div className="flex-1 relative bg-black">
        {[...Array(tabCount)].map((_, index) => (
          <div
            key={index}
            id={`terminal-${index}`}
            className={`absolute inset-0 ${activeTab === index ? 'block' : 'hidden'}`}
          />
        ))}
      </div>
    </div>
  );
};

export default TerminalComponent;