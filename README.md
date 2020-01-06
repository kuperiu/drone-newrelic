# drone-mewrelic
## Description
Post deployments from drone CI to new relic by name


## Example
```
- name: Newrelic - Post Deployment
  image: kuperiu/drone-newrelic
  environment:
    - APPLICATION_NAME="my app"
  secrets:
    - source: "MY_KEY"
      target: "NEW_RELIC_LICENSE_KEY"
```
## Build
```make build```

