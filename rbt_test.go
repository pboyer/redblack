package rbt

import ("testing"; "math/rand")

func verifyRBT(tr *Tree, t *testing.T) {
    // prop1 is checked implicitly
    if tr.root == nil {
        return
    }

    n := tr.root
    if n.c != Black { // prop2
        t.Fatalf("root not black \n%v", tr)
    }

    verifyBST( n, t )
    verifyRBTSubtree( n, t )
}

func verifyBST( n *Node, t *testing.T){
    if n == nil {
        return
    }

    if n.l != nil {
        if n.l.v > n.v {
            t.Fatalf("illegal left child > parent")
        }

        if n.l.p != n {
            t.Fatalf("illegal left child parent ptr")
        }
    }

    if n.r != nil {
        if n.r.v < n.v {
            t.Fatalf("illegal right child < parent")
        }

        if n.r.p != n {
            t.Fatalf("illegal right child parent ptr")
        }
    }

    verifyBST( n.l, t )
    verifyBST( n.r, t )
}

func verifyRBTSubtree(n *Node, t *testing.T) int {
    if n == nil { // prop3
        return 1
    }

    if n.c == Red { // prop4
        if n.l != nil && n.l.c != Black {
            t.Fatalf("illegal red left child of red node")
        }

        if n.r != nil && n.r.c != Black {
            t.Fatalf("illegal red right child of red node")
        }
    }

    // prop5 and recurse
    bhr := verifyRBTSubtree(n.r, t)
    bhl := verifyRBTSubtree(n.l, t)

    if bhl != bhr {
        t.Fatalf("illegal black height incompatibility")
    }

    if n.c == Black {
        return bhl + 1
    }
    return bhl
}

func insertCheck(tr *Tree, val int, t *testing.T){
    _, e := tr.Insert(val)
    if e != nil {
        // dup
        return
    }
    verifyRBT(tr, t)
}

func insertCheckBST(tr *Tree, val int, t *testing.T){
    _, e := tr.insert(val)
    if e != nil {
        t.Fatalf("%v", e)
    }
    verifyBST(tr.root, t)
}

func TestRBT(t *testing.T){
    tr := &Tree{}
    
    for i := 0; i < 1000; i++ {
        insertCheck(tr, rand.Int(), t)
    }
}

func TestBST(t *testing.T){
    tr := &Tree{}

    insertCheckBST(tr, 1, t)
    insertCheckBST(tr, 2, t)
    insertCheckBST(tr, 0, t)
    insertCheckBST(tr, -1, t)
    insertCheckBST(tr, -5, t)

    tr.leftRotate(tr.root)
    verifyBST(tr.root, t)

    tr.rightRotate(tr.root)
    verifyBST(tr.root, t)
}