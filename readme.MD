## Introduction
This project provide an application that receives a number from user and returns the highest prime number lower than the input number. The application relies on the client-server architecture, in which user enter a number on user's device, the input number is propagated to the server, after some validations, the server then tries to figure out what is the result and return it to the user.

## References
1. The demo of the code: https://phuonghau.com
2. The git repository: https://github.com/phuonghau98/stably-togo

## How to run the app
1. Start the web server

```bash
go get
go run main.go
```
1. Start the web client

```bash
cd ui
npm install
npm run start
```

## Decisions taken and tradeoffs
1. Don't use server-side rendering to simplify the project architecture.
   1. Working with such simple requirement like an input and a button only need a html template which we can just return it to the user after it is rendered on our side, and then we can inject the result variable to the returning template to the user afterwards. But I choose to use a SPA app instead which call the rest server's API using REST. Because the SPA app enables us to have an ability to provide the user more fine-grained UI, it also requires less work to have something like validations and animations on the user device.
   2. The obvious tradeoff is that we have a more complex architecture.

2. Don't use dedicated web server to serve client app (react-app)
   1. Let the Golang server to serve the static content like html, css, js... built from the client source code instead of using a dedicated web server like nginx, apache,...That will laverage the ease when deploying the code because we have to 2 sides to deploy, first thing is the react app, the second thing is Golang server. 
   2. The tradeoff of this is that we don't have capability to scale the app vertically because we haven't implemented any kind of caching features which are already packed with dedicated server. So basically, we lack of capabilities to have operating enhancement activities like monitoring requests, caching, request intercepting.

3. I decided to use the `BigInt.ProbablyPrime` function to determine a number is a prime number.
   1. After a few tries and spending time research how to have an efficient algorithm to determine whether a number is prime or not, I realized that there is a lot of it which beats my own implemented one. Some are only suitable for specific ranges, some relies on so many theories, so I decided not to re-invent the wheel but to choose the well-tested, built-in one of Go's `big` library: `BigInt.ProbablyPrime`, it works flawlessly.
   2. The obvious tradeoff is that maybe I'm not going to be hired by this decision, anyway, I just want my application to work well in the end.

## Learning Journey
1. Golang's bigint
I'm not so quite familiar with Golang's ``big`` library, I spent a huge amount of time learning how to take advantage of it to unlock the capability to handle big numbers.