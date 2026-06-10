#include <iostream>
#include <vector>
#include <cmath>
#include <functional>

using namespace std;

template<typename T, int M, int N>
class Matrix
{
    using Func = std::function<T(int,int)>;
    public:
        Matrix(Func f)
        {
            std::cout << "Creating matrix" << std::endl;
            this->f = f;
        }
        Func f;
        T GetVal(int i, int j)
        {
            return f(i, j);
        }
        Matrix operator + (const Matrix& matrix) const
        {
            Func f1 = this->f;
            Func f2 = matrix.f;
            return Matrix([f1, f2](int i, int j) {
            return f1(i, j) + f2(i, j);
            });
        }
        template<int K>
        Matrix<T, M, K> operator*(const Matrix<T, N, K>& matrix) const
        {
            Func f1 = this->f;
            Func f2 = matrix.f;
            return Matrix<T, M, K>([f1, f2](int i, int j) {
                T sum = 0;
                for (int k = 0; k < N; k++) {
                    sum += f1(i, k) * f2(k, j);
                }
            return sum;
            });
        }
        T operator[](int k)
        {
            return f(k/N,k%N);
        }

};

int main()
{
    Matrix<int, 5, 4> m1([](int i, int j){return i*i + j*j;});
    Matrix<int, 5, 4> m2([](int i, int j){return 2*i*j;});
    cout << m1.GetVal(2, 3) << endl;
    cout << m2.GetVal(2, 3) << endl;
    Matrix<int, 5, 4> m3 = m1 + m2;
    cout << m3.GetVal(2, 3) << endl;
    cout << m3[11] << endl;
    Matrix<int, 4, 3> m4([](int i, int j){return -i;});
    Matrix<int, 5, 3> m5 = m2 * m4;
    cout << m5[4] << endl;
    return 1;
}