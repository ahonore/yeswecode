# input file generated with Faker 24.8.0
# names from fr, en, de, es, pt, it, ro

import io
from faker import Faker

seed = 26401243054462397853
input_file = "input.txt"

f = Faker(['en-US', 'fr-FR', 'de-DE', 'es-ES', 'pt-PT', 'it-IT', 'ro-RO'])
Faker.seed(seed)
with io.open(input_file, "w") as in_f:
    for i in range(200000):
        p = f.first_name() + " " + f.last_name()
        in_f.write(p)
        in_f.write("\n")
