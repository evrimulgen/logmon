import random
from time import sleep

from logline import LogLine
from logutils import writeLogHeader


LOG_PATH = "./log.txt"

# Open log.txt in append mode with a buffer size of 0 
f = open(LOG_PATH, "a", 0) 
writeLogHeader(f)

for _ in range(1000):
    sleep(random.uniform(0.01, 1))
    f.write(LogLine().getLine())

f.close()
