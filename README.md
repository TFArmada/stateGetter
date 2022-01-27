# stateGetter
***stateGetter*** is a simple Go command line tool and GitHub action that downloads ***latest*** Terraform Cloud remote state files for given workspace.
 
## GitHub Action
```yaml
    steps:
      -
        name: stateGetter get
        uses: TFArmada/stateGetter@master
        with:
          filename: state.tfstate #Filename to save the state as
          organization: sumup #Terraform Cloud Organization 
          workspace: autosg-classic-live #A workspace in the Terraform Cloud Organization
        env:
          TFE_TOKEN: ${{ secrets.TFE_TOKEN }} #Terraform Cloud Token
```
 
## Use as standalone binary
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

## Using the binary in GitHub Action
```yaml
name: Generate
on: [push]
jobs:
  prepare:
    name: Prepare
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v2
      - name: Download stateGetter
        shell: bash 
        run: |
          # Fetch stateGetter binary
          sudo wget https://github.com/TFArmada/stateGetter/releases/download/v0.1/stateGetter -O /usr/local/bin/stateGetter
          # Validate checksum against expected value
          SUM=a9fc23d9fbe20797a16a963176faa1df64d1cbcadbe85f42a46aa41df4f1932c
          if [[ "$SUM" != "$(sha256sum /usr/local/bin/stateGetter | awk '{ print $1 }')" ]]
          then
            echo "stateGetter binary checksum does not match expected"
            exit 1;
          fi
          sudo chmod +x /usr/local/bin/stateGetter
      - name: Execute stateGetter
        run: /usr/local/bin/stateGetter -organization ganekov -workspace armada-accounts-prime -filename accounts.tfstate
        shell: bash
        env:
          TFE_TOKEN: ${{ secrets.tfe_token }}
      - name: Prepare Terraform
        uses: hashicorp/setup-terraform@v1
        with:
          terraform_version: 1.0.0
      - name: Show Outputs
        run: terraform state list -state=accounts.tfstate
        shell: bash
```