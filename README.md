# go-gusanos
Golang recreation of [gusanos](https://github.com/wesz/gusanos), the Liero clone.

## Plan

1. Roughly rewrite the original codebase, using 
    - the [ebitengine](github.com/hajimehoshi/ebiten) instead of [Allegro](github.com/liballeg/allegro5),
    - no sound engine (for now),
    - [nano](github.com/lonng/nano) instead of zoidcom (the original webpage is available only through Wayback Machine)
    - comment `CHECK` for any places in code, that I should go back to understand more or refactor.

2. Produce executable running without (many) errors, if it will be somehow playable at this stage, that will be a plus.

3. Continuous refactor, hopefully documentation.

4. Take my time.

5. Every legit pull request is more than welcome.