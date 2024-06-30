flag_file = "flag.txt"
fi_file = "fi.txt"

def check_solution(fn: int) -> int:
    f0 = 0
    f1 = 1
    i = 2
    while True:
        t = f1
        f1 += f0
        f0 = t
        if f1 == fn:
            break
        i += 1
    return i

with open(flag_file, "r") as f:
    flag = f.read()

with open(fi_file, "r") as f:
    fi = f.read()

i = check_solution(int(fi))
assert i == int(flag)
print(i)
