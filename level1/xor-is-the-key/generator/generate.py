import random
import io
import base64
import matplotlib.pyplot as pplot
from jinja2 import Environment, select_autoescape, FileSystemLoader

flag_file = "flag.txt"

seed = 1711747128910664
r = random.Random(seed)
k1 = int.to_bytes(r.getrandbits(256), 32, "little")
k2 = int.to_bytes(r.getrandbits(256), 32, "little")
k3 = int.to_bytes(r.getrandbits(256), 32, "little")
flag = int.to_bytes(r.getrandbits(256), 32, "little")

with io.open(flag_file, "w") as flag_f:
    flag_f.write(flag.hex())

def xor_bytes(b0: bytes, b1: bytes) -> bytes:
    return bytes(a ^ b for a, b in zip(b0, b1))

k2k1 = xor_bytes(k2, k1)
k2k3 = xor_bytes(k2, k3)
flagk1k3k2 = xor_bytes(flag, xor_bytes(k1, k2k3))

data = {
    "k1value": k1.hex(),
    "k2k1value": k2k1.hex(),
    "k2k3value": k2k3.hex(),
    "flagk1k3k2": flagk1k3k2.hex(),
    }

jin_env = Environment(autoescape=select_autoescape(), loader=FileSystemLoader(["."]))
tpl = jin_env.get_template("README.md.tpl")
readme = tpl.render(data)
with io.open("README.md", "w") as rm:
    rm.write(readme)
    rm.write("\n")

def check_solution(k1: str, k2k3: str, flagk1k3k2: str) -> str:
    k1b = bytes.fromhex(k1)
    k2k3b = bytes.fromhex(k2k3)
    flagk1k3k2b = bytes.fromhex(flagk1k3k2)
    flagb = xor_bytes(flagk1k3k2b, xor_bytes(k1b, k2k3b))
    return flagb.hex()

assert check_solution(data["k1value"], data["k2k3value"], data["flagk1k3k2"]) == flag.hex()
