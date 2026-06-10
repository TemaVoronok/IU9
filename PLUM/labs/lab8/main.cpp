#include <iostream>
#include "Graph.h"

using namespace std;


int main()
{
    Graph<5, false>g;
    g.AddEdge(3, 2);
    g.AddEdge(4, 1);
    g.AddEdge(4, 2);
    g.AddEdge(2, 1);

    cout << g.CheckAdjacency(1, 4) << endl;
    cout << g.CheckAdjacency(2, 2) << endl;

    g.print();

    Graph<3, true>g1;
    g1.AddEdge(1, 2);
    g1.AddEdge(2, 0);
    g1.AddEdge(0, 1);
    g1.print();
    return 0;
}