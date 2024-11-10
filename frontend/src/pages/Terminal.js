import React, { useEffect, useRef, useState } from 'react';
import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import { WebLinksAddon } from 'xterm-addon-web-links';
import { SearchAddon } from 'xterm-addon-search';
import { X, Maximize2, Minimize2 } from 'lucide-react';
import 'xterm/css/xterm.css';
import { useParams } from 'react-router-dom';

const TerminalComponent = () => {
  const [terminals, setTerminals] = useState({});
  const [wsConnections, setWsConnections] = useState({});
  const [activeTab, setActiveTab] = useState(0);
  const [tabCount, setTabCount] = useState(1);
  const [isFullscreen, setIsFullscreen] = useState(false);
  const terminalRefs = useRef({});
  const fitAddons = useRef({});
  const { uuid } = useParams();
  const containerRef = useRef(null);

  const createTerminal = (tabId) => {
    if (terminalRefs.current[tabId]) {
      return;
    }

    const term = new Terminal({
      fontSize: 14,
      fontFamily: 'Menlo, Monaco, "Courier New", monospace',
      theme: {
        background: '#1a1b26',
        foreground: '#a9b1d6',
        cursor: '#c0caf5',
        selection: '#33467C',
        black: '#32344a',
        brightBlack: '#444b6a',
        red: '#f7768e',
        brightRed: '#ff7a93',
        green: '#9ece6a',
        brightGreen: '#b9f27c',
        yellow: '#e0af68',
        brightYellow: '#ff9e64',
        blue: '#7aa2f7',
        brightBlue: '#7da6ff',
        magenta: '#ad8ee6',
        brightMagenta: '#bb9af7',
        cyan: '#449dab',
        brightCyan: '#0db9d7',
        white: '#787c99',
        brightWhite: '#acb0d0'
      },
      cursorBlink: true,
      scrollback: 10000,
      allowTransparency: true
    });

    const fitAddon = new FitAddon();
    const searchAddon = new SearchAddon();
    const webLinksAddon = new WebLinksAddon();

    term.loadAddon(fitAddon);
    term.loadAddon(searchAddon);
    term.loadAddon(webLinksAddon);

    fitAddons.current[tabId] = fitAddon;

    const container = document.getElementById(`terminal-${tabId}`);
    if (!container) return;

    term.open(container);
    fitAddon.fit();

    const host = window.location.host;
    const wsURL = `ws://${host}/api/v1/ssh?uuid=${uuid || noUUID()}`;
    const ws = new WebSocket(wsURL);

    ws.onopen = () => {
      term.write('\x1b[1;34mConnected to SuSHI v0...\x1b[0m\r\n\r\n');
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
      term.write('\x1b[1;31mWebSocket error: ${event}\x1b[0m\r\n');
    };

    ws.onclose = () => {
      term.write('\x1b[1;31mConnection closed.\x1b[0m\r\n');
    };

    terminalRefs.current[tabId] = term;
    setTerminals(prev => ({ ...prev, [tabId]: term }));
    setWsConnections(prev => ({ ...prev, [tabId]: ws }));
  };

  const startHeartbeat = (ws) => {
    const interval = setInterval(() => {
      if (ws.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify({ type: 'heartbeat', data: '' }));
      }
    }, 5 * 60 * 1000);
    return () => clearInterval(interval);
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

  const closeTab = (tabId, event) => {
    event.stopPropagation();
    if (tabCount === 1) return;

    // Close WebSocket connection
    if (wsConnections[tabId]) {
      wsConnections[tabId].close();
    }

    // Clean up terminal
    if (terminalRefs.current[tabId]) {
      terminalRefs.current[tabId].dispose();
      delete terminalRefs.current[tabId];
      delete fitAddons.current[tabId];
    }

    // Remove from state
    const newTerminals = { ...terminals };
    const newWsConnections = { ...wsConnections };
    delete newTerminals[tabId];
    delete newWsConnections[tabId];
    setTerminals(newTerminals);
    setWsConnections(newWsConnections);

    // Update tab count and active tab
    setTabCount(prev => prev - 1);
    if (activeTab === tabId) {
      setActiveTab(Math.max(0, tabId - 1));
    }
  };

  const switchTab = (tabId) => {
    setActiveTab(tabId);
    setTimeout(() => {
      if (fitAddons.current[tabId]) {
        fitAddons.current[tabId].fit();
      }
    }, 0);
  };

  const toggleFullscreen = () => {
    setIsFullscreen(!isFullscreen);
    setTimeout(() => {
      Object.values(fitAddons.current).forEach(addon => addon.fit());
    }, 100);
  };

  useEffect(() => {
    createTerminal(0);

    const handleResize = () => {
      Object.values(fitAddons.current).forEach(addon => addon.fit());
    };

    window.addEventListener('resize', handleResize);

    return () => {
      window.removeEventListener('resize', handleResize);
      Object.values(wsConnections).forEach(ws => {
        if (ws.readyState === WebSocket.OPEN) {
          ws.close();
        }
      });
      Object.values(terminalRefs.current).forEach(term => {
        term.dispose();
      });
    };
  }, []);

  return (
    <div 
      ref={containerRef}
      className={`bg-base-200 ${isFullscreen ? 'fixed inset-0 z-50' : 'h-screen'} flex flex-col`}
    >
      <div className="flex items-center p-2 bg-gray-900 text-white">
        <div className="flex flex-1 overflow-x-auto">
          {[...Array(tabCount)].map((_, index) => (
            <div
              key={index}
              className={`flex items-center tab px-4 py-2 cursor-pointer hover:bg-gray-700 border-r border-gray-700
                ${activeTab === index ? 'bg-gray-800' : 'bg-gray-900'}`}
              onClick={() => switchTab(index)}
            >
              <span className="mr-2">Terminal {index + 1}</span>
              {tabCount > 1 && (
                <X
                  size={14}
                  className="hover:text-red-400"
                  onClick={(e) => closeTab(index, e)}
                />
              )}
            </div>
          ))}
        </div>
        <div className="flex items-center gap-2 px-2">
          <button
            onClick={addTab}
            className="px-3 py-1 bg-gray-700 hover:bg-gray-600 rounded"
          >
            +
          </button>
          <button
            onClick={toggleFullscreen}
            className="p-1 hover:bg-gray-700 rounded"
          >
            {isFullscreen ? <Minimize2 size={18} /> : <Maximize2 size={18} />}
          </button>
        </div>
      </div>
      <div className="flex-1 relative bg-[#1a1b26]">
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