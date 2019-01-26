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

try:
    for line in sys.stdin:

        input = line[:-1].split()

        buttons_to_reset = ["null", "null"]

        duration = int(input[0])

        typeCommand = input[1]

        start_time = time.time()


        while (time.time() - start_time) * 1000.0 < duration:
            if typeCommand == "attack":
                unpacked = input[2].split("-")
                buttons_to_reset[0] = attack(unpacked[0])
                buttons_to_reset[1] = move(unpacked[1])

            elif typeCommand == "reset":
                MyVirtual.set_value('BtnBack', 1)
                buttons_to_reset[0] = 'BtnBack'

        if buttons_to_reset[0] != "null":
            MyVirtual.set_value(buttons_to_reset[0], 0)
        if buttons_to_reset[1] != "null":
            MyVirtual.set_value(buttons_to_reset[1], 0)

except Exception as err:
    f.write("\nan error occurred....\n" + err.text)
    f.flush()
    f.close()