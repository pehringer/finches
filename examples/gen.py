import random
import numpy as np

def ackley1d(x):
	"""Ackley function (1D): f(x) = -20*exp(-0.2*|x|) - exp(cos(2*pi*x)) + 20 + e."""
	return -20 * np.exp(-0.2 * np.abs(x)) - np.exp(np.cos(2 * np.pi * x)) + 20 + np.e

def rastrigin1d(x):
	"""Rastrigin function (1D): f(x) = x^2 - 10*cos(2*pi*x) + 10."""
	return x**2 - 10 * np.cos(2 * np.pi * x) + 10

def piecewise_complex(x):
    """
    A more complex piecewise function demonstrating varied behaviors:

    f(x) =
      sin(x)                  if x < -2
      x^3 + x                 if -2 <= x < 0
      2*x + 1                 if 0 <= x < 2
      x * log(x)              if 2 <= x < 5
      exp(-x) + sin(x)        if x >= 5
    """
    if x < -2:
        return np.sin(x)
    elif x < 0:
        return x**3 + x
    elif x < 2:
        return 2 * x + 1
    elif x < 5:
        return x * np.log(x)
    else:
        return np.exp(-x) + np.sin(x)

with open("ackley1d.csv", "w") as f:
	f.write("input,output\n")
	for i in range(0, 256):
		input = random.uniform(-10, 10)
		output = ackley1d(input)
		f.write(f"{round(input, 8)},{round(output, 8)}\n")
