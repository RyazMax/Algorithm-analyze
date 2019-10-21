import matplotlib.pyplot as plt
import numpy as np

filename = input('Имя файла')

size = []
regMul = []
vinMul = []
optVinMul = []

with open(filename, 'r') as file:
    for line in file.readlines():
        splited = line.split()
        size += [splited[0]]
        regMul += [splited[1]]
        vinMul += [splited[2]]
        optVinMul += [splited[3]]

regMul = np.array(regMul, dtype = float) / 1000000
vinMul = np.array(vinMul,dtype=float) / 1000000
optVinMul = np.array(optVinMul, dtype=float) / 1000000

plt.plot(size, regMul, label='Обычный алгоритм')
plt.plot(size, vinMul, label='Алгоритм Винограда')
plt.plot(size, optVinMul, label='Оптимизированный Винограда')
plt.legend()
plt.xlabel("Размер матрицы(эл)")
plt.ylabel("Время умножения (мс)")

plt.show()