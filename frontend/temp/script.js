// script.js

// Establish WebSocket connection to the server
const roomName = 'general';  // or dynamically get the room from user selection
const ws = new WebSocket(`ws://localhost:8080/ws?room=${roomName}`);

// Event listener for connection opening
ws.onopen = () => {
    console.log("Connected to WebSocket");
};

// Event listener for receiving messages from WebSocket
ws.onmessage = (event) => {
    const message = JSON.parse(event.data);
    displayMessage(message, 'received');
};

// Event listener for errors
ws.onerror = (error) => {
    console.error("WebSocket error:", error);
};

// Event listener for connection close
ws.onclose = () => {
    console.log("WebSocket connection closed");
};

// Function to send message to WebSocket server
function sendMessage() {
    const inputElement = document.querySelector('input');
    const messageText = inputElement.value.trim();
    
    if (messageText === "") return;  // Don't send empty messages

    const message = {
        text: messageText,
        user: "Logan Smith"  // Replace with actual user info
    };

    // Send the message as a JSON string
    ws.send(JSON.stringify(message));

    // Display the message on the frontend (from sender)
    displayMessage(message, 'sent');
    
    // Clear the input field
    inputElement.value = "";
}

// Event listener for "Send" button
document.querySelector('button').addEventListener('click', sendMessage);

// Optional: Send message on pressing Enter key
document.querySelector('input').addEventListener('keypress', function (e) {
    if (e.key === 'Enter') {
        sendMessage();
    }
});


// Function to display messages in the chat window
function displayMessage(message, messageType) {
    const chatContainer = document.querySelector('.chat-messages');
    const messageElement = document.createElement('div');
    
    messageElement.classList.add('message', messageType);  // 'sent' or 'received'
    messageElement.innerHTML = `
        <p>${message.text}</p>
        <span>${new Date().toLocaleTimeString()}</span>
    `;
    
    chatContainer.appendChild(messageElement);
    
    // Scroll to the bottom of the chat container to show the latest message
    chatContainer.scrollTop = chatContainer.scrollHeight;
}
