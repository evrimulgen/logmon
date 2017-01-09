import logutils
from logfields import LOG_FIELDS

LOG_DATA = {
    "c-ip": "172.224.24.114", 
    "cs-username": "-", 
    "s-ip": "206.73.118.24", 
    "s-port": "80",
    "cs-uri-query": "-", 
    "cs-bytes": "248", 
    "time-taken": "31", 
    "cs(User-Agent)": "Mozilla/4.0+(compatible;+MSIE+5.01;+Windows+2000+Server)", 
    "cs(Referrer)": "http://64.224.24.114/",
}

class LogLine:
    data = LOG_DATA.copy()
    
    def __init__(self):
        self.data["date"] = logutils.getDate()
        self.data["time"] = logutils.getTime()
        self.data["cs-method"] = logutils.getMethod()
        self.data["cs-uri-stem"] = logutils.getURI()
        self.data["sc-status"] = logutils.getStatus()
        self.data["sc-bytes"] = logutils.getSCBytes()

    def getLine(self):
        line = ""
        for field in LOG_FIELDS:
            line += self.data[field] + " "
        return line + "\n"
