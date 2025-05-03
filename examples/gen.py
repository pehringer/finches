import random

def unflattenable(x, y, z):
    if x > 0:
        if y > 0:
            if z > 0:
                return x * y * z
            else:
                return x + y - z
        else:
            if x + y < 0:
                return (x - y) * z
            else:
                return (x + z) ** 2
    else:
        if z > 5:
            if y < -5:
                return abs(x - y) + z
            else:
                return (z - x) / (abs(y) + 1)
        else:
            if x * y < z:
                return (x + y + z) ** 2
            else:
                return x - y + z


with open("unflattenable.csv", "w") as f:
	f.write("input0,input1,input2,expected\n")
	for i in range(0, 100):
		input0 = random.uniform(-10, 10)
		input1 = random.uniform(-10, 10)
		input2 = random.uniform(-10, 10)
		expected = unflattenable(input0, input1, input2)
		f.write(f"{round(input0, 4)},{round(input1, 4)},{round(input2, 4)},{round(expected, 4)}\n")



