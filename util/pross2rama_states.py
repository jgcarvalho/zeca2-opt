import math
from sys import argv

aa = {'CYS': 'C', 'ASP': 'D', 'SER': 'S', 'GLN': 'Q', 'LYS': 'K',
     'ILE': 'I', 'PRO': 'P', 'THR': 'T', 'PHE': 'F', 'ASN': 'N',
     'GLY': 'G', 'HIS': 'H', 'LEU': 'L', 'ARG': 'R', 'TRP': 'W',
     'ALA': 'A', 'VAL':'V', 'GLU': 'E', 'TYR': 'Y', 'MET': 'M'}


alpha0 = [(-60,-40,25)]
alpha1 = [(60,40,25)]

beta0 = [(-95, 120, 20), (-115, 130, 35), (-125, 145, 30), (-135, 158, 30), (-145, 170, 30), (-155, 180, 30), (-155, -180, 30)]

delta0 =  [(-70, -30, 25),(-80, -20, 25),(-90, -10, 25),(-100, 0, 25),(-110, 10, 25)]
delta1 = [(70, 30, 25),(80, 20, 25),(90, 10, 25),(100, 0, 25),(110, -10, 25)]

gamma0 = [(-85, 70, 20)]
gamma1 = [(75, -50, 20)]

zeta0 = [(-135, 75, 20)]

PII0 = [(-55, 130, 20), (-60, 140, 20), (-65, 150, 20),(-70, 160, 20),(-75, 170, 20), (-80, 180, 20), (-80, -180, 20)]
PII1 = [(55, -130, 20),(60, -140, 20), (65, -150, 20),(70, -160, 20),(75, -170, 20), (80, -180, 20), (80, 180, 20)]

def distance(center, angles):
    return math.sqrt((center[0]-angles[0])**2 + (center[1]-angles[1])**2)

def test(r, angles):
    if distance((r[0],r[1]), angles) < r[2]:
        return True
    else:
        return False

def classify(angles):
    for r in alpha0:
        if test(r, angles):
            return "a0"
    for r in alpha1:
        if test(r, angles):
            return "a1"
    for r in beta0:
        if test(r, angles):
            return "b0"
    for r in delta0:
        if test(r, angles):
            return "d0"
    for r in delta1:
        if test(r, angles):
            return "d1"
    for r in gamma0:
        if test(r, angles):
            return "g0"
    for r in gamma1:
        if test(r, angles):
            return "g1"
    for r in zeta0:
        if test(r, angles):
            return "z0"
    for r in PII0:
        if test(r, angles):
            return "p0"
    for r in PII1:
        if test(r, angles):
            return "p1"
    return "??"

with open(argv[1]) as f:
    seq = []
    ss = []
    is_aa = False
    for l in f.readlines():
        if l[0] == "[":
            is_aa = True
            continue
        if is_aa:
            d = l.strip().split()
            # print d
            # print "{}, {}, {}, {}, {}".format(d[0], aa[d[1]], d[2], d[3], classify((float(d[4]),float(d[5]))))
            seq.append(aa[d[1]] + "??")
            ss.append(aa[d[1]] + classify((float(d[4]),float(d[5]))))
    print seq
    print ss
        # print ", {}".format(classify((float(d[2]), float(d[3]))) + classify((float(d[4]), float(d[5]))) + classify((float(d[6]), float(d[7]))))
        # print ",{},{},{}".format(classify((float(d[2]), float(d[3]))), classify((float(d[4]), float(d[5]))), classify((float(d[6]), float(d[7]))))
