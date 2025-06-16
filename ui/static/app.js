const chatContainer = document.getElementById('chat-container');
const userInput = document.getElementById('user-input');
const sendButton = document.getElementById('send-button');

sendButton.addEventListener('click', sendMessage);
userInput.addEventListener('keypress', function (e) {
    if (e.key === 'Enter') {
        sendMessage();
    }
});

function sendMessage() {
    const message = userInput.value.trim();
    if (message) {
        appendMessage('You', message);
        userInput.value = '';
        fetchResponse(message);
    }
}

function appendMessage(sender, message) {
    const messageElement = document.createElement('div');
    messageElement.classList.add('message');
    messageElement.innerHTML = `<strong>${sender}:</strong> ${message}`;
    chatContainer.appendChild(messageElement);
    chatContainer.scrollTop = chatContainer.scrollHeight; // Scroll to the bottom
}

function fetchResponse(message) {
    fetch('/api/chat', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.reply) {
            appendMessage('Claude', data.reply);
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });
}