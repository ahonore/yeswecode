import numpy as np
from scipy.io.wavfile import read
import zipfile as zf
import io
from io import BytesIO

image_file = "image.png"
input_file = "input.zip"

repeat_data = 128

# load audio file from zip
with zf.ZipFile(input_file, 'r') as out_f:
    data_bytes = out_f.read(out_f.filelist[0])
samplerate, data = read(BytesIO(data_bytes))
# normalize data between -1 and 1
data = data/samplerate
# regroup data by blocs of repeat_data
data = np.reshape(data, (-1, repeat_data))
# compute frequencies on repeat_data levels
fq = np.fft.fftfreq(repeat_data)
# perform fft on each bloc and get the index of the maximum amplitude on a bloc
out_idx = np.array([np.argmax(np.abs(np.fft.fft(d))) for d in data])
# get the real frequency from each indices
out_freq = np.array([abs(fq[i]) * samplerate for i in out_idx])
# get all the distinct frequencies sorted: give the relation 4 bits value (index of freq_table) -> frequency
freq_table = np.sort(np.array(list(set(out_freq))))
# get the final data by mapping previously computed frequencies to value from 0 to 15 thanks to freq_table
# transform the array to array of couple of 4-bit integers
final_data = np.array([np.where(freq_table == i)[0][0] for i in out_freq]).reshape(-1, 2)
# transform to an array of bytes by gathering couples into 1 byte each
final_data = np.array([(i[0] << 4) | i[1] for i in final_data])
# compare to the given image containing the flag to validate
with io.open(image_file, "rb") as f:
    in_b = np.frombuffer(f.read(), dtype=np.uint8)
assert np.equal(final_data, in_b).all()
print("OK")
