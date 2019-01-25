import pyxinput
import sys
import time

MyVirtual = pyxinput.vController()

try:
    for line in sys.stdin:

        input = line[:-1].split()

        command = input[0]
        duration = int(input[1])

        start_time = time.time()

        button_to_reset = "null"

        while (time.time() - start_time) * 1000.0 < duration:
            if command == "left":
                MyVirtual.set_value('AxisLx', -1)
                button_to_reset = 'AxisLx'

            elif command == "right":
                MyVirtual.set_value('AxisLx', 1)
                button_to_reset = 'AxisLx'

            elif command == "down":
                MyVirtual.set_value('AxisLy', -1)
                button_to_reset = 'AxisLy'

            elif command == "up":
                MyVirtual.set_value('AxisLy', 1)
                button_to_reset = 'AxisLy'

            elif command == "special":
                MyVirtual.set_value('BtnA', 1)
                button_to_reset = 'BtnA'

            elif command == "light":
                MyVirtual.set_value('BtnX', 1)
                button_to_reset = 'BtnX'

            elif command == "medium":
                MyVirtual.set_value('BtnY', 1)
                button_to_reset = 'BtnY'

            elif command == "heavy":
                MyVirtual.set_value('BtnB', 1)
                button_to_reset = 'BtnB'

        if button_to_reset != "null":
            MyVirtual.set_value(button_to_reset, 0)

except Exception as err:
    f = open("debug.txt", "a")
    f.write("\nan error occurred....\n" + err.text)
    f.flush()
    f.close()