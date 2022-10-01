The `sam` matrix consists of the following column elements. Each can have a value of 0 (`false`) or 1 (`true`)

`[0]` - Is Available

`[1]` - Has Ammo

`[2]` - Has Power

`[3]` - Is Connected to IADS

`[4]` - Has High Range capabilities (30,000+ ft)

`[5]` - Has Mid-High Alt capabilities (20,000+ ft)

`[6]` - Has Mid Alt capabilities (12,000+ ft)

`[7]` - Has Low Range capabilties (<50ft )

An example for an SA-10 would be:

11101111

- Is alive
- Has ammo and power
- Operating autonomously 
- Has all tracked capabilities

An example of an SA-6, suppressed SA-2, and SA-9 would be:

11110110

00000000

11110000

Which produces a mask of:
11110110

If a fighter is detected by an EWR the system will calculate it's future position in 7 minutes, 


If a fighter is coming in at above 30,000ft, the system will evaluate elem 4 (High Range). If any elem of the mask is 0, the IADS is considered `AtRisk`and the system will ready alert fighters for dispatch if the current available pool is less than 1. If elem 4 of the mask is 0, the IADS will consider the sector in front of the fighter to be `Defeatable` and directly push tasks to fighters in the available pool to intercept the aircraft.