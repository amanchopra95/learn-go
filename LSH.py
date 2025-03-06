# LSH

a = "flying fish flew by the space station"

# Shingling
def shingle(text: str, k: int):
    shingle_set = []
    for i in range(len(text) - k + 1):
        shingle_set.append(text[i:i + k])

    return set(shingle_set)



a_shingle = shingle(a, 2)
print(a_shingle)

