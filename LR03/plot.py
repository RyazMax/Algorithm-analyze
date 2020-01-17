from matplotlib import pyplot as plt
import numpy as np

filename = 'random.log'

size = []
insertSort = []
mergeSort = []
radixSort = []

with open(filename, 'r') as file:
    for line in file.readlines():
        splited = line.split()
        size += [splited[0]]
        insertSort += [splited[1]]
        mergeSort += [splited[2]]
        radixSort += [splited[3]]

print("Finished reading")

def smoove(arr):
    for i in range(2, len(arr) - 2):
        arr[i] = (arr[i - 2] + arr[i - 1] + arr[i] + arr[i + 1] + arr[i + 2]) / 5
        
step = 100
scale = 1e6
size = np.array(size[::step], dtype = float)
insertSort = np.array(insertSort[::step], dtype = float) / scale
mergeSort = np.array(mergeSort[::step], dtype = float) / scale
radixSort = np.array(radixSort[::step], dtype = float) / scale

#smoove(insertSort)
#smoove(mergeSort)
#smoove(radixSort)

#plt.plot(size, insertSort, label='Сортировка вставками')
plt.plot(size, mergeSort, label='Сортировка слиянием')
plt.plot(size, radixSort, label='Поразрядная сортировка')
plt.legend()
plt.xlabel("Размер массива(эл)")
plt.ylabel("Время сортировки(мс)")

file_to_save = input('Файл для сохранения')
if file_to_save != "":
    plt.savefig(file_to_save)

plt.show()