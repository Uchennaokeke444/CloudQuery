In Order for CloudQuery to sync resources from your HubSpot setup, you will need to authenticate with your HubSpot account. You will need to create a [HubSpot Private App](https://developers.hubspot.com/docs/api/private-apps), and copy the App Token to the spec.
If not specified `HUBSPOT_APP_TOKEN` environment variable will be used instead.

```bash copy
export HUBSPOT_APP_TOKEN=<your_app_token> # optional, if not using spec configuration
```
