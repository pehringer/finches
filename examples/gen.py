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

with open("piecewise.csv", "w") as f:
	f.write("input,output\n")
	for i in range(0, 1000):
		input = random.uniform(-5, 5)
		output = piecewise(input)
		f.write(f"{round(input, 10)},{round(output, 10)}\n")
