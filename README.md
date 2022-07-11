# Ecommerce Chatbot

This document seeks to describe in a general way the architecture, design and components that are part of the chatbot.

## Table of Contents  
[Architecture](#architecture)  
[Components](#components)  
[Installation](#installation)   
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

The frontend is responsible for interacting with the user. The component was developed using the React framework.
The communication between the frontend and backend is done through websockets.
The decision to use websockets over rest is that they allow for a higher amount of efficiency compared to REST because they do not require the HTTP request/response overhead for each message sent and received.

The frontend source code can be found [here](https://github.com/mariomateusg/frontend-chatbot)

<a name="backend"/>

### Backend bot

The backend is responsible for deploying the web server, where connections are established via websocket. Requests received by the frontend are processed and then the Chatbot API is used to answer user questions. The backend allows the connection of several users.

The backend source code can be found [here](https://github.com/mariomateusg/backend-chatbot)

<a name="chatbot"/>

### Chatbot API

The chatbot API contains the business logic to respond to user requests received by the Backend. The communication between the backend and the API is done through 2 methods:

* Greeting: 
* Reply


For the analysis of the requests, the NLP library developed in Go was used, which through a knowledge base allows semantic analysis and generate responses to requests. The documentation of this library can be found at this [link](https://github.com/james-bowman/nlp)

The chatbot API source code can be found [here](https://github.com/mariomateusg/ecommerce-chatbot)

<a name="postgres"/>

### Postgres DB

The knowledge base that allows the Chabot API to be trained to respond to user requests is stored in the database.

<a name="Installation"/>

## Installation

The deployment must be done on kubernetes. For the installation you must execute the following commands:

* kubectl apply -f postgres-config.yaml
* kubectl apply -f postgres-secret.yaml
* kubectl apply -f postgres.yaml

After the creation of the database. The script for creating the tables and inserting the data must be executed. After the creation of the database. The script for creating the tables and inserting the data must be executed. The script is [here](https://github.com/mariomateusg/ecommerce-chatbot/blob/main/database/models.sql)

Once the script is executed, continue with the following commands:


