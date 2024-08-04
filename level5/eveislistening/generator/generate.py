import random
import io
from jinja2 import Environment, select_autoescape, FileSystemLoader

# https://datatracker.ietf.org/doc/html/rfc7296#appendix-B.1
# https://datatracker.ietf.org/doc/html/rfc3526
# https://fr.wikipedia.org/wiki/%C3%89change_de_cl%C3%A9s_Diffie-Hellman


input_file = "input.bin"
flag_file = "flag.txt"

seed = 1711747128910664
public_g = 20713387470590905255573036258836721934728492752390521374064507177089990210962343331928418955536103194720235532725338679759288301525140653512087629238835871170633543949795607865540123535064826464925528784090341658306794839708039239341640351981655246845864770402104880102447526054707525127448850020961623527
public_p = 106528951761249025729412125479197260910308638225544451426813760411773819654979331756107858688322178730446171344806416830002019734743798381012666677175332885430568316533798811252472855340838402509111994536576627148671845860618445807934056330241652934528282514178025398366887626499360801730469435657805629799361

def get_message(flag_f: str) -> bytes:
    with io.open(flag_f) as f:
        return f.read().encode("ascii")


def get_int(int_file: str) -> int:
    with io.open(int_file) as intf:
        return int(intf.read())


def create_secret_key(g: int, p: int, rnd: random.Random, bits: int) -> int:
    a = rnd.getrandbits(bits)
    b = rnd.getrandbits(bits)
    A = pow(g, a, p)
    B = pow(g, b, p)
    k = pow(A, b, p)
    assert k == pow(B, a, p)
    return k


def cipher_message(m: bytes, key: int, bits: int) -> bytes:
    msg_padded = m
    len_pad = len(msg_padded) % (bits//8)
    if len_pad != 0:
        pad = "\x00" * ((bits//8) - len_pad)
        msg_padded += pad.encode("ascii")
    return (int.from_bytes(msg_padded, "little") ^ key).to_bytes(bits//8, 'little')


def check_solution(ciph_msg: bytes, g: int, p: int, bits: int) -> bytes:
    keys: dict[int, bytes] = {}
    i = 0
    while True:
        k = pow(g, i, p)
        if k in keys:
            break
        keys[k] = (int.from_bytes(ciph_msg, "little") ^ k).to_bytes(bits//8, 'little')
        i += 1
    solution: bytes = bytes()
    max_ct = 0
    for k, v in keys.items():
        # count printable ascii chars (values 32 to 127), get the one with max count
        ct = sum([32 <= i <= 127 for i in v])
        if max_ct < ct:
            solution = v
            max_ct = ct
    # cleaning
    return solution.decode("ascii").rstrip("\x00").encode("ascii")

def generate_readme(g: int, p: int) -> None:
    data = {
        "public_g": str(g),
        "public_p": str(p),
    }
    jin_env = Environment(autoescape=select_autoescape(), loader=FileSystemLoader(["."]))
    tpl = jin_env.get_template("README.md.tpl")
    readme = tpl.render(data)
    with io.open("README.md", "w") as rm:
        rm.write(readme)
        rm.write("\n")

message = get_message(flag_file)
key = create_secret_key(public_g, public_p, random.Random(seed), 1024)
cmsg = cipher_message(message, key, 1024)
assert check_solution(cmsg, public_g, public_p, 1024) == message
with io.open(input_file, "wb") as in_f:
    in_f.write(cmsg)
generate_readme(public_g, public_p)
