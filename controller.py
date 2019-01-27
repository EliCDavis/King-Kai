import pyxinput
import sys
import time

MyVirtual = pyxinput.vController()
f = open("debug.txt", "a")

def otherInput(command):
    buttons_to_reset = ["null", "null"]

    if command == "assist1":
        MyVirtual.set_value('BtnShoulderL', 1)
        buttons_to_reset[0] = 'BtnShoulderL'

    elif command == "assist2":
        MyVirtual.set_value('TriggerL', 1)
        buttons_to_reset[0] = 'TriggerL'

    elif command == "dragonrush": # done
        # rb but also is l+m
        MyVirtual.set_value('BtnShoulderR', 1)
        buttons_to_reset[0] = 'BtnShoulderR'

    elif command == "superdash": # done
        # rt but also is h+s
        MyVirtual.set_value('TriggerR', 1)
        buttons_to_reset[0] = 'TriggerR'

    elif command == "vanish": # done
        # m+h
        MyVirtual.set_value('BtnY', 1) 
        MyVirtual.set_value('BtnB', 1)
        buttons_to_reset[0] = 'BtnY'
        buttons_to_reset[1] = 'BtnB'

    elif command == "sparking":
        MyVirtual.set_value('TriggerR', 1)
        buttons_to_reset[0] = 'TriggerR'
        MyVirtual.set_value('BtnShoulderR', 1)
        buttons_to_reset[1] = 'BtnShoulderR'

    return buttons_to_reset

def move(direction):
    if direction == "left":
        MyVirtual.set_value('AxisLx', -1)
        return 'AxisLx'

    elif direction == "right":
        MyVirtual.set_value('AxisLx', 1)
        return 'AxisLx'

    elif direction == "down":
        MyVirtual.set_value('AxisLy', -1)
        return 'AxisLy'

    elif direction == "up":
        MyVirtual.set_value('AxisLy', 1)
        return 'AxisLy'

    elif direction == "neutral":
        return 'null'
    
    f.write("no direction: " + str(direction) + "\n")
    return "null"

def attack(button):

    if button == "special":
        MyVirtual.set_value('BtnA', 1)
        return 'BtnA'

    elif button == "light":
        MyVirtual.set_value('BtnX', 1)
        return 'BtnX'

    elif button == "medium":
        MyVirtual.set_value('BtnY', 1)
        return 'BtnY'

    elif button == "heavy":
        MyVirtual.set_value('BtnB', 1)
        return 'BtnB'

    f.write("no attack button: " + str(button) + "\n")
    return "null"


def special(time_elapsed, button, circle):

    update_rate = 70

    buttons_to_reset = ["null", "null"]

    midDirection = "left" if circle == "quarterback" else "right"

    if time_elapsed < update_rate:
        move("down")

    elif  time_elapsed < update_rate * 2:
        move("down")
        buttons_to_reset[0] = move(midDirection)

    elif time_elapsed < update_rate * 3:
        MyVirtual.set_value('AxisLy', 0)

    else:
        buttons_to_reset[0] = move(midDirection)
        buttons_to_reset[1] = attack(button)

    return buttons_to_reset


def bar_button(button):
    if button == "one":
        MyVirtual.set_value('BtnShoulderR', 1)
        return  "BtnShoulderR"

    elif button == "two":
        MyVirtual.set_value('TriggerR', 1)
        return  "TriggerR"

    f.write("no bar button: " + str(button) + "\n")
    return "null"
    

def bar(time_elapsed, button, circle):

    update_rate = 50

    buttons_to_reset = ["null", "null"]

    midDirection = "left" if "back" in circle else "right"

    if "quarter" in circle:
        if time_elapsed < update_rate:
            move("down")

        elif  time_elapsed < update_rate * 2:
            move("down")
            buttons_to_reset[0] = move(midDirection)

        elif time_elapsed < update_rate * 3:
            MyVirtual.set_value('AxisLy', 0)

        else:
            buttons_to_reset[0] = move(midDirection)
            buttons_to_reset[1] = bar_button(button)

    if "half" in circle:
        if time_elapsed < update_rate:
            move("down")

        elif  time_elapsed < update_rate * 2:
            move("down")
            buttons_to_reset[0] = move(midDirection)

        elif  time_elapsed < update_rate * 3:
            buttons_to_reset[0] = move(midDirection)
            MyVirtual.set_value('AxisLy', 0)

        elif  time_elapsed < update_rate * 4:
            move("up")
            buttons_to_reset[0] = move(midDirection)

        elif time_elapsed < update_rate * 5:
            buttons_to_reset[0] = move("up")
            MyVirtual.set_value('AxisLx', 0)

        else:
            buttons_to_reset[0] = move("up")
            buttons_to_reset[1] = bar_button(button)


    return buttons_to_reset


try:
    for line in sys.stdin:

        input = line[:-1].split()

        buttons_to_reset = ["null", "null" , "null"]

        duration = int(input[0])

        typeCommand = input[1]

        start_time = time.time()

        while (time.time() - start_time) * 1000.0 < duration:
            if typeCommand == "attack":
                unpacked = input[2].split("-")
                buttons_to_reset[0] = attack(unpacked[0])
                buttons_to_reset[1] = move(unpacked[1])

            elif typeCommand == "command":
                buttons_to_reset = otherInput(input[2])

            elif typeCommand == "move":
                buttons_to_reset[0] = move(input[2])

            elif typeCommand == "jump":
                buttons_to_reset[0] = move("up")
                buttons_to_reset[0] = move(input[2])

            elif typeCommand == "special":
                unpacked = input[2].split("-")
                time_elapsed = (time.time() - start_time) * 1000.0
                buttons_to_reset = special(time_elapsed, unpacked[0], unpacked[1])

            elif typeCommand == "bar":
                unpacked = input[2].split("-")
                time_elapsed = (time.time() - start_time) * 1000.0
                buttons_to_reset = bar(time_elapsed, unpacked[0], unpacked[1])

            elif typeCommand == "reset":
                MyVirtual.set_value('BtnBack', 1)
                buttons_to_reset[0] = 'BtnBack'
                buttons_to_reset[1] = move("down")


        for button in buttons_to_reset:
            if button != "null":
                MyVirtual.set_value(button, 0)

except Exception as err:
    f.write("\nan error occurred:\n\t" + str(err))
    f.flush()
    f.close()