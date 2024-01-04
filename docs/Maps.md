The arenas, where the game encounters take place.

### Gusanos maps
Their minimal structure consists of:
- level.png - sprite with actual map layer being rendered on the screen,
- material.png - sprite with color-coded collision blocks, defining how the players and game objects will collide with different map parts,
- config.cfg - configuration file written in [[OMFGScript]].

### Legacy Liero maps (LEV)
These are map files from OG Liero. They come in single .lev file, no idea for now how to parse them, the first glance tells me that it's some kind of binary, but it may be a lot simpler in reality (I just hope so).

### Collisions
Going to try SolarLune/resolv for collisions detection, check if it's suitable for my needs, and maybe learn how to tackle this by myself (in the future).