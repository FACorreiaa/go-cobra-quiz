# Go Quiz REST API with Cobra CLI

- Boostrap Cobra with cobra-cli
- Use air for live reload
- lint configuration inside ./config
- Chi for group routing
- Ideally I would have everything inside two separate folders: cmd for the UI and internal for the API.
- With this structure the server is running inside the internal logic and CLI UI inside cmd
- The service file contains a single service instance, which will cause problems if we want to change one of the repos. Ideally we would change this per feature, but for this use case I have put everything inside one handler, one service and one repository folder.
- I pass ctx without needing it (for now). In case we plug a DB, to have an history of results, for example, the ctx is there to be used. But because its not used, lint will not pass successfully.

# Run the API
- make up

# Run CLI
docker run go-cobra-quiz-cli "your command here"
or
make cli-run COMMAND="your command here"


## list of commands:
- start
- setuser
- submit
- list
- ranking


TODO:
- [ ] Fix CLI setname command
- [ ] Prepare configs for env variables or deployments keys
- [ ] Add more test coverage
- [ ] Decouple NewService inside api/service implementation for each feature
- [ ] Make github action linter pass.
