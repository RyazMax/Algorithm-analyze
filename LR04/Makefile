main:
	g++ main.cpp -o main.out

tests: test_matrix.cpp matrix.hpp
	g++ test_matrix.cpp -o test.out
	./test.out

clean:
	rm *.out *.o