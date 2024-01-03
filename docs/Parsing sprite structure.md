Sprite maps from default gusanos mod are built in specific way, which needs a custom loader to fetch all the data from them.
Sprites have uneven widths: on the top and left white border, the center point of each sprite is marked with red (255,0,0) pixel, the split point with black (0,0,0) one. Transparency is marked with magenta (255,0,255).

### Algorithm (first thought)
1. Load pixels from row x:0. Check pixel value, store red as anchor x points, black as split x points.
2. Load pixels from column y:0. Check pixel value, store as above.
3. Create new image without top and left border from original one.
4. Store new image height and width.
5. Return struct with:
- cropped image,
- anchor points,
- split points,

### Issues
Getting sub sprite, by (row, col) - using split points.