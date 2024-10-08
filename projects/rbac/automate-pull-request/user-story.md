To create an automated pull request (PR) in Azure DevOps (ADO), here’s a concise user story detailing the steps involved:

---

### **User Story: Automated Pull Request in Azure DevOps**

**Title**: Automate the creation of pull requests in Azure DevOps  
**As a** DevOps Engineer,  
**I want** to automate the creation of pull requests in Azure DevOps,  
**so that** code changes can be reviewed and merged more efficiently without manual intervention.

#### **Acceptance Criteria**:

1. **Trigger Automation**: 
   - The automation is triggered when code is pushed to a specific branch (e.g., `feature/*` or `bugfix/*`).
   
2. **Create Pull Request**:
   - Upon detecting the branch push, the automation should create a pull request targeting the `main` or `develop` branch.
   - The pull request title should follow a predefined format, including the branch name and ticket reference (if applicable).

3. **Assign Reviewers**:
   - The automation will assign specific team members as reviewers based on predefined rules (e.g., based on the type of changes or files modified).

4. **Set Auto-Complete**:
   - Enable the "Auto-Complete" option in the PR, which merges the pull request automatically after all policies and checks are met.

5. **Add Labels**:
   - Labels (e.g., `feature`, `bugfix`, `urgent`) should be applied automatically based on branch naming conventions.

6. **Notify Team**:
   - Send an automated notification to a specific channel (e.g., Microsoft Teams or Slack) when the PR is created.

#### **Technical Details**:
- **Azure DevOps Pipeline**:
   - Leverage Azure Pipelines or ADO REST API to implement the automation.
   - Ensure YAML pipeline includes a task for creating a PR (e.g., using `az repos pr create` command or similar).
   
- **Branch Filtering**:
   - Set up branch filters to only trigger the PR creation for specific branches (`feature/*`, `bugfix/*`).

- **Permissions**:
   - Ensure the service principal or automation user has appropriate permissions to create pull requests and assign reviewers.

---

### **Tasks**:
1. Set up the Azure Pipeline to detect code changes in specific branches.
2. Implement the script or use ADO API to automatically create a pull request.
3. Define reviewer assignment rules based on project needs.
4. Enable Auto-Complete for the pull request in the pipeline.
5. Add automated labeling based on branch name.
6. Configure notifications for the team once the PR is created.

### **Dependencies**:
- Access to Azure DevOps repository and required permissions for PR creation.
- Defined rules for reviewer assignments and labeling.

---

This story ensures the automation of PR creation in ADO is handled seamlessly, reducing manual effort and streamlining the code review process. Let me know if you need further clarification!