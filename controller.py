import pyxinput
import sys

MyVirtual = pyxinput.vController()

# f = open("debug.txt", "a")
# f.write("aaaaa")
# f.flush()
# while True:
#     line = sys.stdin.read(-1)
#     f.write("line: " + line)
#     f.flush()
#     if line == "left\n":
#         MyVirtual.set_value('AxisLx', 1)

for line in sys.stdin:
    # f.write("line: " + line)
    # f.flush()
    if line == "left\n":
        MyVirtual.set_value('AxisLx', 1)
    # sys.stdin.read()



# while True:
#     MyVirtual.set_value('AxisLx', -1)
