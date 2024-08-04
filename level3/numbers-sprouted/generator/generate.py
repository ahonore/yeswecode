import random
import io

seed = 315458545295787278707809819354934510695
input_file = "input.txt"
flag_file = "flag.txt"

# v list[int]) -> int:
def find_flag(v) -> int:
    i = max(v)
    level_start_id = 0
    found = False
    level = 0
    while not found:
        level_size = 1<<level
        max_id = level_start_id+level_size
        if i in v[level_start_id:max_id]:
            id = v.index(i, level_start_id, max_id) - level_start_id
            return (id+1)*(level+1)*i
            
        level_start_id = max_id
        level += 1

# check example from readme
bt = list(map(lambda x: int(x), "271 422 688 389 872 754 147 231 564 112 211 145 275 124 333".split()))
assert find_flag(bt) == 5232

# binary tree of level 15 -> 2^16-1 values (65535)
# pick up unique values in the range (110000, 200000)
r = random.Random(seed)
bin_tree = r.sample(range(10000, 2000000), 65535)

with io.open(input_file, "w") as in_f:
    in_f.write(" ".join(list(map(lambda x: str(x), bin_tree))))

flag = find_flag(bin_tree)
with io.open(flag_file, "w") as in_f:
    in_f.write(str(flag))
