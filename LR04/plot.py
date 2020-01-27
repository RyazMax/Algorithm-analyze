import numpy as np

from matplotlib import pyplot as plt 

filename = input("Имя файла ")

size = []
mul = []
mults = [[] for i in range(6)]

with open(filename, 'r') as fin:
    for line in fin.readlines():
        splitted = line.split()
        size += [int(splitted[0])]
        mul += [int(splitted[1])]
        for j in range(6):
            mults[j] += [int(splitted[2 + j])]

scale = 1e-6
size = np.array(size, dtype=int)
mul = np.array(mul, dtype=float) * scale

for i in range(6):
    mults[i] = np.array(mults[i], dtype=float) * scale

plt.plot(size, mul, label='Последовательный')
for i in range(6):
    plt.plot(size, mults[i], label=f'Многопоточный - {2 ** i} поток')

plt.legend()
plt.xlabel("Число строк/cтобцов")
plt.ylabel("Время перемножения(мс)")

savefilename = input("Куда сохранить? ")

plt.savefig(savefilename)
plt.show()
