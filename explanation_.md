## Exercise 01 — Basic Variable Types
- **Story in my words:** A student record needs an age, a body mass, whether the student has checked in, and a student ID. I store each in a variable of a suitable type and print them on one line.
- **Input(s) and their types:** `age` = 19 (int), `mass` = 62.5 (float64), `checkedIn` = true (bool), `studentID` = "00123" (string). No user input.
- **Processing / logical rule:** No calculation. Go infers each type from its value, and the ID is a string so the leading zeros are kept.
- **Expected output (before coding):** `19 | 62.5 | true | 00123`
- **Actual output:** `19 | 62.5 | true | 00123`
- **Normal test:** The values above print exactly as written.
- **Boundary or invalid test:** If `studentID` were the int `00123`, it would print `83` (octal literal), which shows why an ID must be a string.
- **One assumption or limitation:** The values are hard-coded, so the program cannot handle different students without editing the code.
- **What I changed in `main.go`:** Declared the four variables with `:=` and printed them with `fmt.Println`, using `"|"` as a separator.
- **Why I believe the result is correct:** Each printed value matches what I assigned, and the ID keeps its leading zeros.

## Exercise 02 — Price Calculation
- **Story in my words:** A customer buys 2 items at 18000 each and has a discount of 5000. I compute the subtotal and then the final total.
- **Input(s) and their types:** `unitPrice` = 18000 (int), `quantity` = 2 (int), `discount` = 5000 (int).
- **Processing / logical rule:** `subtotal = unitPrice * quantity`; `total = subtotal - discount`.
- **Expected output (before coding):** Subtotal: 36000, Total: 31000
- **Actual output:** `Subtotal: 36000` and `Total: 31000`
- **Normal test:** 18000 × 2 = 36000, and 36000 − 5000 = 31000.
- **Boundary or invalid test:** With `discount` = 40000, the total would be −4000, which makes no sense for a price.
- **One assumption or limitation:** The program assumes the discount never exceeds the subtotal, and it uses int (no cents or decimals).
- **What I changed in `main.go`:** Added the `subtotal` calculation line and subtracted the discount to get `total`.
- **Why I believe the result is correct:** I checked the arithmetic by hand and it matches the printed values.

## Exercise 03 — Factor Check
- **Story in my words:** Given two numbers, I want to know whether the first divides the second evenly (x is a factor of y).
- **Input(s) and their types:** `x` and `y`, both int, read with `fmt.Scan`.
- **Processing / logical rule:** If either number is ≤ 0, print `invalid input` and stop. Otherwise print whether `y % x == 0`.
- **Expected output (before coding):** For `3 12` → `true`.
- **Actual output:** `true`
- **Normal test:** `3 12` → `true`; `5 12` → `false`.
- **Boundary or invalid test:** `0 5` → `invalid input`, which avoids a modulo-by-zero panic; `-3 12` → `invalid input`.
- **One assumption or limitation:** Only positive integers are accepted. If the user types text, `Scan` fails, both values stay 0, and the program prints `invalid input`.
- **What I changed in `main.go`:** Added the validation check before the modulo so division by zero can never happen.
- **Why I believe the result is correct:** 12 / 3 leaves remainder 0, so `true` is right; 12 / 5 leaves remainder 2, so `false` is right.

## Exercise 04 — Population Change
- **Story in my words:** A town starts with 1000 people. Over a year, some are born, some arrive, some die, and some leave. I compute the final population.
- **Input(s) and their types:** `start` = 1000, `births` = 20, `arrivals` = 5, `deaths` = 10, `departures` = 4 (all int).
- **Processing / logical rule:** `end = start + births + arrivals - deaths - departures`.
- **Expected output (before coding):** Population: 1011
- **Actual output:** `Population: 1011`
- **Normal test:** 1000 + 20 + 5 − 10 − 4 = 1011.
- **Boundary or invalid test:** If deaths and departures were large enough (e.g., 2000), the population would go negative.
- **One assumption or limitation:** Values are hard-coded, and there is no check that the result stays at or above zero.
- **What I changed in `main.go`:** Wrote the `end` formula adding the two increases and subtracting the two decreases.
- **Why I believe the result is correct:** Increases total 25, decreases total 14, and 1000 + 25 − 14 = 1011.

