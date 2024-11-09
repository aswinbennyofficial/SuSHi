import React, { useState, useEffect, useRef } from 'react';
import { useParams } from 'react-router-dom';
import { Terminal as XTerm } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import { WebLinksAddon } from 'xterm-addon-web-links';
import { AttachAddon } from 'xterm-addon-attach';
import 'xterm/css/xterm.css';

const Terminal = () => {
  const { uuid } = useParams();
  const [activeTab, setActiveTab] = useState(0);
  const [terminals, setTerminals] = useState([]);
  const terminalRef = useRef(null);
  const wsRef = useRef(null);
  const xtermRef = useRef(null);

  // Initialize terminal
  useEffect(() => {
    if (!terminalRef.current) return;

    // Initialize xterm.js
    const term = new XTerm({
      cursorBlink: true,
      cursorStyle: 'block',
      fontSize: 14,
      fontFamily: 'Menlo, Monaco, "Courier New", monospace',
      theme: {
        background: '#1e1e1e',
        foreground: '#ffffff',
      },
      allowTransparency: true,
      scrollback: 1000,
      cols: 80,
      rows: 24,
    });

    // Create addons
    const fitAddon = new FitAddon();
    const webLinksAddon = new WebLinksAddon();

    // Load addons
    term.loadAddon(fitAddon);
    term.loadAddon(webLinksAddon);

    // Open terminal in the container
    term.open(terminalRef.current);
    
    // Connect to WebSocket
    const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
    const ws = new WebSocket(`${protocol}//${window.location.host}/api/v1/terminal/${uuid}`);
    
    // Create and load attach addon after WebSocket connection
    ws.onopen = () => {
      const attachAddon = new AttachAddon(ws);
      term.loadAddon(attachAddon);
      // Initial terminal fit
      fitAddon.fit();
      // Show connection message
      term.write('\x1b[1;32mConnected to terminal.\x1b[0m\r\n');
    };

    // Handle WebSocket errors
    ws.onerror = (error) => {
      term.write('\x1b[1;31mWebSocket connection error.\x1b[0m\r\n');
      console.error('WebSocket error:', error);
    };

    // Handle WebSocket close
    ws.onclose = () => {
      term.write('\x1b[1;31mDisconnected from terminal.\x1b[0m\r\n');
    };

    // Store references
    wsRef.current = ws;
    xtermRef.current = term;

    // Handle window resize
    const handleResize = () => {
      fitAddon.fit();
      // Send new dimensions to server if needed
      const dimensions = { cols: term.cols, rows: term.rows };
      if (ws.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify({ type: 'resize', ...dimensions }));
      }
    };

    window.addEventListener('resize', handleResize);

    // Copy on selection
    term.onSelectionChange(() => {
      if (term.hasSelection()) {
        const selection = term.getSelection();
        navigator.clipboard.writeText(selection);
      }
    });

    // Paste on Ctrl+V or Command+V
    term.attachCustomKeyEventHandler((event) => {
      if ((event.ctrlKey || event.metaKey) && event.key === 'v' && event.type === 'keydown') {
        navigator.clipboard.readText().then(text => {
          ws.send(text);
        });
        return false;
      }
      return true;
    });

    // Cleanup function
    return () => {
      window.removeEventListener('resize', handleResize);
      if (wsRef.current) {
        wsRef.current.close();
      }
      if (xtermRef.current) {
        xtermRef.current.dispose();
      }
    };
  }, [uuid]);

  // Handle adding new terminal tab
  const addNewTerminal = () => {
    const newTabIndex = terminals.length;
    setTerminals([...terminals, { id: newTabIndex }]);
    setActiveTab(newTabIndex);
  };

  return (
    <div className="flex flex-col h-screen">
      <div className="flex items-center border-b border-gray-200 bg-gray-800">
        <div className="flex space-x-2 p-2">
          <button
            className={`px-4 py-2 rounded-t-lg transition-colors duration-200 ${
              activeTab === 0 
                ? 'bg-blue-500 text-white' 
                : 'bg-gray-700 text-gray-300 hover:bg-gray-600'
            }`}
            onClick={() => setActiveTab(0)}
          >
            Terminal 1
          </button>
          {terminals.map((term, index) => (
            <button
              key={term.id}
              className={`px-4 py-2 rounded-t-lg transition-colors duration-200 ${
                activeTab === index + 1
                  ? 'bg-blue-500 text-white'
                  : 'bg-gray-700 text-gray-300 hover:bg-gray-600'
              }`}
              onClick={() => setActiveTab(index + 1)}
            >
              Terminal {index + 2}
            </button>
          ))}
        </div>
        <button 
          className="ml-auto p-2 text-white hover:bg-gray-700 rounded-lg mr-2 transition-colors duration-200"
          onClick={addNewTerminal}
        >
          +
        </button>
      </div>
      
      <div className="flex-1 bg-[#1e1e1e]">
        <div
          ref={terminalRef}
          className="h-full"
          style={{ padding: '12px' }}
        />
      </div>
    </div>
  );
};

export default Terminal;