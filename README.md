# King Kai

*Leave your warriors with King Kai and he'll teach them some new techniques*

King Kia is a reinforcement learning agent meant to discover combos in Dragonball FighterZ

Work in progress...

## Requirements

 * Golang
 * FighterZ
 * 1920x1080 resolution, no lower *or* higher
 * **Exactly** Python 3.6.1, any other version doesn't work

 1. Follow setup instructions for [PYXInput](https://github.com/bayangan1991/PYXInput)
 2. Training settings:
    * Turn off the ability to summon shenron
    * Make sure the combo damage info is visible
    * Quick health recovery
    * Make characters start with 7 bars
 3. FighterZ can't be on fullscreen, must be on windowed fullscreen

### Testing

Go to training mode and pick your characters, stage, and whatever. Once the stage has loaded unplug your controller. Pull up a console and type something like:

```bash
echo "100 attack light-down" | python .\controller.py
```

What you should see is your character do a 2L. If they don't do anything you failed to set up PYXInput correctly.

## Notes..

Tesseract gave around 3-4 frames a second. Very spotty..

Custom OCR gives around 11-19 frames a second..

Screen capturing only the regions I need instead of the entire screen increases framerate to 30-142 fps...

Here's a really cool design choice that FighterZ did but hurts me:

> They put the combo counter behind the special effects. That means it can get hidden by the effects. But it's cool cause after you do a super move, and the dust settles, you see your resulting combo counter. This sucks for me though because I can't just look at a specific part of the screen to see if the combo ended.

Any kind of keyboard emulation does not seem to work -_- Despite keyboard actually being able to work

