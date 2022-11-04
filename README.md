# Actions Build Time

Requirements:
- environment variable which has your Github Token (GH_TOKEN)

This tool will return the following:
- number of successful runs
- average build time

Example:
```bash
❯ actions-build-time --org "society6" --repo "s6-web" --workflowId "38035010" --branch "pipeline-actions"
Repository                      Branch                          Successful Runs         Average Build Time
society6/s6-web                 pipeline-actions                4                       9m19s

```