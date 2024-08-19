## SuSHI (Work in Progress)

SuSHI is a web-based platform that helps to make SSH connections to remote machines from any location, with a browser based terminal. 

### Features
- Secure SSH Connections: Connect to remote machines securely using SSH.
- Web-Based Terminal: Access the terminal interface directly in your web browser.
- Real-Time Communication: Utilize WebSockets for real-time communication between the client and server.
- User-Friendly Interface: Simple and intuitive interface for ease of use.
- Encrypted Private Keys: Private keys are stored in database securely using AES-CFB encryption with a key derived from PBKDF2-HMAC-SHA256 (10000 iterations), utilizing a unique salt and IV for each encryption operation.


### Screenshot
![screenhot 1](./static/images/homepage.png)
![screenhot 2](./static/images/dashboard.png)
![screenhot 3](./static/images/connected-terminals.png)
![screenhot 5](./static/images/add-machine.png)
![screenhot 4](./static/images/enter-password.png)


## Running app
- `mkdir db`
- `mkdir db/data`
- `sudo chown -R 1001:1001 ./db/data`
- Have the docker-compose.yaml file with all the env values
- `docker compose up`
- By default the server will be up on `localhost:8080`


- It you have issues `connection closed` while opening terminal change the `wss://` to `ws://` in the `static/terminal.html` (line no. 95)
