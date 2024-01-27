# todoist-to-org-mode

## Motivation
Recently discovered [orgzly](https://orgzly.com/) and decided to make the switch to org-mode based tasks.

## Manual

### Dependencies
- Golang - The script is using `go v1.21` but should be fairly backward compatible with older versions
- Todoist API - The script uses `v2` of the [todoist api](https://developer.todoist.com/rest/v2/#delete-a-section).

## Flags
- Api key
  - name: `--key`
  - short: `-k`

| full name  | short name | description                                                                                                                            |
|------------|------------|----------------------------------------------------------------------------------------------------------------------------------------|
| `--key`    | `-k`       | api key for your todoist account. Can be found [here](https://app.todoist.com/app/settings/integrations/developer). Can have multiple. |
| `--output` | `-o`       | Output file name. If left empty will print results in console.                                                                         |

## Improvements
- Currently recurring tasks are not supported.

## Similar projects
- https://github.com/pmiddend/org-todoist
- https://github.com/drmfinlay/todoist-org-mode
