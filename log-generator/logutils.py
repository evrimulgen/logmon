from datetime import datetime
from numpy.random import choice

from logfields import LOG_FIELDS


METHODS = [
    ["GET", 0.91],
    ["POST", 0.03],
    ["UPDATE", 0.03],
    ["DELETE", 0.03],
]

URIS = [
    ["my.site.com/pages/create", 0.1],
    ["my.site.com/pages/update", 0.1],
    ["my.site.com/pages/delete", 0.1],
    ["my.site.com/login", 0.1],
    ["my.site.com/register", 0.6],
]

STATUS = [
    ["200", 0.9],
    ["201", 0],
    ["301", 0],
    ["302", 0],
    ["401", 0.05],
    ["404", 0],
    ["500", 0.05],
]

def getDate():
    return datetime.now().strftime("%Y-%m-%d")

def getTime():
    return datetime.now().strftime("%H:%M:%S")

def getMethod():
    return choice([i[0] for i in METHODS], p=[i[1] for i in METHODS])

def getURI():
    return choice([i[0] for i in URIS], p=[i[1] for i in URIS])

def getStatus(): 
    return choice([i[0] for i in STATUS], p=[i[1] for i in STATUS])

''' Write W3C log header '''
def writeLogHeader(f):
    dt = datetime.now().strftime("%d-%b-%Y %H:%M:%S")
    fields = ""
    for field in LOG_FIELDS:
        fields += field + " "

    f.write("#Version: 1.0\n")
    f.write("#Date: " + dt + "\n")
    f.write("#Fields: " + fields + "\n")


if __name__ == "__main__":
    for i in range(100):
        print getDate(), getTime(), getMethod(), getURI(), getStatus()
