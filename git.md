To push an existing local directory to GitHub, you'll need to follow these steps:

1. Create a new repository on GitHub:
   - Go to GitHub.com and log in
   - Click the "+" icon in the top right and select "New repository"
   - Name your repository (e.g., "obsidian")
   - Don't initialize with a README, .gitignore, or license

2. Open a terminal and navigate to your local directory:
   ```
   cd Documents/obsidian/
   ```

3. Initialize the local directory as a Git repository:
   ```
   git init
   ```

4. Add the files in your new local repository:
   ```
   git add .
   ```

5. Commit the files:
   ```
   git commit -m "Initial commit"
   ```

6. Add the URL for the remote repository on GitHub:
   ```
   git remote add origin https://github.com/YOUR-USERNAME/obsidian.git
   git remote add origin https://github.com/davidmarkgardiner/obsidian.git
   ```
   Replace YOUR-USERNAME with your actual GitHub username.

7. Push the changes to GitHub:
   ```
   git push -u origin main
   ```
   Note: If your default branch is named "master" instead of "main", use "master" in the command above.

If you encounter any issues with the last step, it might be because your local branch name doesn't match the remote. In that case, you can try:

```
git push -u origin HEAD:main
```

This will push your current HEAD to the main branch on GitHub.

Remember, if your repository contains sensitive information, make sure to add a `.gitignore` file to exclude any private files before pushing to a public repository.