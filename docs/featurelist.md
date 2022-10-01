# Core Features

- All configuration should follow top down inheritance. Configuration should first contain a global default that creates a functional IADS > User set IADS settings > User set command center settings > User set SAM site settings

- Ability to drop onto a DCS server and create IADS on all missions without the need for Lua to explicitly enable it.

- Ability to swap between configurations

- Ability to find and setup SAM and EWR sites by either name tagging or descriptor tagging

- Ability to defend against HARMs

- Ability to sort and prioritize threats by their vector, not just their presence in a threat ring.

- Ability to dispatch interceptors 

- Ability to estimate the need to dispatch interceptors if the IADs is not well suited for the threat

- Ability to generate logical power and communication networks use abstracts and static objects to find best paths of connection.

- Ability to manually define power and communication connections as desired by the mission maker