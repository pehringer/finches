# fahrenheit.csv

Columns: C, F
- C (Celsius): uniform random from −40 to 100
- F (Fahrenheit)
```
F = 1.8 * C + 32
```

# heat.csv

Columns: F, RH, HI
- F (Fahrenheit): uniform random from 80 to 110
- RH (Relative Humidity %): uniform random from 40 to 100
- HI (Heat Index)
```
HI = -42.379 + 2.04901523 * F + 10.14333127 * RH - 0.22475541 * F * RH - 0.00683783 * F^2 - 0.05481717 * RH^2 + 0.00122874 * F^2 * RH + 0.00085282 * F * RH^2 - 0.00000199 * F^2 * RH^2
```

# savings.csv

Columns: D, B
- D (Deposit $): uniform random from 0 to 500
- B (Balance $)
- N (Deposits): 3 to 10 per sequence
```
B = 100 * 1.05^N + D1 * 1.05^(N-1) + D2 * 1.05^(N-2) + ... + DN
```

# portfolio.csv

Columns: P, S, V
- P (Price $): random walk starting from 40 to 60, changing up to ±5% each period
- S (Shares bought): uniform random whole number from 0 to 10
- V (Value $)
- N (Periods): 3 to 12 per sequence

```
V = 1000 - 4.95 * N - (S1 * P1 + S2 * P2 + ... + SN * PN) + (S1 + S2 + ... + SN) * PN
```

# heartbeat.csv

Columns: T, ECG
- T (Time): uniform random from 0 to 1
- ECG (Voltage)
```
ECG = 0.15 * exp(-(T - 0.20)^2 / 0.00125) - 0.15 * exp(-(T - 0.36)^2 / 0.0002) + exp(-(T - 0.40)^2 / 0.000288) - 0.25 * exp(-(T - 0.44)^2 / 0.0002) + 0.30 * exp(-(T - 0.70)^2 / 0.0032)
```

# shipping.csv

Columns: O, T
- O (Order $): uniform random from 0 to 100
- T (Total $)
```
if O >= 50: T = O
else:       T = O + 5.99
```

# coffee.csv

Columns: D, Q, T
- D (Drink code): 1 to 3 half the time, otherwise 0 or 4 to 9
- Q (Quantity): uniform random whole number from 1 to 10
- T (Total $)
```
switch D:
  case 1:  T = 4.25 * Q
  case 2:  T = 3.95 * Q
  case 3:  T = 4.75 * Q
  default: T = 2.50 * Q
```
