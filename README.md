# footballsimulation

A football simulation engine in golang.

Terms in the project.
- Formation : A Set of 11 Roles. 
- Role : A Position + a Player in a Formation.
- Player : A Player who has an Ability and skilled positions.
- Ability : A set of Attributes representing mental, physical and technical skills.
- Attribute : A skill that has a name and a point(ex, pass 10)

- Zones : A field or stadium that the game can be played. A Zones consists of 10 Zone(1 Zones = 10 Zone).
- Zone : A Zone means a part of the field, which is GK, DL, DR, DC, ML, MC, MR, FL, FC, FR.
- Entry : A Player + a position + game stats in a game. An Entry can be involved with multiple zones.

