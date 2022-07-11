# Ecommerce Chatbot

This document seeks to describe in a general way the architecture, design and components that are part of the chatbot.

## Table of Contents  
[Architecture](#architecture)  
[Emphasis](#emphasis)  
...snip...    
<a name="architecture"/>
## Architecture
The architecture of the solution comprises 4 major components.

Listed below:

* Frontend bot.
* Backend bot.
* Chatbot API.
* Postgres DB.

The following image shows the architecture diagram of the solution

![alt text](https://github.com/mariomateusg/ecommerce-chatbot/blob/main/Architecture.jpg?raw=true)

In the proposed architecture, the frontend communicates with the backend through websockets where user requests are sent to the backend.

The backend communicates with the Chatbot API to generate the responses to the requests to the users

Finally, the knowledge base that is used by the chatbot is stored in the database
