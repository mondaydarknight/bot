# bot
A simple bot service to interact with LINE bot channel.

## Getting started
Set up a local environemnt with Docker.
```console
# Copy the environment example to the dot file.
$ cp env.example .env
# Start bot service of Docker containers via Makefile command
$ make up
# Stop bot service of Docker containers
$ make down
```

## Implementation
Design a simple workflow that applies golang web server and MongoDB database to meet the requirement, implement a [layered architectutre](https://martinfowler.com/bliki/PresentationDomainDataLayering.html) to isolate the domain model of business logic and eliminate the dependency on infrastructure, user interface or application logic. Define the user interface layer (controller) that is responsible for presenting information to the user and interpreting user commands, and design an infrastructure layer that access external service like database and LINE bot service.

Table schema: `messages`
| Column       | Type         | Description                             |
| ------------ | ------------ | --------------------------------------- |
| userId       | String       | The unique key identifier of user ID.   |
| text         | String       | The push message sent from LINE bot.    |
| createdAt    | Date         | The created date and time.              |

## API reference docs
List API endpoints for bot service.

### GET `/api/v1/linebot/messages`
Retrieve a list of push messages of the user.

#### Query Parameters
| Field          | Type   | Required | Description |
| -------------- | ------ | -------- | ----------- |
| userId         | String |          | The field to filter the rows by the user ID. |

#### Responses
`200`
| Field       | Type    | Required | Description |
| ----------- | ------- | -------- | ----------- |
| messages    | Message | Y        | A list of push messages. |

`Message`
| Field       | Type    | Required | Description                           |
| ----------- | ------- | -------- | ------------------------------------- |
| userId      | String  | Y        | The unique key identifier of user ID. |
| message     | String  | Y        | The push message sent from LINE bot.  |

`404`, `500`
| Field   | Type   | Required | Description        |
| ------- | ------ | -------- | ------------------ |
| code    | Int    | Y        | The error code.    |
| message | String | Y        | The error message. |

```sh
GET /api/v1/linebot/messages

HTTP/1.1 200 OK
{
  "messages": [
    {
      "userId": "U13c4f2c3120a19260ba1c44ee23dae58",
      "text": "helloworld",
      "createdAt": "2023-01-09T03:08:16.965Z"
    },
    {
      "userId": "U13c4f2c3120a19260ba1c44ee23dae58",
      "text": "foo bar baz test",
      "createdAt": "2023-01-09T03:08:18.869Z"
    }
  ]
}
```

### POST `/api/v1/linebot/messages`
Send a push message to LINE bot channel.

#### Body Parameters
| Field    | Type   | Required | Description                           |
| -------- | ------ | -------- | ------------------------------------- |
| userId   | String | Y        | The unique key identifier of user ID. |
| text     | String | Y        | The push message sent to LINE bot.    |

#### Responses
`200` (Empty response)

`400`, `500`
| Field   | Type   | Required | Description        |
| ------- | ------ | -------- | ------------------ |
| code    | Int    | Y        | The error code.    |
| message | String | Y        | The error message. |

```sh
POST /api/v1/linebot/messages

HTTP/1.1 200 OK
{}

HTTP/1.1 400 Bad Request
{
  "code": 400,
  "message": "Key: 'Message.Text' Error:Field validation for 'Text' failed on the 'required' tag"
}
```

### POST `/api/v1/linebot/webhook`
Store the push message via webhook.

#### Body Parameters ([learn more about LINE webhook API doc](https://developers.line.biz/en/reference/messaging-api/#request-body))
| Field       | Type   | Required | Description                           |
| ----------- | ------ | -------- | ------------------------------------- |
| destination | String | Y        | User ID of a bot that should receive webhook events. The user ID value is a string that matches the regular expression, U[0-9a-f]{32}. |
| events      | String | Y        | Array of webhook event objects. The LINE Platform may send an empty array that doesn't include a webhook event object to confirm communication.    |

#### Responses
`200` (Empty response)

`500`
| Field   | Type   | Required | Description        |
| ------- | ------ | -------- | ------------------ |
| code    | Int    | Y        | The error code.    |
| message | String | Y        | The error message. |

```sh
POST /api/v1/linebot/webhook

HTTP/1.1 200 OK
{}
```
