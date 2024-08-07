検証が終わり次第削除します

## playwright がまとめて更新できている

```json
{
  "$schema": "https://docs.renovatebot.com/renovate-schema.json",
  "extends": ["config:recommended", "group:allNonMajor"],
  "prHourlyLimit": 0,
  "prConcurrentLimit": 0,
  "packageRules": [
    {
      "managers": ["npm"],
      "commitMessage": "npm - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "{{parentDir}}-",
      "packagePatterns": [".*"],
      "ignoreDeps": ["playwright"]
    },
    {
      "matchPackageNames": ["node"],
      "commitMessage": "node - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "{{parentDir}}-"
    },
    {
      "matchPackageNames": ["pnpm"],
      "commitMessage": "pnpm - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "{{parentDir}}-"
    },
    {
      "managers": ["npm", "dockerfile", "devcontainer"],
      "matchPackageNames": ["@playwright/test", "mcr.microsoft.com/playwright"],
      "commitMessage": "playwright - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "playwright-"
    },
    {
      "managers": ["gomod"],
      "commitMessage": "gomod - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "{{parentDir}}-",
      "packagePatterns": [".*"]
    },
    {
      "matchPackageNames": ["golang"],
      "commitMessage": "golang - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "golang-"
    },
    {
      "matchPackageNames": ["mysql"],
      "commitMessage": "mysql - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "mysql-"
    }
  ]
}
```

## 共通化、playwright 以外よさそう？

```json
{
  "$schema": "https://docs.renovatebot.com/renovate-schema.json",
  "extends": ["config:recommended"],
  "prHourlyLimit": 0,
  "prConcurrentLimit": 0,
  "ignoreDeps": ["playwright"],
  "packageRules": [
    {
      "managers": ["npm"],
      "extends": ["group:allNonMajor"],
      "commitMessage": "npm - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "{{parentDir}}-",
      "packagePatterns": [".*"]
    },
    {
      "managers": ["npm", "dockerfile", "devcontainer"],
      "matchPackageNames": ["@playwright/test", "mcr.microsoft.com/playwright"],
      "commitMessage": "playwright - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "playwright-"
    },
    {
      "managers": ["gomod"],
      "extends": ["group:allNonMajor"],
      "commitMessage": "gomod - {{{commitMessagePrefix}}} {{{commitMessageAction}}} {{{commitMessageTopic}}} {{{commitMessageExtra}}} {{{commitMessageSuffix}}}",
      "additionalBranchPrefix": "{{parentDir}}-",
      "packagePatterns": [".*"]
    }
  ]
}
```
