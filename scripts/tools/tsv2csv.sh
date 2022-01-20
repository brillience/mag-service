#!/usr/bin/env python
import csv, sys
csv.writer(sys.stdout).writerows(csv.reader(sys.stdin,delimiter='\t'))