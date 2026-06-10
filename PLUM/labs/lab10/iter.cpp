#include <iostream>
#include <vector>

using namespace std;

class Integers
{
    public:
        class Iterator
        {
            public:
                using iterator_category = std::forward_iterator_tag;
                using difference_type   = std::ptrdiff_t;
                using value_type        = int;
                using pointer           = int*;
                using reference         = int&;

                Iterator(pointer ptr, pointer end) : ptr(ptr), end(end) 
                {
                    while (this->ptr != end && *this->ptr >= 0)
                    {
                        ++this->ptr;
                    }
                }

                reference operator*() const {return *ptr;}
                pointer operator->() const {return ptr;}

                Iterator& operator++() 
                {
                    ++ptr;
                    while (ptr != end && *ptr >= 0)
                    {
                        ++ptr;
                    } 
                    return *this; 
                }
                Iterator operator++(int) { Iterator tmp = *this; ++(*this); return tmp; }
                friend bool operator== (const Iterator& a, const Iterator& b) { return a.ptr == b.ptr; };
                friend bool operator!= (const Iterator& a, const Iterator& b) { return a.ptr != b.ptr; }; 

            private:
                pointer ptr;
                pointer end;
        };

        vector<int> data;
        Integers() {}
        Integers(initializer_list<int> x)
        {
            data = x;
        }
        void Add(int x)
        {
            data.push_back(x);
        }
        Iterator begin() {return Iterator(data.data(), data.data() + data.size());}
        Iterator end() {return Iterator(data.data() + data.size(), data.data() + data.size());}
};

int main()
{
    Integers in2 = Integers({124, -1, 243, -5, -6, 3, -9});
    for (int i = 0; i < in2.data.size(); i++){
        cout << in2.data[i] << endl;
    }
    cout << "negative" << endl;
    for (int x : in2) {
        cout << x << " ";
    }
    cout << endl;
    return 1;
}