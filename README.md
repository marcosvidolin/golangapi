# QuestionsAndAnswers.com

## Statement

You are to design the backend side of a system for the following business idea.
We want to build a site called QuestionsAndAnswers.com that will compete with Quora/Stackoverflow and others with 1 major difference. We will only allow 1 answer per question. If someone thinks they have a better answer, they will have to update the existing answer for that question instead of adding another answer. In essence, each question can only have 0 or 1 answer.

The backend should support the following operations:

- Get one question by its ID
- Get a list of all questions
- Get all the questions created by a given user
- Create a new question
- Update an existing question (the statement and/or the answer)
- Delete an existing question

No user tracking or security needed for this version.
Database design is up to you.

We would like to receive code that runs, so remember to focus on the MVP functionality. You can document what's missing that you wish you had more time for? Please think about the different problems you might encounter if the business idea is successful. This would include considerations such as increased load, increased data, and an upvoting feature.

## How to Run

Building

```shell
docker build  -t questionapi .
```

Running

```shell
docker run -p 8080:8080 questionapi
```

## Usage

Examples:

- Get one question by its ID

```shell
curl -X GET http://localhost:8080/questions/613f46933cd7eb9c67e9a6c3
```

- Get a list of all questions

```shell
curl -X GET http://localhost:8080/questions
```

- Get all the questions created by a given user

```shell
curl -X GET http://localhost:8080/questions?author=marcosvidolin
```

- Create a new question

```shell
curl -X POST http://localhost:8080/questions
--data '{"body": "How to create a question?"}'
```

- Update an existing question (the Statement)

```shell
curl -X PUT http://localhost:8080/questions/613f46933cd7eb9c67e9a6c3/answers/613f46933cd7eb9c67e9a6c3 \
--data '{"body": "Change the questions statement... ?"}'
```

- How to create an answer

```shell
curl -X PUT http://localhost:8080/questions \
--data {"body": "How to create an answer?"}
```

- Update an existing answer

```shell
curl -X PUT http://localhost:8080/questions/613f46933cd7eb9c67e9a6c3 \
--data '{"body": "Change the answer... !"}'
```

- Delete an existing question

```shell
curl -X DELETE http://localhost:8080/questions/613f46933cd7eb9c67e9a6c3

```
