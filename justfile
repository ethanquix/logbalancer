help:
    just --list

clean:
    rm -rf gen
gen: clean
    buf generate

example_deploy:  # example to deploy on gcloud
    export KO_DOCKER_REPO := "us-east1-docker.pkg.dev/gcp_username/example/"
    ko build -P ./cmd/cloudrun
    gcloud run deploy example-logbalancer \
    --image=us-east1-docker.pkg.dev/nofo-prod/nofo-images/github.com/ethanquix/logbalancer/cmd/cloudrun:latest \
    --allow-unauthenticated \
    --service-account=XXX@developer.gserviceaccount.com \
    --timeout=1000 \
    --memory=1Gi \
    --max-instances=1 \
    --set-env-vars=PROD=true \
    --set-secrets=internal_services_token=internal_services_token:latest \
    --set-secrets=slack_token=slack_token:latest \
    --set-env-vars=slack_token=abcdefgh \
    --ingress=all \
    --region=us-east1 \
    --project=example