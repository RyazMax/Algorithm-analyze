import matplotlib.pyplot as plt

it_num = [1, 2, 3, 4, 5, 6]

rec_time = [118875, 376220, 1416050, 9622086, 98594223, 429212561]
it_time = [166305, 291932, 1896924, 10090621, 52452050, 446451502]

plt.plot(it_num, rec_time, label='Рекурсивная реализация')
plt.plot(it_num, it_time, label='Итеративная реализация')

plt.xlabel('Число итераций')
plt.ylabel('Время работы(нс)')
plt.legend()

plt.show()