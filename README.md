# Actions Build Time

Requirements:
- environment variable which has your Github Token (GH_TOKEN)

This tool will return the following:
- all GitHub workflows in a given repository
- number of successful runs per workflow
- average build time of each workflow

Example:
```bash
$ actions-run-times list workflows --owner "derailed" --repo "k9s"

   Repository       Workflow                             ID            Successful Runs     Average Build Times   
                                                                                                                 
   derailed/k9s     K9s Lint                             76629439      227                 29m1s                 
   derailed/k9s     K9s Test                             6893216       543                 12m6s                 
   derailed/k9s     docker in /. - Update \#978135623    133662603     79                  27s                 
```