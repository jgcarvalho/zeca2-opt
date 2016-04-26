
prob_basal = 0.001

def calc_prob(pattern):
    total = 0
    sum_prob = 0.0
    ratio = {}
    probs = {}

    # inicializacao
    for i in states:
        ratio[i] = 0.0
        probs[i] = prob_basal


    for i in states:
        for j in freq:
            if pattern[0] == "??" and pattern[2] =="??":
                if j[1]==i:
                    ratio[i] += freq[j]
                    total += freq[j]
            elif not pattern[0] == "??" and pattern[2] =="??":
                if pattern[0]== j[0] and j[1]==i:
                    ratio[i] += freq[j]
                    total += freq[j]
            elif pattern[0] == "??" and not pattern[2] =="??":
                if pattern[2]== j[2] and j[1]==i:
                    ratio[i] += freq[j]
                    total += freq[j]
            elif not pattern[0] == "??" and not pattern[2] =="??":
                if pattern[0]== j[0] and pattern[2]== j[2] and j[1]==i:
                    ratio[i] += freq[j]
                    total += freq[j]

    for i in ratio:
        ratio[i] /= (total + 1)
        # print i, ratio[i]

    # calc prob["??"]
    if pattern[1] == "??":
        for i in ratio:
            for j in ratio:
                if i != j:
                    probs["??"] += (ratio[i] * ratio[j])
        sum_prob += probs["??"]
    else:
        for i in ratio:
            if pattern[1] != i:
                probs["??"] += (ratio[pattern[1]] * ratio[i])
        sum_prob += probs["??"]


    # calc states prob
    for i in probs:
        if i != "??":
            probs[i] += ratio[i] * (1.0 - probs["??"])
            sum_prob += probs[i]

    # normalize
    for i in probs:
        probs[i] = probs[i]/sum_prob

    # # print test
    # t1 = 0.0
    # for i in probs:
    #     print i, probs[i]
    #     t1 += probs[i]
    # print "Total", t1
    return probs

def print_static_rules(aa):
    for l in states:
        for c in states:
            for r in states:
                print "[ ### ][ {} ][ {} ] -> ".format(aa[1]+c,aa[2]+r),
                print "{",
                n = len(states)
                count = 1
                for i in states:
                    if i == "??":
                        print "{} : {}".format(aa[1]+i, 1.0),
                    else:
                        print "{} : {}".format(aa[1]+i, 0.0),
                    if count != n:
                        print ",",
                    count += 1
                print "}"

                print "[ {} ][ {} ][ ### ] -> ".format(aa[0]+c,aa[1]+r),
                print "{",
                n = len(states)
                count = 1
                for i in states:
                    if i == "??":
                        print "{} : {}".format(aa[1]+i, 1.0),
                    else:
                        print "{} : {}".format(aa[1]+i, 0.0),
                    if count != n:
                        print ",",
                    count += 1
                print "}"


def print_rule(aa, ss):
    probs = calc_prob(ss)
    print "[ {} ][ {} ][ {} ] ->".format(aa[0]+ss[0],aa[1]+ss[1],aa[2]+ss[2]),
    print "{",
    n = len(probs)
    c = 1
    for i in probs:
        print "{} : {}".format(aa[1]+i, probs[i]),
        if c != n:
            print ",",
        c += 1
    print "}"


states = ["??","a0","a1","b0","d0","d1","g0","g1","z0","p0","p1"]
freq = {}
probs = {}
for l in states:
    for c in states:
        for r in states:
            freq[(l,c,r)] = 0
            probs[(l,c,r)] = {}
# print freq


with open("AAA.states") as f:
    for l in f.readlines():
        d = l.strip().split(',')
        freq[(d[8],d[9],d[10])] += 1

    for l in states:
        for c in states:
            for r in states:
                print_rule(('A','A','A'),(l,c,r))
    print_static_rules(('A','A','A'))
