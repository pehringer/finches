import math
import random

def affine(x):
	# f(x) = ax + b
	return 2*x + 3

def quadratic(x):
	# f(x) = ax^2 + bx + c
	return 2*x**2 + 3*x + 4

def piecewise(x):
	#        { 0.5x + sin(x) if x < 0
	# f(x) = { x^2 − 2       if 0 <= x < 1
	#        { tan(0.2x)     if x >= 1
	if x < 0:
		return 0.5 * x + math.sin(x)
	elif x < 1:
		return x**2 - 2
	else:
		return math.tan(0.2 * x)

def polynomial(x, y, z):
	# f(x, y, z) = 2*x*y + 3*z^2 - x + 5
	return 2 * x * y + 3 * z**2 - x + 5

def nested(x, y):
	if x < 0:
		if y > -1:
			return 0.5 * x + math.sin(y)
		else:
			return x**2 - 2 * y
	else:
		if y < 1:
			return math.tan(0.2 * x) - y
		else:
			log_arg = abs(x * y) + 1
			return math.log(log_arg) + 2

with open("nested.csv", "w") as f:
	f.write("inputX,inputY,output\n")
	for i in range(0, 1000):
		x = random.uniform(-5, 5)
		y = random.uniform(-5, 5)
		output = nested(x, y)
		f.write(f"{round(x, 10)},{round(y, 10)},{round(output, 10)}\n")
