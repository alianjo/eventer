# Kubernetes Telegram Event Operator

## Overview

This project is a Kubernetes operator responsible for sending Kubernetes events from a cluster to a Telegram channel.

## Table of Contents

- [Kubernetes Telegram Event Operator](#kubernetes-telegram-event-operator)
  - [Overview](#overview)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Configuration](#configuration)
  - [Installation](#installation)

## Prerequisites

* Kubernetes cluster (1.23 or later)
* Telegram Bot API token
* Telegram channel ID

## Configuration

To configure the operator, update the configuration of the `Deployment` resource with the following fields:

* `TELEGRAM_BOT_TOKEN`: Telegram Bot API token
* `TELEGRAM_CHANNEL_ID`: Telegram channel ID

Example `Deployment` resource in `manifests/deployment.yaml`:

```yml
        env:
        - name: TELEGRAM_CHANNEL_ID
          value: "chat_id"
        - name: TELEGRAM_BOT_TOKEN
          value: "your_bot_token"
```

## Installation

To install the operator, run the following command:

```bash
make deploy
```