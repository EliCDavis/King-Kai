# King Kai

*Leave your warriors with King Kai and he'll teach them some new techniques (combos)*

King Kia is a reinforcement learning agent meant to discover combos in Dragonball FighterZ

Work in progress...

## Requirements

 * Golang
 * FighterZ
 * 1920x1080 resolution, no lower *or* higher

## Notes..

Tesseract gave around 3-4 frames a second. Very spotty..

Custom OCR gives around 11-19 frames a second..

Screen capturing only the regions I need instead of the entire screen increases framerate to 30-142 fps...

Here's a really cool design choice that FighterZ did but hurts me:

> They put the combo counter behind the special effects. That means it can get hidden by the effects. But it's cool cause after you do a super move, and the dust settles, you see your resulting combo counter. This sucks for me though because I can't just look at a specific part of the screen to see if the combo ended.

