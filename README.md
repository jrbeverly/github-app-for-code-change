# GitHub App for Generated Commits

Running through a bunch of things to be done with this

## Notes

- Is it possible to push empty commits, which would need to be handled
- Commits should be crafted first, then attempts being made to apply the change
- Would need to have better visibility into the "crafted" files that will be committed
- Does this really require a GitHub Application? Feels like this could be decoupled
- Need to ensure that requests with the API go through the proper retry processes
- Need to ensure that the configuration is split from auth, so its easier to rotate the auth without fiddling with the config
- Should the app really be responsible for the work-component of crafting the change?
- Should all changes be made as a pull request + request to merge from the GitHub App? This would avoid conflicts as much, and let GitHub be responsible for changes?

Does this kind of thing really need a GitHub Application? Feels like this is something that can be split away from GitHub as a kind of entity that is responsible for the following actions:

1) Pull the original source from the git repositories
2) Apply the processes on the source files (regular text processing)
3) Apply the change, and publish using standard git procedures

This wouldn't require any additional overhead complexity with communicating with the GitHub repository. Or could this be something that is abstracted into a common application that is responsible for something like

- Receive a set of "Changes" as proposed within git
- Perform the change on behave of this system

This would allow commits to be "queued" and reduce a lot of the headache with dealing with this processes within the tool, and instead allow other tools to just communicate with the "Git Operations" API, which is much more fault tolerant.
