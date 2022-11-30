#!/bin/sh
set -euo pipefail

TAG=$1
if [ -z "$TAG" ]; then
    echo "Usage: cd.sh <tag>"
    echo "Example: cd.sh 684b4e6"
    exit 1
fi

if ! command -v gh > /dev/null; then
    echo "Please install gh"
    exit 1
fi

{
    cd terraform

    # Update variables
    echo "tag = \"$TAG\"" > terraform.tfvars

    # Create PR
    git checkout -b "deploy/$TAG"
    git add terraform.tfvars
    git commit -m "Deploy $TAG"
    git push
    gh pr create --title "Auto-Deploy: $TAG" --body "Deploys $TAG"
    git checkout master
}