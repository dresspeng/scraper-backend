# scraper

Online scraper for building a dataset for ML.


## installation

Install Golang and MongoDB

If pbm with package `<package>: command not found`:

    export GOPATH="$HOME/go"
    PATH="$GOPATH/bin:$PATH"

## Localstack

https://docs.localstack.cloud/get-started/#localstack-cli
https://github.com/localstack/localstack

or with docker:

    docker run --rm -it -p 4566:4566 localstack/localstack
## Docker

    sudo docker build -t scraper-img .
    sudo docker run -it -p 8080:8080 -p 27017:27017 --rm --name scraper-run --env-file <state>.env scraper-img

## Run

### Run without docker

    ENV=local go run src/main.go

### Build without docker

    go build -o scraper src/main.go
    ENV=local ./scraper

## License

must share photos generated with https://creativecommons.org/licenses/by-sa/2.0/

## Env

Create a local.env file:

    MONGODB_URI=mongodb://localhost:27017
    SCRAPER_DB=scraper
    TAGS_UNDESIRED_COLLECTION=tagsUndesired
    TAGS_DESIRED_COLLECTION=tagsDesired
    PRODUCTION=imagesProduction
    PENDING=imagesPending
    UNDESIRED=imagesUndesired
    VALIDATION=imagesValidation
    USERS_UNDESIRED_COLLECTION=usersUndesired
    IMAGES_BUCKET=<s3_bucket_name>
    FLICKR_PRIVATE_KEY=***
    FLICKR_PUBLIC_KEY=***
    UNSPLASH_PRIVATE_KEY=***
    UNSPLASH_PUBLIC_KEY=***
    PEXELS_PUBLIC_KEY=***
    ENV=local

ENV is either `production`, `staging`, `development` or `local`

## linter

https://github.com/mgechev/revive

    revive -config revive.toml

## Dependencies

    go mod tidy

## AWS Bash scripts

    sudo chmod +x <script-file>
    ./<script-file>

