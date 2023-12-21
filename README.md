# Slackbot Golang Project

A simple Golang Slackbot using the [Slacker](https://github.com/shomali11/slacker) library.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Welcome to the Slackbot Golang project! This bot is designed to respond to a "ping" command with a witty reply. It utilizes the Slacker library for easy interaction with the Slack API.

## Features

- Responds to the "ping" command with a custom message.
- Prints command events for analytics purposes.

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/slackbot-golang.git
```
1 Change into the project directory:
```
cd slackbot-golang
```
2 Set env token of app and bot
```
export SLACK_BOT_TOKEN=xoxb-YourBotToken
export SLACK_APP_TOKEN=xapp-YourAppToken
```
3 Build and run the project:
```
go build
./slackbot-golang
```
![image](https://github.com/AnkurKumarShukla/Golang-slack-bot/assets/80956033/62678a41-3672-4082-aed7-f58830c8c331)

## Usage
Once the bot is up and running, it will listen for the "ping" command in your Slack workspace. You can also check the command events for analytics by monitoring the console.
![image](https://github.com/AnkurKumarShukla/Golang-slack-bot/assets/80956033/f9ddcd0d-c779-44ff-8050-728285c6fbbc)

## Available Command
  -/ping: Receive a witty reply from the bot.

## Contributing

Contributions are welcome! Please follow these steps:

1. **Fork the repository.**

2. **Create a new branch:**

    ```bash
    git checkout -b feature/my-feature
    ```

3. **Make your changes and commit them:**

    ```bash
    git commit -am 'Add my feature'
    ```

4. **Push to the branch:**

    ```bash
    git push origin feature/my-feature
    ```

5. **Open a pull request.**

## License

This project is licensed under the [MIT License](LICENSE).






