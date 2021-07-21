# About the Project

The backend for https://github.com/SocietyClub/RCV

# Getting Started

Make sure to install Golang beforehand: https://golang.org/dl/

### Generate key for service account

1. Log into Google Cloud and head to Service accounts under the IAM & Admin page. https://console.cloud.google.com/iam-admin/serviceaccounts 

![Service account img](img/RCVSA.png)

2. Select rcv-service-account and under action click *Manage Keys*
3. Create a new key, select json file format and save it for later. Do not save the key in this repo in case we accidently deploy it :) 
3. Provide authentication credentials to your application code by setting the environment variable GOOGLE_APPLICATION_CREDENTIALS

![Keys](img/envkey.png)

For more details check out: https://cloud.google.com/firestore/docs/quickstart-servers#go

*Working on script for this soon...*

4. `cd RCV-backend/helloworld`
5. `go build`
6. `.\helloworld.exe` if you are using windows
7. Once your done, `go clean` to remove object files and cached files

# Generating Go from an OpenAPI Spec
See [how-to-use-open-api-generator.md](how-to-use-open-api-generator.md) for details

# Deployment
Only package maintainers have access to deploy. These instructions are for them.

Once you are authorized to deploy through gcloud via your google account, you may download gcloud locally (https://cloud.google.com/sdk/docs/quickstart?authuser=5) or use the terminal from the online google console.

1. Go to https://console.cloud.google.com/home/dashboard?project=societyclub-rcv-backend and open the terminal (or download gcloud locally)

![gcp terminal button](img/README-gcp-terminal-button.png)

2. `gcloud config set project societyclub-rcv-backend`
    You can check if the config was set properly via `gcloud config list`
3. `git clone https://github.com/SocietyClub/RCV-backend.git`
4. `cd RCV-backend/helloworld`
5. `git pull`
6. `gcloud app deploy`


