# Wanted Features

`All configuration should follow top down inheritance.`
- [ ] Configuration
    - [ ] Global Default Configuration
    - [ ] User Definable IADS-wide Settings
    - [ ] User Definable Command Center Settings
    - [ ] User Definable SAM site settings
    - [ ] User Definable settings filtered by SAM name or Descriptor
    - [ ] SAM Tactics
        - [ ] DCS Default (Always on)
        - [ ] Radars off unless engaging
        - [ ] Active / Standby clustering
        - [ ] SAM Trapping / Ambush setup
        - [ ] Engage at rMAX, NEZ, specific % range, or random value between rMAX and NEZ
- [ ] Define Objectives for IADS
    - [ ] Defend self by default
    - [ ] Defend airbases 
    - [ ] Defend defined coordinates
    - [ ] Defend a unit list
- [ ] Server Autostart Mode
- [ ] Mission Drop-in Mode
- [ ] Support Multiple Configurations
- [ ] Unit Discovery
    - [ ] Discover units by Name Prefix
    - [ ] Discover units by Descriptor
    - [ ] Discover power and communication infrastructure by static object type
    - [ ] Configruation option to add new SAM units automatically after simulation start
    - [ ] Ability to handle unknown unit types
- [ ] Support for mobile SAM sites
- [ ] Support for mixed unit groups that include SHORAD units
- [ ] HARM and JSOW Defense
    - [ ] SAM sites assignable as point defense
    - [ ] Point Defense changes priorities based on state of ward
    - [ ] Ward changes priorities based on state of point defense
    - [ ] HARM track prediction to determine which SAM site is in danger
- [ ] Target Detection and Prioritization
    - [ ] Prioritize by vector
    - [ ] Prioritize by threat to specified defense objective
- [ ] Interceptor Dispatcher
    - [ ] Support for a JEZ
    - [ ] Config options to prioritize use of fighter or use of SAMs
    - [ ] Function to pop task to Interceptor to fly for best intercept
- [ ] Ability to evaluate health state of IADS
    - [ ] Dispatch Interceptor if IADS does not believe it can counter them
    - [ ] Ability to repair SAM sites
- [ ] Power Grid Simulation 
    - [ ] Ability to automatically generate power grids based on logical graphs
    - [ ] Ability to detect suitable backup generators for SAM sites
    - [ ] Disable SAMs if power is removed
- [ ] Communication Network Simulation
    - [ ] Ability to automatically generate communication and command structures based on logical graphs
    - [ ] Ability to manually overide autogen and manually define some connections
    - [ ] Support for optionally enforcing channel and range limits