# Actions Build Time

Requirements:
- environment variable which has your Github Token (GH_TOKEN)

This tool will return the following:
- number of successful runs
- average build time

Example:
```bash
$ actions-build-time --org "jgrigorian" --repo "certscan" --workflowId "38035010" --branch "master"
Repository                      Branch                          Successful Runs         Average Build Time
jgrigorian/certscan             master                          4                       9m19s
```