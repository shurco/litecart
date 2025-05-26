# Commit Message Instructions

- Language: Use english
- Use conventional commit message format.
- The commit message should have a short description (50 characters or less) followed by a blank line and then a longer description.
- The short description should be in the format: `<icon> <type>(<scope>): <short description>`
  - `type`: The type of change (e.g., feat, fix, docs, style, refactor, test, chore).
    - `feat`: ✨ A new feature
    - `fix`: 🐛 A bug fix
    - `docs`: 📝 Documentation only changes
    - `style`: 💄 Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
    - `refactor`: ♻️ A code change that neither fixes a bug nor adds a feature
    - `test`: ✅ Adding missing tests or correcting existing tests
    - `chore`: 🔧 Changes to the build process or auxiliary tools and libraries such as documentation generation
    - `perf`: ⚡️ A code change that improves performance
    - `ci`: 👷 Changes to CI configuration files and scripts
    - `build`: 🏗️ Changes that affect the build system or external dependencies
    - `revert`: ⏪ Reverts a previous commit
    - `wip`: 🚧 Work in progress
    - `security`: 🔒 Security-related changes
    - `i18n`: 🌐 Internationalization and localization
    - `a11y`: ♿ Accessibility improvements
    - `ux`: 🎨 User experience improvements
    - `ui`: 🖌️ User interface changes
    - `config`: 🔧 Configuration file changes
    - `deps`: 📦 Dependency updates
    - `infra`: 🌐 Infrastructure changes
    - `init`: 🎉 Initial commit
    - `analytics`: 📈 Analytics or tracking code
    - `seo`: 🔍 SEO improvements
    - `legal`: ⚖️ Licensing or legal changes
    - `typo`: ✏️ Typo fixes
    - `comment`: 💬 Adding or updating comments in the code
    - `example`: 💡 Adding or updating examples
    - `mock`: 🤖 Adding or updating mocks
    - `hotfix`: 🚑 Critical hotfix
    - `merge`: 🔀 Merging branches
    - `cleanup`: 🧹 Code cleanup
    - `deprecate`: 🗑️ Deprecating code or features
    - `move`: 🚚 Moving or renaming files
    - `rename`: ✏️ Renaming files or variables
    - `split`: ✂️ Splitting files or functions
    - `combine`: 🧬 Combining files or functions
    - `add`: ➕ Adding files or features
    - `remove`: ➖ Removing files or features
    - `update`: ⬆️ Updating files or features
    - `downgrade`: ⬇️ Downgrading files or features
    - `patch`: 🩹 Applying patches
    - `optimize`: 🛠️ Optimizing code
    - `docs`: 📝 Documentation changes
    - `test`: ✅ Adding or updating tests
    - `fix`: 🐛 Bug fixes
    - `feat`: ✨ New features
    - `style`: 💄 Code style changes (formatting, etc.)
    - `refactor`: ♻️ Code refactoring
    - `perf`: ⚡️ Performance improvements
    - `ci`: 👷 Continuous integration changes
    - `build`: 🏗️ Build system changes
    - `revert`: ⏪ Reverting changes
    - `wip`: 🚧 Work in progress
    - `security`: 🔒 Security improvements
    - `i18n`: 🌐 Internationalization changes
    - `a11y`: ♿ Accessibility improvements
    - `ux`: 🎨 User experience improvements
    - `ui`: 🖌️ User interface changes
    - `config`: 🔧 Configuration changes
    - `deps`: 📦 Dependency updates
    - `infra`: 🌐 Infrastructure changes
    - `init`: 🎉 Initial commit
    - `analytics`: 📈 Analytics changes
    - `seo`: 🔍 SEO improvements
    - `legal`: ⚖️ Legal changes
    - `typo`: ✏️ Typo fixes
    - `comment`: 💬 Comment changes
    - `example`: 💡 Example changes
    - `mock`: 🤖 Mock changes
    - `hotfix`: 🚑 Hotfix changes
    - `merge`: 🔀 Merge changes
    - `cleanup`: 🧹 Cleanup changes
    - `deprecate`: 🗑️ Deprecation changes
    - `move`: 🚚 Move changes
    - `rename`: ✏️ Rename changes
    - `split`: ✂️ Split changes
    - `combine`: 🧬 Combine changes
    - `add`: ➕ Add changes
    - `remove`: ➖ Remove changes
    - `update`: ⬆️ Update changes
    - `downgrade`: ⬇️ Downgrade changes
    - `patch`: 🩹 Patch changes
    - `optimize`: 🛠️ Optimize changes
  - `scope`: The scope of the change (e.g., component or file name). Include this if the change is specific to a particular part of the codebase.
- `short description`: A brief summary of the change.
- The long description should provide additional context and details about the change.
  - Explain why the change was made.
  - Describe what is being used and why.
  - Include any relevant information that might be useful for understanding the change in the future.
  - Reference any related issues or pull requests at the end of the long description.
- If the commit fixes an issue or task, include `Fixes #<issue-number>` or `Closes #<issue-number>` at the end of the long description.
- If the commit introduces a breaking change, include `BREAKING CHANGE: <description of the breaking change>` at the end of the long description.

## Example

### Commit Message Example

```
✨ feat(auth): Add user authentication

Added user authentication using JWT. This includes login, registration, and token verification endpoints.

- Implemented JWT-based authentication.
- Added login and registration endpoints.
- Added middleware for token verification.

Fixes #123
```

### Breaking Change Example

```
♻️ refactor(api): Update API endpoints

Refactored the API endpoints to follow RESTful conventions. This change affects all existing API calls.

- Updated endpoint URLs to follow RESTful conventions.
- Modified request and response formats.

BREAKING CHANGE: All existing API calls need to be updated to the new endpoint URLs.
```

## Example
