

Based on the search results, here's a simple workflow using the Azure DevOps CLI:

1. **Create a new branch**:
```
az repos pr create --project=myproject --repository=myrepo --source-branch=main --target-branch=newbranch --title="New Feature" --description="Implementing new feature XYZ"
```
This creates a new branch `newbranch` from `main`.

2. **Add code changes to the new branch**:
Make code changes and commit them to the `newbranch`.

3. **Create a pull request**:
```
az repos pr create --project=myproject --repository=myrepo --source-branch=newbranch --target-branch=main --title="New Feature" --description="Implementing new feature XYZ"
```
This creates a new pull request to merge `newbranch` into `main`.

4. **Approve the pull request**:
```
az repos pr approval add --id=123 --project=myproject --repository=myrepo --user=reviewer@email.com
az repos pr approval add --id=123 --project=myproject --repository=myrepo --user=lead@email.com
```
This approves the pull request with ID 123 from two reviewers.

5. **Complete the pull request**:
```
az repos pr complete --id=123 --project=myproject --repository=myrepo --merge-commit-message="Merged PR 123: New Feature"
```
This merges the pull request into the target branch (`main`).

You can automate this entire workflow using a script or Azure Pipelines. Note that you'll need to replace `myproject`, `myrepo`, branch names, and email addresses with your actual values.