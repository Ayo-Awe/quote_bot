# Quote Bot

This repository features a telegram quote bot written in golang. The bot sends an inspirational quote upon request from the user.

## Commands

- `/start` : Initialize a conversation with the bot
- `/quote` : Get a new quote

## Dependencies

The quotes used by the bot are sourced from [Quotable]("https://github.com/lukePeavey/quotable"), a free and open source quotations API

## Todo

- [ ] Abstract quote client to make switching to a new quotes API simple
- [ ] Implement a get quote by category feature

## Deployment

`Quote bot` currently runs as a web service on [Render]("https://render.com"). It handles incoming requests via webhooks rather than long polling.