# Actions Run Times

This tool will return the following:
- all GitHub workflows in a given repository
- number of successful runs per workflow
- average build time of each workflow

Example:
```bash
$ actions-run-times list workflows --owner "jgrigorian" --repo "certscan"

   Repository              Workflow                         ID            Successful Runs     Average Run Times

   jgrigorian/certscan     Build Release Binaries           145003030     1                   5m27s
   jgrigorian/certscan     Enforce Conventional Commits     145003031     2                   15s
   jgrigorian/certscan     release-please                   145002944     5                   2m22s              
```