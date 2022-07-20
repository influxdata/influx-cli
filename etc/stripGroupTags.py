import sys
import re

blacklist = [
    "Data I/O endpoints",
    "Security and access endpoints",
    "System information endpoints"
]

if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        data = f.read()
        blItems = "|".join(blacklist)
        data = re.sub(f"- +({blItems})", "", data)
        print(data)