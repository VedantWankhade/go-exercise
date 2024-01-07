package main

import (
	"fmt"
	"strings"
)

func main() {
    // data := make(map[string]int)
    // put(data, []string{"hello", "hi"}, 2)
    // fmt.Println(data)
    // put(data, []string{"hello", "hi"}, 3)
    // fmt.Println(data)
    // put(data, []string{"yes", "no"}, 99)
    // fmt.Println(data)
    // fmt.Println(get(data, []string{"yes", "no"}))
    // fmt.Println(data)

    // data := make(map[string]int)
    // sPut(data, "hello", 2)
    // fmt.Println(data)
    // sPut(data, "hello", 5)
    // fmt.Println(data)
    // sPut(data, "HeLlo", 89)
    // fmt.Println(data)
    // sPut(data, "ghji", 99)
    // fmt.Println(data)
    // fmt.Println(sGet(data, "ghji"))
    // fmt.Println(data)

    graph := make(map[string]map[string]bool)
    addEdge(graph, "node1", "node2")
    addEdge(graph, "node1", "node3")
    addEdge(graph, "node2", "node4")
    fmt.Println(graph)
}

// demonstrates how to use slices as keys in map by a workaround
// returns a string from array that can be used as a key in map
func keyfy(slice []string) string {
    return fmt.Sprintf("%q", slice)
}

func put(data map[string]int, slice []string, val int) {
    data[keyfy(slice)] = val
}

func get(data map[string]int, slice []string) int {
    return data[keyfy(slice)]
}

// --------------------------------------
// demonstrates how to change behaviour for even the comparable key types
// strings as keys are checked with == which does not ignore case 
// following way can be used to store keys ignoring cases
func ignCase(s string) string {
    return strings.ToLower(s)
}

func sPut(data map[string]int, sKey string, val int) {
   data[ignCase(sKey)] = val 
}

func sGet(data map[string]int, sKey string) int {
    return data[ignCase(sKey)]
}
//------------------------------------------------------

// go doesnt have sets, but maps can be used as sets, the keys of map as values of the set
// and the values of the map can just be booleans
// following demonstrates a graph that is a map with strings as keys. and a map(set) of strings

func addEdge(graph map[string]map[string]bool, from, to string) {
    edges := graph[from]
    // must check if the map at the given key is not initialized, if its not then cannot assign to any of its values
    if edges == nil {
        edges = make(map[string]bool)
        graph[from] = edges
    }
    // we are setting the 'to' = true but we are essentially just adding 'to' to the set of nodes from a 'key' node
    edges[to] = true
}

func hasEdge(graph map[string]map[string]bool, from, to string) bool {
    return graph[from][to]
}