## Exercise 05 — Fibonacci-Style Sequence
- **Story in my words:** Starting with 5 and 3, each new number is the sum of the two before it. I print the next three terms.
- **Input(s) and their types:** `first` = 5, `second` = 3 (both int).
- **Processing / logical rule:** `third = first + second`; `fourth = second + third`; `fifth = third + fourth`.
- **Expected output (before coding):** `8 11 19`
- **Actual output:** `8 11 19`
- **Normal test:** 5 + 3 = 8, 3 + 8 = 11, 8 + 11 = 19.
- **Boundary or invalid test:** Starting values of 0 and 0 would give `0 0 0`, a sequence that never grows.
- **One assumption or limitation:** Only three new terms are computed, and each step is written by hand instead of in a loop.
- **What I changed in `main.go`:** Added the three lines that compute `third`, `fourth`, and `fifth`, each using the previous two values.
- **Why I believe the result is correct:** Each term equals the sum of the two before it, which I checked by hand.

## Exercise 06 — Weight on Earth and Mars
- **Story in my words:** A person with a mass of 70 kg has a weight (a force in newtons) on Earth. On Mars the same person weighs 38% as much.
- **Input(s) and their types:** `massKg` = 70.0, `earthGravity` = 9.8, `marsFactor` = 0.38 (all float64).
- **Processing / logical rule:** `earthWeight = massKg * earthGravity`; `marsWeight = earthWeight * marsFactor`. Print each with one decimal place.
- **Expected output (before coding):** Earth: 686.0 N, Mars: 260.7 N
- **Actual output:** `Earth: 686.0 N` and `Mars: 260.7 N`
- **Normal test:** 70 × 9.8 = 686; 686 × 0.38 = 260.68, which rounds to 260.7.
- **Boundary or invalid test:** A mass of 0 gives 0.0 N on both planets; a negative mass would give a negative weight, which is not physical.
- **One assumption or limitation:** Gravity values are rounded constants, and the output is rounded to one decimal place by `%.1f`.
- **What I changed in `main.go`:** Added the `earthWeight` line and used `Printf` with `%.1f` for formatting.
- **Why I believe the result is correct:** Weight = mass × gravity, and the Mars result is a fixed 38% of the Earth result, which my hand calculation confirms.

Exercise 07 — Genetic Algorithm Basics
- **Story in my words:** Two bit-string "parents" are combined into a child (crossover), then one bit of the child is forced to 1 (mutation). A fitness function scores a string by counting its 1s.
- **Input(s) and their types:** `a` = "1001", `b` = "0110" (both string).
- **Processing / logical rule:** `child = a[:2] + b[2:]`; `mutated = child[:1] + "1" + child[2:]`; `fitness` counts the '1' characters.
- **Expected output (before coding):** Child 1010 with fitness 2; mutated 1110 with fitness 3.
- **Actual output:** Parent A: 1001
Parent B: 0110
Child: 1010 fitness: 2
Mutated: 1110 fitness: 3
- **Normal test:** The first half of A ("10") plus the second half of B ("10") gives "1010".
- **Boundary or invalid test:** If the strings were shorter than 2 characters, the slicing would panic; `fitness` of "0000" is 0 and of "1111" is 4.
- **One assumption or limitation:** All strings must have length 4 and contain only '0' and '1', and the mutation position is fixed instead of random.
- **What I changed in `main.go`:** Wrote the `child` crossover line and the `mutated` line that replaces index 1 with "1".
- **Why I believe the result is correct:** I sliced the strings by hand, and counting the 1s gives 2 for "1010" and 3 for "1110".

## Exercise 08 — Cube Check
- **Story in my words:** I want to know whether the second number is the cube of the first.
- **Input(s) and their types:** `x` and `y`, both int, read with `fmt.Scan`.
- **Processing / logical rule:** If `x * x * x == y`, print `true`; otherwise print `false`.
- **Expected output (before coding):** For `3 27` → `true`.
- **Actual output:** `true`
- **Normal test:** `3 27` → `true`; `3 28` → `false`.
- **Boundary or invalid test:** `0 0` → `true`; `-2 -8` → `true` (negative cubes work); a very large `x` can overflow int and give a wrong answer.
- **One assumption or limitation:** Only integers are supported, and there is no overflow protection for large values.
- **What I changed in `main.go`:** Replaced the condition with `x * x * x == y` (multiplying x by itself three times).
- **Why I believe the result is correct:** 3 × 3 × 3 = 27, so the first test is true, and 27 ≠ 28, so the second is false.

