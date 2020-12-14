with open("input.txt") as f:
    ops = [int(i) for i in f.readline().split(",")]


class Computer:
    def __init__(self, code):
        self.mem = (code[:] + [0] * 10000)[:10000]
        self.index = 0
        self.relative = 0
        self.inputs = []
        self.output = ""

    def parameter_mode(self, op, param):
        if op[3 - param] == "0":
            return self.mem[self.index + param]
        elif op[3 - param] == "1":
            return self.index + param
        elif op[3 - param] == "2":
            return self.mem[self.index + param] + self.relative

    def run_program(self):
        while self.index in range(len(self.mem)):
            opcode = str(self.mem[self.index]).zfill(5)
            if opcode.endswith("99"):
                break
            first = self.parameter_mode(opcode, 1)
            second = self.parameter_mode(opcode, 2)
            third = self.parameter_mode(opcode, 3)
            if opcode.endswith("1"):
                self.mem[third] = self.mem[second] + self.mem[first]
                self.index += 4
            elif opcode.endswith("2"):
                self.mem[third] = self.mem[second] * self.mem[first]
                self.index += 4
            elif opcode.endswith("3"):
                self.mem[first] = self.inputs.pop(0)
                self.index += 2
            elif opcode.endswith("4"):
                self.index += 2
                if self.mem[first] == 10:
                    print(self.output)
                    if '#'in self.output:
                        print("   ABCDEFGHI")
                    self.output = ""
                else:
                    if self.mem[first] in range(0x110000):
                        self.output += chr(self.mem[first])
                    else:
                        print(self.mem[first])
            elif opcode.endswith("5"):
                self.index = self.mem[second] if self.mem[first] != 0 else self.index + 3
            elif opcode.endswith("6"):
                self.index = self.mem[second] if self.mem[first] == 0 else self.index + 3
            elif opcode.endswith("7"):
                self.mem[third] = 1 if self.mem[first] < self.mem[second] else 0
                self.index += 4
            elif opcode.endswith("8"):
                self.mem[third] = 1 if self.mem[first] == self.mem[second] else 0
                self.index += 4
            elif opcode.endswith("9"):
                self.relative += self.mem[first]
                self.index += 2


# part one
# jump if A, B, or C is a hole and D is ground i.e. J = (!A | !B | !C) & D
script = list(map(ord, "NOT A J\nNOT B T\nOR T J\nNOT C T\nOR T J\nAND D J\nWALK\n"))
droid = Computer(ops)
droid.inputs = script
droid.run_program()

# part two
# our solution for part one fails in the case of !C & D:
#  @
# #### # #   #####
#   ABCDEFGHI
# If H is not jumpable, we will fall into the hole at E. We need H to be jumpable too in that case, so:
# J = (!A | !B) & D | (!C & D & H) = (!A | !B | (!C & H)) & D
script = list(map(ord, "NOT C J\nAND H J\nNOT A T\nOR T J\nNOT B T\nOR T J\nAND D J\nRUN\n"))
droid = Computer(ops)
droid.inputs = script
droid.run_program()