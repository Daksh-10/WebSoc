# Event push (Websockets) for data retrieval (REST API)

**Problem Statement**

In building UI for high throughput systems, we need to use a mix of techniques without which
the UI becomes very busy and a crash of the browser is likely. In this assignment, build a simple
system where the websocket based UI will retrieve data from the backend that you have.

Set a counter on the backend which gets updated on the UI in real time connected via
websockets. The backend pushes data to the UI say once every 2 seconds. This websocket
payload is used by a REST API call in the UI to retrieve a random string (10 characters) which is generated along with the counter in the backend. So, it’s a bidirectional data flow where a
push will lead to a pull. Every time the data comes in, it’s updated in the UI as shown below.

Websocket is a push from the backend. REST API is a pull from the UI.

Backend: Go
Frontend: React, NextJS

**Installation**
**Backend :**
1. Clone the repository
2. Change Directory to backend ```cd backend```
3. Install the dependencies ```go mod tidy```
4. Run the program ```go run main.go```

**Frontend :**
1. Change Directory to frontend ```cd frontend```
2. Install the dependencies ```npm i```
3. Run the program ```npm run dev```
