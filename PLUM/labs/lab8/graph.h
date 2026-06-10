#pragma once
#include <iostream>
#include <vector>


template<int N>
class Dsa
{
    std::vector<int> dsa;
    public:
        Dsa() : dsa(N)
        {
            for (int i = 0; i < N; i++)
            {
                dsa[i] = i;
            }
        }
        int Find(int u)
        {
            return dsa[u];
        }
        void Union(int u, int v)
        {
            if (dsa[u] == dsa[v])
            {
                return;
            } 
            int t = dsa[v];
            for (int i = 0; i < N; i++)
            {
                if (dsa[i] == t)
                {
                    dsa[i] = dsa[u];
                }
            }
        }
};

template<int N, bool Acyclic>
class Graph
{
    protected:
        std::vector<std::vector<int>> matrix;
    public:
        Graph()
        {
            std::cout << "Creating graph" << std::endl;
            matrix = std::vector<std::vector<int>>(N, std::vector<int>(N, 0));
        }
        void AddEdge(int u, int v)
        {
            matrix[u][v] = matrix[v][u] = 1;
        }
        bool CheckAdjacency(int u, int v)
        {
            return matrix[u][v] == 1;
        }
        void print()
        {
            std::cout << "printing graph" << std::endl;
            for (auto i : matrix)
            {
                for (auto j : i)
                {
                    std::cout << j << " ";
                }
                std::cout << std::endl;
            }
        }
};

template<int N>
class Graph<N, true> : public Graph<N, false> 
{
    public: 
        Dsa<N>dsa;
        void AddEdge(int u, int v)
        {
            if (dsa.Find(u) == dsa.Find(v))
            {
                std::cout << u << " and " << v << " create cycle. not added" << std::endl;
                return;
            }
            dsa.Union(u, v);
            this->matrix[u][v] = this->matrix[v][u] = 1;
        }
};