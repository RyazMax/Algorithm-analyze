main:
	g++ main.cpp -o main.out -pthread

profile:
	g++ profile.cpp -o profile.out -pthread -O2

tests: test_matrix.cpp matrix.hpp
	g++ test_matrix.cpp -o test.out
	./test.out

clean:
	rm *.out *.o