from sys import argv

with open(argv[1]) as f:
	motif = argv[2]
	is_aa = False
	lines = f.readlines()
	for l in range(len(lines)-2):
		if lines[l][2] == '#':
			is_aa = True
			continue
		if is_aa:
			if lines[l][13] == motif[0] and lines[l+1][13] == motif[1] and lines[l+2][13] == motif[2]:
				ss0 = lines[l][16]
				ss1 = lines[l+1][16]
				ss2 = lines[l+2][16]
				if lines[l][16] == " ":
					ss0 = "C"
				if lines[l+1][16] == " ":
					ss1 = "C"
				if lines[l+2][16] == " ":
					ss2 = "C"
				print "{}{}{},{}{}{}, {}, {}, {}, {}, {}, {}".format(lines[l][13], lines[l+1][13], lines[l+2][13], ss0, ss1, ss2, lines[l][103:109].strip(), lines[l][109:116].strip(), lines[l+1][103:109].strip(), lines[l+1][109:116].strip(), lines[l+2][103:109].strip(), lines[l+2][109:116].strip())
