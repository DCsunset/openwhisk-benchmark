#!/usr/bin/python3
import glob
import sys

resultFiles = glob.glob("*.txt")

throughput = 0
latency = 0
errors = 0
for resultFile in resultFiles:
	with open(resultFile, "r") as f:
		lines = f.readlines()
	
	throughput += float(lines[0].split(" ")[1])
	latency += float(lines[1].split(" ")[1])
	errors += int(lines[2].split(" ")[1])

latency /= len(resultFiles)

clientNum = len(resultFiles)
args = sys.argv[1:]
if len(args) < 1:
	print("Client Num:", clientNum)
	print("Throughput:", throughput)
	print("Latency:", latency)
	print("Errors:", errors)
else:
	filename = args[0]
	with open(filename, "a") as f:
		print(f"{clientNum},{throughput},{latency},{errors}", file=f)

