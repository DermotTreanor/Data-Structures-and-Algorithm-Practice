def anagrams(word):
    if len(word) == 1:
        return [word]
    listy = []
    for i in anagrams(word[:-1]):
        for j in range(len(i)+1):    
            listy.append(i[:j] + word[-1] + i[j:] )
    return listy


for i in anagrams("abcd"):
    print(i)
