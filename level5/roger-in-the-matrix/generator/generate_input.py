import numpy as np
from scipy.io.wavfile import write
import zipfile as zf
import io
import os

image_file = "image.png"
audio_file = "input.wav"
input_file = "input.zip"

wave_frequency = 44100
repeat_data = 128

# create array from the message to hide, with each bit repeated 'repeat_data' times
with io.open(image_file, "rb") as in_f:
    d = in_f.read()

# transform the data into int of 4 bits
mask = np.array(list(map(lambda x: np.array([(x >> 4) & 0xF, x & 0xF]), d))).flatten()
# repeat each value repeat_data times
mask = np.repeat(mask, repeat_data)
# generate 16 freq between 400 and 16000 (step 975) for each value 0 to 15
frequencies = [400+i*975 for i in range(16)]
# tranform each value to its frequency
mask_freq = np.array([frequencies[i] for i in mask])
num_of_samples = len(mask_freq)
# apply frequencies to the wave
wave = np.sin(np.multiply(np.arange(num_of_samples) * 2 * np.pi, mask_freq) / wave_frequency)
# save audio file and zip it
write(audio_file, wave_frequency, (wave*np.iinfo('int16').max).astype('int16'))
with zf.ZipFile(input_file, 'w', compression=zf.ZIP_DEFLATED, compresslevel=9) as out_f:
    out_f.write(audio_file)
os.remove(audio_file)
