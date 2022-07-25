build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o terraform_infra/target/golambdabin -ldflags '-w' main.go && zip -j terraform_infra/target/golambda.zip target/golambdabin

check:
	terraform -chdir=terraform_infra fmt -check

init:
	terraform -chdir=terraform_infra init

validate:
	terraform -chdir=terraform_infra validate -no-color

plan:
	terraform -chdir=terraform_infra plan -no-color -input=false

apply:
	terraform -chdir=terraform_infra apply --auto-approve

destroy:
	terraform -chdir=terraform_infra destroy --auto-approve