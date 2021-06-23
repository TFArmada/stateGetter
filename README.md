# stateGetter
***stateGetter*** is a simple Go command line tool that downloads ***latest*** Terraform Cloud remote state files for given workspace.
 
It's intended use is in GitHub Actions to dynamically prepare Terraform provider configuration based
 on remote terraform state outputs.
 
## Use
```shell script
Usage of ./stateGetter:
  -filename string
        Output file name (default "stateGetter.tfstate")
  -organization string
        TFE Organization (required)
  -workspace string
        TFE Workspace (required)
```
### Example
```shell script
export TFE_TOKEN=<secret tfe token>
./stateGetter -organization ganekov -workspace armada-accounts-prime
```

## Github Action Example 
*TODO*