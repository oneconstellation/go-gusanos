A mod contains its own sprites - few of them seem as required, like:
- skin.png - a sprite map depicting a player character, in *default* mod has dimensions 42x82, and contains 36 sprites
	- This one is tricky, because of uneven frame widths. Need to use the markers on the offset to grab the sprite bounds and center point.