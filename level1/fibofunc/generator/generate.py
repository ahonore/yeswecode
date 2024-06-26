import random
import io
import base64
import matplotlib.pyplot as pplot
from jinja2 import Environment, select_autoescape, FileSystemLoader

seed = 1711531345694243359
flag_file = "flag.txt"
fi_file = "fi.txt"

f0 = 0
f1 = 1
for i in range(2, random.Random(seed).randint(20000, 60000)):
    t = f1
    f1 += f0
    f0 = t

index = i
with io.open(flag_file, "w") as ff:
    ff.write(str(index))

with io.open(fi_file, "w") as ff:
    ff.write(str(f1))

data = {
    "fivalue": f1
}

jin_env = Environment(autoescape=select_autoescape(), loader=FileSystemLoader(["."]))
tpl = jin_env.get_template("README.md.tpl")
readme = tpl.render(data)
with io.open("README.md", "w") as rm:
    rm.write(readme)
    rm.write("\n")
