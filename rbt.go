package rbt

import ( "fmt" )

type Color bool

const (
    Red Color = false
    Black = true
)

type Node struct {
    v int
    p, l, r *Node
    c Color
}

func (n *Node) String() string {
    c := "Red"
    if n.c == Black {
        c = "Black"
    }
    return fmt.Sprintf("((%v, %v)\n\t\t%v\n\t\t%v))", n.v, c, n.l, n.r)
}

type Tree struct {
    root *Node
}

func (t *Tree) String() string {
    if t.root == nil {
        return "<nil>"
    } else {
        return t.root.String()
    }
}

func (t *Tree) Insert(v int) (*Node, error) {
    n,e := t.insert(v)
    if e != nil {
        return nil, e
    }
    t.fix(n)
    return n, nil
}

func (t *Tree) insert(v int) (*Node, error) {
    if t.root == nil {
        t.root = &Node{ v : v }
        return t.root,nil
    }
    n := t.root
    for {
        if n.v == v {
            return nil, fmt.Errorf("duplicate node")
        }

        if v < n.v {
            if n.l == nil {
                n.l = &Node{ v : v, p : n }
                return n.l, nil
            }
            n = n.l
        } else {
            if n.r == nil {
                n.r = &Node{ v : v, p : n }
                return n.r, nil
            }
            n = n.r
        }
    }

}

func (t *Tree) leftRotate(n *Node) {
    
    //    a
    // b      c
    //      d   e

    //       c
    //   a      e
    // b   d

    if n == nil || n.r == nil {
        panic("illegal left rotate")
    }

    ap := n.p
    a := n
    c := a.r
    d := c.l

    if d != nil {
        d.p = a
    }
    a.r = d

    a.p = c
    c.l = a

    c.p = ap
    if ap != nil {
        if ap.l == a {
            ap.l = c
        } else {
            ap.r = c
        }
    }

    if a == t.root {
        t.root = c
    }
}

func (t *Tree) rightRotate(n *Node) {
    //       a
    //   b     c
    // d   e

    //    b
    // d      a
    //      e   c
    
    if n == nil || n.l == nil {
        panic("illegal right rotate")
    }

    ap := n.p
    a := n
    b := a.l
    e := b.r

    if e != nil {
        e.p = a
    }
    a.l = e

    b.r = a
    a.p = b

    b.p = ap
    if ap != nil {
        if ap.l == a {
            ap.l = b
        } else {
            ap.r = b
        }
    }

    if a == t.root {
        t.root = b
    }
}

func (t *Tree) fix(z *Node) {
    for z.p != nil && z.p.c == Red {
        if z.p == z.p.p.l {
            y := z.p.p.r
            if y != nil && y.c == Red { 
                y.c = Black
                z.p.p.c = Red
                z.p.c = Black
                z = z.p.p
            } else {
                if z == z.p.r {
                    z = z.p
                    t.leftRotate(z)
                }
                z.p.c = Black
                z.p.p.c = Red
                t.rightRotate(z.p.p)
            }
        } else {
            y := z.p.p.l
            if y != nil && y.c == Red { 
                y.c = Black
                z.p.p.c = Red
                z.p.c = Black
                z = z.p.p
            } else {
                if z == z.p.l {
                    z = z.p
                    t.rightRotate(z)
                }
                z.p.c = Black
                z.p.p.c = Red
                t.leftRotate(z.p.p)
            }
        }
    }
    t.root.c = Black
}