## Exercise 09 — Full Genetic Algorithm
- **Story in my words:** A population of 6 random 4-bit strings evolves for up to 10 generations. Each generation picks parents, combines them, and sometimes flips bits, trying to reach "1111".
- **Input(s) and their types:** Constants: `populationSize` = 6, `chromosomeSize` = 4, `maxGenerations` = 10 (int), `mutationRate` = 0.1 (float64). Candidates are strings of '0'/'1'.
- **Processing / logical rule:** Score each candidate by counting 1s and print the best. Stop if the best equals 4. Otherwise build a new population: `selectParent` takes the fitter of two random picks, `crossover` joins the first half of one parent with the second half of the other, and `mutate` flips each bit with a 10% chance.
- **Expected output (before coding):** Up to 10 lines of the form `Generation N: Best = XXXX, Fitness = F`, ending with `Optimal solution found!` if a "1111" appears.
- **Actual output:** Paste your own run here. Example shape: `Generation 1: Best = 1110, Fitness = 3` … `Generation 3: Best = 1111, Fitness = 4` then `Optimal solution found!`
- **Normal test:** Run it several times. Fitness usually rises across generations and often reaches 4 before generation 10.
- **Boundary or invalid test:** If all 6 random candidates are "0000", the best fitness stays 0, and since the code only updates the best when `score > bestFitness`, it prints `Best = ` with an empty string.
- **One assumption or limitation:** Results are random, so runs differ. There is no elitism, so the best candidate can be lost between generations, and crossover always splits at position 2.
- **What I changed in `main.go`:** Added comments explaining parent preference, crossover, mutation, and the stop condition.
- **Why I believe the result is correct:** Fitness is the count of 1s, the stop check compares it with `chromosomeSize`, and the printed best candidate's fitness matches its number of 1s.

## Exercise 10 — Go Basic Types Demo
- **Story in my words:** A tour of Go's basic types: integer division, floating-point division, booleans, strings, runes, and constants.
- **Input(s) and their types:** `a` = 7, `b` = 2 (int); `score` = 72, `attendance` = 70 (int); `radius` = 5.0 (float64); `letter` = 'A' (rune); `pi` = 3.14 (constant).
- **Processing / logical rule:** Integer `/` and `%` give quotient and remainder; `float64` conversion gives true division; `passed = score >= 60 && attendance >= 75`; strings are joined with `+`; a rune plus 1 is the next character; `area = pi * radius * radius`.
- **Expected output (before coding):**int: 7 / 2 = 3, 7 % 2 = 1
float64: 3.5
bool: passed = false
string: Hello, Go (len 9)
rune: A -> 65 -> B
const: area = 78.5
- **Actual output:** Same as expected.
- **Normal test:** 7 / 2 = 3 remainder 1; converted to float it is 3.5; 3.14 × 5 × 5 = 78.5.
- **Boundary or invalid test:** `attendance` = 75 would make `passed` true (boundary of `>=`); `b` = 0 would crash with an integer division by zero.
- **One assumption or limitation:** `pi` = 3.14 is only an approximation, and integer division drops the decimal part.
- **What I changed in `main.go`:** Added the section comments and the lines for each type, including `float64()` conversions and the `_, _, _ = ...` line so unused variables don't cause a compile error.
- **Why I believe the result is correct:** I checked each line by hand: attendance 70 is below 75, so `passed` is false; "Hello, Go" has 9 characters; 'A' is code 65, so the next rune is 'B'.
.

A — Git and GitHub (Terminal)
1. fork
2. clone
3. switch
4. status, diff
5. add
6. push
7. main
8. origin
B — Go Types, Constants and Operators
9. 3, integer
10. 1, remainder
11. 3.5, automatically
12. rune, byte
13. string, zeros
14. const, compile-time
15. false, true, false
16. not
17. assigns, compares
18. 14, 20
19. false, approximately
20. A, Itoa
C — Complete the Program
21. Fill in the three blanks
The answers are: const, float64, >.
Complete program:

package main

import "fmt"

const taxRate = 0.1

func main() {
	price := 20000.0
	qty := 3
	subtotal := price * float64(qty)
	total := subtotal + subtotal*taxRate
	isBig := total > 50000
	fmt.Println(total, isBig)
}



22. Predict the output
66000 true


Explanation:
- subtotal = 20000 × 3 = 60000
- total = 60000 + (60000 × 0.1) = 66000
- isBig = 66000 > 50000, so the result is true.
23. Which line would stop the program from compiling if you removed the conversion?
Answer: subtotal := price * qty