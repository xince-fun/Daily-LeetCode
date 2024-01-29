# 最近公共祖先

如何计算树上任意两点 $x$ 和 $y$ 的最近公共祖先 $lca$ 呢？

设节点 $i$ 的深度为 $depth[i]$。这可以通过一次 DFS 预处理出来。

假设 $depth[x]≤depth[y]$（否则交换两点）。我们可以先把更靠下的 $y$ 更新为 $y$ 的第 $depth[y]−depth[x]$ 个祖先节点，这样 $x$ 和 $y$ 就处在同一深度了。

如果此时 $x=y$，那么 $x$ 就是 $lca$。否则说明 $lca$ 在更上面，那么就把 xxx 和 yyy 一起往上跳。

由于不知道 $lca$ 的具体位置，只能不断尝试，先尝试大步跳，再尝试小步跳。设 $\left\lfloor\log_2 n \right\rfloor$，循环直到 $i<0$。每次循环：
- 如果 $x$ 的第 $2^i$ 个祖先节点不存在，即 $pa[x][i]=−1$，说明步子迈大了，将 $i$ 减 $1$，继续循环。
- 如果 $x$ 的第 $2^i$ 个祖先节点存在，且 ${pa}[x][i]\ne \textit{pa}[y][i]$
，说明 $lca$ 在 $pa[x][i]$ 的上面，那么更新 $x$ 为 $pa[x][i]$，更新 $y$ 为 $pa[y][i]$，将 $i$ 减 $1$，继续循环。否则，若 $pa[x][i]=pa[y][i]$，那么 $lca$ 可能在 $pa[x][i]$ 下面，由于无法向下跳，只能将 $i$ 减 $1$，继续循环。
上述做法能跳就尽量跳，不会错过任何可以上跳的机会。所以循环结束时，$x$ 与 $lca$ 只有一步之遥，即 $lca=pa[x][0]$。

> 注：你也可以用二分来理解上述算法。在 $x$ 到根节点的这条路径上猜一个点 $z$ 当作 $lca$，且 $x$ 与 $z$ 相距 $2^i$ 步。那么把 $x$ 和 $y$ 同时向上跳 $2^i$ 步，如果 $x≠y$，就说明 $lca$ 在 $z$ 的上面，否则 $lca$ 要么是 $z$，要么在 $z$ 的下面。这样一种二段性既说明了二分的正确性，又说明了每次上跳之后，步长一定要减半（类比二分查找，把搜索的区间长度减半）。

考虑到通常题目是用 $edges$ 的方式输入的，所以下面的模板先用 $edges$ 建图，再用 DFS 预处理。

```go
type TreeAncestor struct {
    depth []int
    pa    [][]int
}

func Constructor(edges [][]int) *TreeAncestor {
    n := len(edges) + 1
    m := bits.Len(uint(n))
    g := make([][]int, n)
    for _, e := range edges {
        x, y := e[0], e[1] // 节点编号从 0 开始
        g[x] = append(g[x], y)
        g[y] = append(g[y], x)
    }

    depth := make([]int, n)
    pa := make([][]int, n)
    var dfs func(int, int)
    dfs = func(x, fa int) {
        pa[x] = make([]int, m)
        pa[x][0] = fa
        for _, y := range g[x] {
            if y != fa {
                depth[y] = depth[x] + 1
                dfs(y, x)
            }
        }
    }
    dfs(0, -1)

    for i := 0; i < m-1; i++ {
        for x := 0; x < n; x++ {
            if p := pa[x][i]; p != -1 {
                pa[x][i+1] = pa[p][i]
            } else {
                pa[x][i+1] = -1
            }
        }
    }
    return &TreeAncestor{depth, pa}
}

func (t *TreeAncestor) GetKthAncestor(node, k int) int {
    for ; k > 0; k &= k - 1 {
        node = t.pa[node][bits.TrailingZeros(uint(k))]
    }
    return node
}

// 返回 x 和 y 的最近公共祖先（节点编号从 0 开始）
func (t *TreeAncestor) GetLCA(x, y int) int {
    if t.depth[x] > t.depth[y] {
        x, y = y, x
    }
    y = t.GetKthAncestor(y, t.depth[y]-t.depth[x]) // 使 y 和 x 在同一深度
    if y == x {
        return x
    }
    for i := len(t.pa[x]) - 1; i >= 0; i-- {
        px, py := t.pa[x][i], t.pa[y][i]
        if px != py {
            x, y = px, py // 同时上跳 2^i 步
        }
    }
    return t.pa[x][0]
}
```

对于二叉树，可以

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root
    }

    if left == nil {
        return right
    }
    return left
}
```