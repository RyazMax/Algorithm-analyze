import numpy as np

from matplotlib import pyplot as plt


size = []
kmp = []
kmpo = []
boyer = []

with open('kmp.log', 'r') as fin:
    for line in fin.readlines():
        splitted = line.split()
        size += [splitted[0]]
        kmp += [splitted[1]]

with open('kmpoonline.log', 'r') as fin:
    for line in fin.readlines():
        splitted = line.split()
        kmpo += [splitted[1]]

with open('boyer.log', 'r') as fin:
    for line in fin.readlines():
        splitted = line.split()
        boyer += [splitted[1]]

step = 104
scale = 1e-3

size = np.array(size, dtype=int)[::step]
kmp = np.array(kmp, dtype=float)[::step] * scale
kmpo = np.array(kmpo, dtype=float)[::step] * scale
boyer = np.array(boyer, dtype=float)[::step] * scale

plt.plot(size, kmp, label='Кнута-Морриса-Пратта')
plt.plot(size, kmpo, label='Кнута-Морриса-Пратта оптимизированный')
plt.plot(size, boyer, label='Бойера-Мура')
plt.legend()
plt.xlabel("Размер строки(эл)")
plt.ylabel("Время сортировки(мкс)")

file_to_save = input('Файл для сохранения')
if file_to_save != "":
    plt.savefig(file_to_save)

plt.show()
