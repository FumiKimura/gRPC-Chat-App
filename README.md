# Golang + gRPC Terminal Chat App

This was two weeks long project during part-time immersive bootcamp. The main objective was to learn completely new programming language, first programming language was JavaScript, and create application

## Language Requirements

- It must be a language you have never used before
- No visual languages
- TypeScript is not allowed (JavaScript is first language)

## Basic Requirements

- A demo of your app that compiles is mandatory
- Show some of the source code
- Explain the strengths of the new language
- Describe what types of projects it is best suited for
- Compare your new language with JavaScript!
- Share challenges you experienced while learning the new language

## Advanced Requirements

- Test coverage!
- Have both frontend and backend in the new language.

## My Language Choice and Motivation Behind it

### My choice: **Golang**

### Why?

My interest lies on backend and infrastructure side as well as lower layer. Since C and C++ were not allowed to select. I decided to learn golang because I wanted to learn and build application using micro-service as well as get experience in using pointer (lower-level stuff).

## Getting Started

I created docker-compose.yml to setup environment. If you have Docker Desktop installed all you need to do is start container to start server and run client application.

### 1. Clone project to your project folder

Create project folder and run following command:

```bash
git clone https://github.com/FumiKimura/ccp2-project-polygottal.git
```

### 2. Run Docker-Compose up to start server

In root directory of the project folder:

```bash
docker-compose up
```

If you see, "Started Listening to server...," server is up and ready.

### 3. Run client side application

From root project folder, cd into client directory and run the main.go

```bash
cd client
go run client.go
```

Then, enter your username and start talking. You need to say something to start receiving message from other clients.

### 4. Run multiple client application to start chat

Open another terminal and run client application. Yay there should be messages between clients routed by the server.

### 5. Exiting client application

In the client side app, input following !exit to quit client application. I have not implemented my code to hand Ctrl + C (Termination Signal) yet. This is something I will keep on working. If you have any suggestions, please kindly reach out.
