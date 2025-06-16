# My First AI Agent 🕵🏻‍♂️

## Chat Application with Anthropics API

This project is a chat application that allows users to interact with a chat agent powered by the Anthropics API. The application is structured into several components, each serving a specific purpose.

## Project Structure

```
code-editing-agent
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── agent.go        # Implementation of the Agent type
│   └── anthropic.go    # Definitions and methods for the Anthropics API
├── ui
│   ├── static
│   │   ├── index.html   # Main HTML file for the user interface
│   │   ├── style.css     # Styles for the user interface
│   │   └── app.js        # JavaScript code for user interactions
│   └── server.go        # Web server to serve static files and handle communication
├── go.mod               # Go module definition file
├── go.sum               # Checksums for module dependencies
└── README.md            # Documentation for the project
```

## Setup Instructions

1. **Clone the Repository**
   ```bash
   git clone <repository-url>
   cd code-editing-agent
   ```

2. **Install Dependencies**
   Ensure you have Go installed, then run:
   ```bash
   go mod tidy
   ```

3. **Run the Application**
   Start the server by running:
   ```bash
   go run ui/server.go
   ```
   This will serve the static files and allow you to interact with the chat agent.

4. **Access the Chat Interface**
   Open your web browser and navigate to `http://localhost:8080` to start chatting with the agent.

## Usage Guidelines

- Type your messages in the input field and press Enter to send them.
- The chat agent will respond based on the input provided.
- Use 'ctrl-c' in the terminal to stop the server.

## Contributing

Feel free to submit issues or pull requests if you have suggestions or improvements for the project.