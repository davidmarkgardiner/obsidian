
GO program and ado pipline that will automate the process of extracting specific configuration data from Kubernetes kustomization files and preparing this data to be used in a pipeline, this is then sent in payload  to re onboard RBAC role binding after AKS  cluster rebuild. 


```json
{
    "Topic": "Kubernetes kustomization configuration extraction",
    "Story": "As a DevOps engineer, I want to automate the extraction of specific configuration data from Kubernetes kustomization files so that I can use it in my pipeline for further processing or deployment tasks, ensuring the management of configurations for multiple environments or services is efficient and error-free.",
    "Criteria": "Given that I am a DevOps engineer, when I execute the program x, then it should recursively search through the specified directory structure for Kubernetes kustomization files and extract the required configuration data. Given that multiple kustomization files are found, when the program runs, then it should process each file and prepare the extracted data for use in a pipeline. Upon successful execution, then I should receive confirmation that the configurations are ready for the next steps."
}
```