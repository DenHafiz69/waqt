# Waqt - A bootdotdev capstone project

## My goal

This project is not really about the prayer-time-telling-service, there are already a bajilion of those. I write this app mostly to learn about DevOps. Having my own app that I can improve regularly so I can practice CI/CD, deployments, logging, bug-fixing, and much more.

## Game plan

- Backend in Go
	- [x] Create a backend server that GET data from waktu solat API
	- [] Before every prayer time, ping a Telegram bot, reminding me about the prayer time approaching
- Deployment
	- [x] Dockerized everything so it is easily deployable, either locally, or on the cloud
	- [] Setup CICD so the app can be updated regularly
		- [x] Setup CI to check the code quality, linting, testing, etc
		- [] Setup CD to automatically push the app to AWS upon merging
	- [] Terraform config. Run docker on EC2 at first.
- Documentation
	- [] At the end of this project (hopefully by end of October), revisit this readme and update it with "what I learned", "what could be done differently (and why I still do it my way)", "future features"

