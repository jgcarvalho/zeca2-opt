from sys import argv
aa = ['A','C','D','E','F','G','H','I','K','L','M','N','P','Q','R','S','T','V','Y','W']
states = ["**","??","a0","a1","b0","d0","d1","g0","g1","z0","p0","p1"]

begin_rules = {}
end_rules = {}

def init_special_rules():
    for i in aa:
        for j in aa:
            for u in states:
                for v in states:
                    begin_rules[("###",i+u,j+v)] = {}
                    end_rules[(i+u,j+v,"###")] = {}

def calc_special_rules(fn):
    with open(fn) as f:
        for l in f:
            d = l.strip().split()
            for i in range(9,55,4):
                if d[i] in begin_rules[("###",d[3],d[5])].keys():
                    begin_rules[("###",d[3],d[5])][d[i]] += float(d[i+2])/240

                else:
                    begin_rules[("###",d[3],d[5])][d[i]] = float(d[i+2])/240

                if d[i] in end_rules[(d[1],d[3],"###")].keys():
                    end_rules[(d[1],d[3],"###")][d[i]] += float(d[i+2])/240
                else:
                    end_rules[(d[1],d[3],"###")][d[i]] = float(d[i+2])/240

def normalize_special_rules():
    for i in begin_rules:
        total = 0.0
        for j in begin_rules[i]:
            total += begin_rules[i][j]
        for j in begin_rules[i]:
            begin_rules[i][j] /= total

    for i in end_rules:
        total = 0.0
        for j in end_rules[i]:
            total += end_rules[i][j]
        for j in end_rules[i]:
            end_rules[i][j] /= total


def print_special_rules():
    for m in begin_rules:
        print "[ {} ][ {} ][ {} ] -> ".format(m[0],m[1], m[2]),
        print "{",
        n = len(begin_rules[m])
        c = 1
        for i in begin_rules[m]:
            print "{} : {}".format(i, begin_rules[m][i]),
            if c != n:
                print ",",
            c += 1
        print "}"

    for m in end_rules:
        print "[ {} ][ {} ][ {} ] -> ".format(m[0],m[1], m[2]),
        print "{",
        n = len(end_rules[m])
        c = 1
        for i in end_rules[m]:
            print "{} : {}".format(i, end_rules[m][i]),
            if c != n:
                print ",",
            c += 1
        print "}"


if __name__ == '__main__':
    init_special_rules()
    calc_special_rules(argv[1])
    normalize_special_rules()
    print_special_rules()
