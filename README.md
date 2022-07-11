# Ecommerce Chatbot

This document seeks to describe in a general way the architecture, design and components that are part of the chatbot.

## Table of Contents  
[Architecture](#architecture)  
[Components](#components)  
...snip...    
<a name="architecture"/>
## Architecture
The architecture of the solution comprises 4 major components.

Listed below:

* [Frontend bot](#frontend)  
* [Backend bot](#backend) 
* [Chatbot API](#chatbot) 
* [Postgres DB](#postgres) 

The following image shows the architecture diagram of the solution

![alt text](https://github.com/mariomateusg/ecommerce-chatbot/blob/main/Architecture.jpg?raw=true)

In the proposed architecture, the frontend communicates with the backend through websockets where user requests are sent to the backend.

The backend communicates with the Chatbot API to generate the responses to the requests to the users

Finally, the knowledge base that is used by the chatbot is stored in the database

<a name="components"/>

## Components

This section describes the components described in the architecture

<a name="frontend"/>

### Frontend bot

The frontend is responsible for interacting with the user.

The component was developed using the React framework.

The communication between the frontend and backend is done through websockets.

The decision to use websockets over rest is that they allow for a higher amount of efficiency compared to REST because they do not require the HTTP request/response overhead for each message sent and received.


The frontend source code can be found [here](https://github.com/mariomateusg/frontend-chatbot)

<a name="backend"/>

### Backend bot

The backend is responsible for deploying the web server, where connections are established via websocket.

Requests received by the frontend are processed and then the Chatbot API is used to answer user questions.

The backend allows the connection of several users.

The backend source code can be found [here](https://github.com/mariomateusg/backend-chatbot)
