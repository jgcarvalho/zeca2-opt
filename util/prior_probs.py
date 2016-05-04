import math
from sys import argv

prob_basal = 0.001
states = ["**","??","a0","a1","b0","d0","d1","g0","g1","z0","p0","p1"]

def print_rule_init(aa,ss,prob):
    # p = calc_prob(ss,prob)
    print "[ {} ][ {} ][ {} ] ->".format(aa[0]+ss[0],aa[1]+ss[1],aa[2]+ss[2]),
    print "{",
    n = len(prob)
    c = 1
    for i in prob:
        print "{} : {}".format(aa[1]+i, prob[i]),
        if c != n:
            print ",",
        c += 1
    print "}"


def print_rule_transition(aa,ss):
    print "[ {} ][ {} ][ {} ] ->".format(aa[0]+ss[0],aa[1]+ss[1],aa[2]+ss[2]),
    print "{",
    n = len(states)
    c = 1
    for i in states:
        if i == ss[1]:
            print "{} : {}".format(aa[1]+i, 0.9),
        else:
            print "{} : {}".format(aa[1]+i, 0.1/(len(states)-1)),
        if c != n:
            print ",",
        c += 1
    print "}"



def ajust_probs(probs):
    new_probs = {}
    inde = 0.0
    for i in probs:
        for j in probs:
            if i != j:
                inde += probs[i]*probs[j]
    # print inde

    total = 0.0
    for s in states:
        if s == "**":
            new_probs[s] = 1.0 - inde
        else:
            new_probs[s] = probs[s]*(1.0-inde)
        total += new_probs[s]

    for i in new_probs:
        new_probs[i] /= total

    # print new_probs
    return new_probs

def ajust_probs_entropy(probs):
    new_probs = {}
    entropy = 0.0
    max_entropy = 0.0
    rel_entropy = 0.0
    for i in range(len(states)):
        max_entropy -= 1.0/len(states)*math.log10(1.0/len(states))
    # print max_entropy

    for p in probs.values():
        entropy -= p*math.log10(p)
    # print entropy
    rel_entropy = entropy/max_entropy
    # print rel_entropy

    total = 0.0
    for s in states:
        if s == "**":
            new_probs[s] = rel_entropy
        else:
            new_probs[s] = probs[s]*(1.0-rel_entropy)
        total += new_probs[s]

    for i in new_probs:
        new_probs[i] /= total

    # print total
    # print new_probs
    return new_probs




if __name__ == '__main__':
    with open(argv[1]) as f:
        motif = tuple(argv[2])
        total = 0.0
        freq = {}
        probs = {}
        for l in f.readlines():
            d = l.strip().split(',')
            if d[9] not in freq.keys():
                freq[d[9]] = 0
            freq[d[9]] += 1
            total += 1.0
        for i in states:
            if i not in freq.keys():
                probs[i] = 0.001
            else:
                if freq[i]/total < 0.001:
                    probs[i] = 0.001
                else:
                    probs[i] = freq[i]/total

        # print freq
        # print probs
        new_p = ajust_probs(probs)
        # print new_p
        new_p = ajust_probs_entropy(probs)
        # print new_p

        for l in states:
            for c in states:
                for r in states:
                    if c[1] == "**":
                        print_rule_init(motif,(l,c,r), new_p)
                    else:
                        print_rule_transition(motif, (l,c,r))
