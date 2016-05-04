begin = []
end = []
with open("./chameleonic_test_begin") as f:
    for b in f:
        begin.append(b.strip())

with open("./chameleonic_test_end") as f:
    for e in f:
        end.append(e.strip())

print '"###", ',
for b in begin:
    print b,
for e in end:
    for i in range(10):
        print e,
print

print '"###", ',
for e in end:
    print e,
for i in range(10):
    for e in end[0:10]:
        print e,
for i in range(10):
    for e in end[10:20]:
        print e,
for i in range(10):
    for e in end[20:30]:
        print e,
for i in range(10):
    for e in end[30:40]:
        print e,
print
